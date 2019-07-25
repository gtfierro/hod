package main

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"strings"
	"syscall"
	"time"

	"github.com/gtfierro/hod/config"
	hod "github.com/gtfierro/hod/db"
	query "github.com/gtfierro/hod/lang"
	sparql "github.com/gtfierro/hod/lang/ast"
	"github.com/gtfierro/hod/server"
	"github.com/gtfierro/hod/turtle"
	"github.com/gtfierro/hod/version"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/gtfierro/hod/proto"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/pkg/profile"
	//"github.com/immesys/bw2/crypto"
	//"github.com/immesys/bw2bind"
	"github.com/montanaflynn/stats"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func init() {
	fmt.Println(version.LOGO)
}

type ResultMap hod.ResultMap

func benchLoad(c *cli.Context) error {
	if c.NArg() == 0 {
		log.Fatal("Need to specify a turtle file to load")
	}
	filename := c.Args().Get(0)
	p := turtle.GetParser()
	ds, duration := p.Parse(filename)
	rate := float64((float64(ds.NumTriples()) / float64(duration.Nanoseconds())) * 1e9)
	fmt.Printf("Loaded %d triples, %d namespaces in %s (%.0f/sec)\n", ds.NumTriples(), ds.NumNamespaces(), duration, rate)
	return nil
}

func rebuildServer(c *cli.Context) error {
	cfg, err := config.ReadConfig(c.String("config"))
	if err != nil {
		log.Error(err)
		return err
	}
	dbdir := strings.TrimSuffix(cfg.DBPath, "/")
	log.Warningf("Removing %s", dbdir)
	if err := os.RemoveAll(dbdir); err != nil {
		log.Error(errors.Wrapf(err, "Could not remove old dir %s", dbdir))
	}
	return startServer(c)
}

func startCLI(c *cli.Context) error {
	cfg, err := config.ReadConfig(c.String("config"))
	if err != nil {
		log.Error(err)
		return err
	}
	if cfg.EnableCPUProfile {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.EnableMEMProfile {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.EnableBlockProfile {
		defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	}
	mdb, err := hod.NewHodDB(cfg)
	if err != nil {
		log.Error(err)
		defer mdb.Close()
		return err
	}

	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	done := make(chan struct{})
	go runInteractiveQuery(mdb, interruptSignal)
	go func() {
		killSignal := <-interruptSignal
		switch killSignal {
		case os.Interrupt, syscall.SIGINT:
			log.Warning("Caught SIGINT; closing...")
			mdb.Close()
		case syscall.SIGTERM:
			log.Warning("Caught SIGTERM; closing...")
			mdb.Close()
		default:
			log.Warning(killSignal)
		}
		close(done)
		//return nil
	}()
	<-done
	//return runInteractiveQuery(mdb)
	return nil
}

func startServer(c *cli.Context) error {
	cfg, err := config.ReadConfig(c.String("config"))
	if err != nil {
		log.Error(err)
		return err
	}
	cfg.ReloadOntologies = false
	if cfg.EnableCPUProfile {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.EnableMEMProfile {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.EnableBlockProfile {
		defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	}
	db, err := hod.NewHodDB(cfg)
	if err != nil {
		log.Error(err)
		return err
	}
	defer db.Close()
	var srv *http.Server
	if cfg.EnableHTTP {
		srv = server.StartHodServer(db, cfg)
	}

	// TODO: grpc configure/enable
	if true {
		port := 47809
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Error(err)
			return err
		}
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST"},
			AllowCredentials: true,
			Debug:            false,
		})
		s := 100 * 1024 * 1024
		grpcServer := grpc.NewServer(
			grpc.MaxRecvMsgSize(s),
			grpc.MaxSendMsgSize(s),
		)
		proto.RegisterHodDBServer(grpcServer, db)
		go grpcServer.Serve(lis)
		//log.Info("Serving on ", lis.Addr().String())

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(s), grpc.MaxCallSendMsgSize(s))}
		if err := proto.RegisterHodDBHandlerFromEndpoint(context.Background(), mux, ":47809", opts); err != nil {
			log.Error(err)
			return err
		}
		httpmux := http.NewServeMux()
		httpmux.Handle("/", mux)

		go http.ListenAndServe(":47808", c.Handler(mux))
	}

	//	if cfg.EnableBOSSWAVE {
	//		client := bw2bind.ConnectOrExit(cfg.BW2_AGENT)
	//		client.SetEntityFileOrExit(cfg.BW2_DEFAULT_ENTITY)
	//		client.OverrideAutoChainTo(true)
	//
	//		svc := client.RegisterService(cfg.HodURI, "s.hod")
	//		iface := svc.RegisterInterface("_", "i.hod")
	//		queryChan, err := client.Subscribe(&bw2bind.SubscribeParams{
	//			URI: iface.SlotURI("query"),
	//		})
	//		if err != nil {
	//			err = errors.Wrapf(err, "Could not subscribe to HodDB query slot URI %s", iface.SlotURI("query"))
	//			log.Error(err)
	//			return err
	//		}
	//
	//		log.Notice("Serving query URI", iface.SlotURI("query"))
	//
	//		const QueryPIDString = "2.0.10.1"
	//		//var QueryPID = bw2bind.FromDotForm(QueryPIDString)
	//		const ResponsePIDString = "2.0.10.2"
	//		var ResponsePID = bw2bind.FromDotForm(ResponsePIDString)
	//		type hodQuery struct {
	//			Query string
	//			Nonce string
	//		}
	//		type hodResponse struct {
	//			Count   int
	//			Nonce   string
	//			Elapsed int64
	//			Rows    []ResultMap
	//			Error   string
	//			Errors  []string
	//		}
	//
	//		handleBOSSWAVEQuery := func(msg *bw2bind.SimpleMessage) {
	//			var inq hodQuery
	//			po := msg.GetOnePODF(QueryPIDString)
	//			if po == nil {
	//				return
	//			}
	//			if obj, ok := po.(bw2bind.MsgPackPayloadObject); !ok {
	//				log.Error("Payload 2.0.10.1 was not MsgPack")
	//				return
	//			} else if err := obj.ValueInto(&inq); err != nil {
	//				log.Error(errors.Wrap(err, "Could not unmarshal into a hod query"))
	//				return
	//			}
	//			log.Info("Serving query", inq.Query)
	//
	//			var response hodResponse
	//			returnErr := func(err error) {
	//				log.Error(errors.Wrap(err, "Could not parse hod query"))
	//				response = hodResponse{
	//					Nonce: inq.Nonce,
	//					Error: err.Error(),
	//				}
	//				responsePO, err := bw2bind.CreateMsgPackPayloadObject(ResponsePID, response)
	//				if err != nil {
	//					log.Error(errors.Wrap(err, "Could not serialize hod response"))
	//					return
	//				}
	//				if err = iface.PublishSignal("result", responsePO); err != nil {
	//					log.Error(errors.Wrap(err, "Could not send hod response"))
	//					return
	//				}
	//			}
	//
	//			q, err := query.Parse(inq.Query)
	//			if err != nil {
	//				returnErr(err)
	//				return
	//			}
	//
	//			newq := q.Copy()
	//			//toKey, _ := crypto.UnFmtKey(msg.From)
	//			//var dbs []string
	//			//if newq.From.AllDBs || newq.From.Empty() {
	//			//	dbs = db.Databases()
	//			//} else {
	//			//	dbs = newq.From.Databases
	//			//}
	//			newq.From.Databases = []string{}
	//			newq.From.AllDBs = false
	//			//for _, dbname := range dbs {
	//			//key, zero, err := client.ResolveLongAlias(dbname)
	//			//if err != nil {
	//			//	log.Error(errors.Wrap(err, "Could not resolve alias"))
	//			//	continue
	//			//}
	//			//if zero {
	//			//	log.Error(errors.New("No alias by that name"))
	//			//	continue
	//			//}
	//			//if allowVK("06DZ14k6dhRognGUoSHofD8oS8CUaWwq4tH-FPW13UU=/hod", key, toKey, client) {
	//			//newq.From.Databases = append(newq.From.Databases, dbname)
	//			//}
	//			//}
	//
	//			result, err := db.RunQuery(newq)
	//			if err != nil {
	//				log.Error(errors.Wrap(err, "Could not run query"))
	//				returnErr(err)
	//				return
	//			}
	//
	//			response = hodResponse{
	//				Count:   result.Count,
	//				Errors:  result.Errors,
	//				Elapsed: result.Elapsed.Nanoseconds(),
	//				Nonce:   inq.Nonce,
	//			}
	//			for _, row := range result.Rows {
	//				response.Rows = append(response.Rows, ResultMap(row))
	//			}
	//
	//			responsePO, err := bw2bind.CreateMsgPackPayloadObject(ResponsePID, response)
	//			if err != nil {
	//				log.Error(errors.Wrap(err, "Could not serialize hod response"))
	//				return
	//			}
	//			if err = iface.PublishSignal("result", responsePO); err != nil {
	//				log.Error(errors.Wrap(err, "Could not send hod response"))
	//				return
	//			}
	//		}
	//
	//		go func() {
	//			for msg := range queryChan {
	//				go handleBOSSWAVEQuery(msg)
	//			}
	//		}()
	//	}
	interruptSignal := make(chan os.Signal, 1)
	signal.Notify(interruptSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	killSignal := <-interruptSignal
	switch killSignal {
	case os.Interrupt, syscall.SIGINT:
		log.Warning("SIGINT")
		//db.Close()
		break
	case syscall.SIGTERM:
		log.Warning("SIGTERM")
		//db.Close()
	}

	if srv != nil {
		srv.Shutdown(context.Background())
	}
	return nil
}

func doQuery(c *cli.Context) error {
	cfg, err := config.ReadConfig(c.String("config"))
	if err != nil {
		log.Error(err)
		return err
	}
	cfg.ReloadOntologies = false
	if cfg.EnableCPUProfile {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.EnableMEMProfile {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.EnableBlockProfile {
		defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	}
	db, err := hod.NewHodDB(cfg)
	if err != nil {
		log.Error(err)
		return err
	}
	defer db.Close()
	var (
		q   *sparql.Query
		res *hod.QueryResult
	)
	if c.String("query") != "" {
		q, err = query.Parse(c.String("query"))
		if err != nil {
			log.Fatal(err)
		}
	} else if c.String("file") != "" {
		filebytes, err := ioutil.ReadFile(c.String("file"))
		if err != nil {
			log.Fatal(err)
		}
		q, err = query.Parse(string(filebytes))
		if err != nil {
			log.Fatal(err)
		}
	}
	res, err = db.RunQuery(q)
	if err != nil {
		log.Fatal(err)
	}
	return res.DumpToCSV(c.Bool("prefixes"), db, os.Stdout)
}

func dump(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("Need to specify a turtle file to load")
	}
	filename := c.Args().Get(0)
	p := turtle.GetParser()
	ds, _ := p.Parse(filename)
	for _, triple := range ds.Triples {
		var s = triple.Subject.Value
		var p = triple.Predicate.Value
		var o = triple.Object.Value
		for pfx, full := range ds.Namespaces {
			if triple.Subject.Namespace == full {
				s = pfx + ":" + s
			}
			if triple.Predicate.Namespace == full {
				p = pfx + ":" + p
			}
			if triple.Object.Namespace == full {
				o = pfx + ":" + o
			}
		}
		fmt.Printf("%s\t%s\t%s\n", s, p, o)
	}
	return nil
}

func classGraph(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("Need to specify a turtle file to load")
	}
	filename := c.Args().Get(0)
	p := turtle.GetParser()
	ds, _ := p.Parse(filename)

	name := gethash() + ".gv"
	f, err := os.Create(name)
	if err != nil {
		log.Error(err)
		return err
	}

	nodes := make(map[string]struct{})
	edges := make(map[string]struct{})
	for _, triple := range ds.Triples {
		if triple.Predicate.String() == "http://www.w3.org/1999/02/22-rdf-syntax-ns#type" && triple.Object.String() == "http://www.w3.org/2002/07/owl#Class" {
			x := fmt.Sprintf("\"%s\";\n", triple.Subject.Value)
			nodes[x] = struct{}{}
		} else if triple.Predicate.String() == "http://www.w3.org/2000/01/rdf-schema#subClassOf" {
			if strings.HasPrefix(triple.Object.Value, "genid") || strings.HasPrefix(triple.Subject.Value, "genid") {
				continue
			}
			x := fmt.Sprintf("\"%s\" -> \"%s\" [label=\"%s\"];\n", triple.Object.Value, triple.Subject.Value, "hasSubclass")
			edges[x] = struct{}{}
		}
	}

	fmt.Fprintln(f, "digraph G {")
	fmt.Fprintln(f, "ratio=\"auto\"")
	fmt.Fprintln(f, "rankdir=\"LR\"")
	fmt.Fprintln(f, "size=\"7.5,10\"")
	for node := range nodes {
		fmt.Fprintf(f, node)
	}
	for edge := range edges {
		fmt.Fprintf(f, edge)
	}
	fmt.Fprintln(f, "}")
	cmd := exec.Command("dot", "-Tpdf", name)
	pdf, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return err
	}
	f2, err := os.Create(filename + ".pdf")
	if err != nil {
		log.Error(err)
		return err
	}
	_, err = f2.Write(pdf)
	if err != nil {
		log.Error(err)
		return err
	}

	// remove DOT file
	os.Remove(name)
	return nil
}

func dumpGraph(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("Need to specify a turtle file to load")
	}
	filename := c.Args().Get(0)
	p := turtle.GetParser()
	ds, _ := p.Parse(filename)

	name := gethash() + ".gv"
	f, err := os.Create(name)
	if err != nil {
		log.Error(err)
		return err
	}

	nodes := make(map[string]struct{})
	edges := make(map[string]struct{})
	for _, triple := range ds.Triples {
		x := fmt.Sprintf("\"%s\";\n", triple.Subject.Value)
		nodes[x] = struct{}{}
		x = fmt.Sprintf("\"%s\" -> \"%s\" [label=\"%s\"];\n", triple.Subject.Value, triple.Object.Value, triple.Predicate.Value)
		edges[x] = struct{}{}
	}

	fmt.Fprintln(f, "digraph G {")
	fmt.Fprintln(f, "ratio=\"auto\"")
	fmt.Fprintln(f, "rankdir=\"LR\"")
	fmt.Fprintln(f, "size=\"7.5,10\"")
	for node := range nodes {
		fmt.Fprintf(f, node)
	}
	for edge := range edges {
		fmt.Fprintf(f, edge)
	}
	fmt.Fprintln(f, "}")
	cmd := exec.Command("sfdp", "-Tpdf", name)
	pdf, err := cmd.Output()
	if err != nil {
		// try graphviz dot then
		cmd = exec.Command("dot", "-Tpdf", name)
		pdf, err = cmd.Output()
		if err != nil {
			log.Error(err)
			return err
		}
	}
	f2, err := os.Create(filename + ".pdf")
	if err != nil {
		log.Error(err)
		return err
	}
	_, err = f2.Write(pdf)
	if err != nil {
		log.Error(err)
		return err
	}

	// remove DOT file
	os.Remove(name)
	return nil
}

func gethash() string {
	h := md5.New()
	seed := make([]byte, 16)
	binary.PutVarint(seed, time.Now().UnixNano())
	h.Write(seed)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func runInteractiveQuery(db *hod.HodDB, stop chan<- os.Signal) error {
	currentUser, err := user.Current()
	if err != nil {
		log.Error(err)
		return err
	}
	fmt.Println("Successfully loaded dataset!")
	bufQuery := ""

	//setup color for prompt
	c := color.New(color.FgCyan)
	c.Add(color.Bold)
	cyan := c.SprintFunc()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:                 cyan("(hod)> "),
		HistoryFile:            currentUser.HomeDir + "/.hod-query-history",
		DisableAutoSaveHistory: true,
	})
	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			// interactive query blocks receiving shutdown signal, so we send it manually
			stop <- os.Interrupt
			break
		} else if err != nil {
			log.Error(err)
			continue
		}
		if len(line) == 0 {
			continue
		}
		bufQuery += line + " "
		if !strings.HasSuffix(strings.TrimSpace(line), ";") {
			rl.SetPrompt(">>> ...")
			continue
		}
		rl.SetPrompt(cyan("(hod)> "))
		rl.SaveHistory(bufQuery)
		_, err = query.Parse(bufQuery)
		if err != nil {
			log.Error(err)
		} else if res, err := db.RunQueryString(bufQuery); err != nil {
			log.Error(err)
		} else if res != nil {
			hod.DumpQueryResponse(res)
		}
		bufQuery = ""
	}
	return nil
}

type uniqCounter map[string]struct{}

func (uc *uniqCounter) Add(item string) {
	(*uc)[item] = struct{}{}
}
func (uc *uniqCounter) GetCount() float64 {
	return float64(len(*uc))
}

// We load all turtle files into a unified graph and compute
// the following statistics, recording the following
// - in-degree per node
// - out-degree per node
// - # of nodes
// - # of edges
// - # of triples
func ttlStat(c *cli.Context) error {
	if c.NArg() == 0 {
		log.Fatal("Need to provide at least one TTL file")
	}
	numTriples := 0
	uniqueEdges := make(uniqCounter)
	uniqueNodes := make(uniqCounter)
	outdegree := make(map[string]int)
	indegree := make(map[string]int)
	predFreq := make(map[string]int)
	for fileidx := 0; fileidx < c.NArg(); fileidx++ {
		filename := c.Args().Get(fileidx)
		p := turtle.GetParser()
		ds, _ := p.Parse(filename)
		numTriples += ds.NumTriples()
		for _, triple := range ds.Triples {
			uniqueEdges.Add(triple.Predicate.String())
			uniqueNodes.Add(triple.Subject.String())
			uniqueNodes.Add(triple.Object.String())

			if cur, found := outdegree[triple.Subject.String()]; found {
				outdegree[triple.Subject.String()] = cur + 1
			} else {
				outdegree[triple.Subject.String()] = 1
			}
			if cur, found := indegree[triple.Object.String()]; found {
				indegree[triple.Object.String()] = cur + 1
			} else {
				indegree[triple.Object.String()] = 1
			}

			if cur, found := predFreq[triple.Predicate.String()]; found {
				predFreq[triple.Predicate.String()] = cur + 1
			} else {
				predFreq[triple.Predicate.String()] = 1
			}
		}
	}
	// load into arrays so we can do stats
	var outdegreeCounts stats.Float64Data
	for _, count := range outdegree {
		outdegreeCounts = append(outdegreeCounts, float64(count))
	}
	var indegreeCounts stats.Float64Data
	for _, count := range indegree {
		indegreeCounts = append(indegreeCounts, float64(count))
	}
	var predFrequencyCounts stats.Float64Data
	for _, count := range predFreq {
		predFrequencyCounts = append(predFrequencyCounts, float64(count))
	}
	fmt.Printf("# Triples: %d\n", numTriples)
	fmt.Printf("# Unique Nodes: %0.0f\n", uniqueNodes.GetCount())
	fmt.Printf("# Unique Edges: %0.0f\n", uniqueEdges.GetCount())
	// compute density
	// for N nodes in the graph, max edges is N(N-1) for a directed graph with 1 type of edge.
	// For M types of edges, this is M*N*(N-1)
	// density is the number of edges we have out of this theoretical maximum. Each triple corresponds to 1 edge
	maxEdges := big.NewFloat(uniqueNodes.GetCount() * (uniqueNodes.GetCount() - 1) * uniqueEdges.GetCount())
	ntFloat := big.NewFloat(float64(numTriples))
	density := new(big.Float)
	density.Quo(ntFloat, maxEdges)
	fmt.Printf("Density: %s\n", density.String())
	sum_outdeg, _ := outdegreeCounts.Sum()
	min_outdeg, _ := outdegreeCounts.Min()
	max_outdeg, _ := outdegreeCounts.Max()
	mean_outdeg := sum_outdeg / uniqueNodes.GetCount()
	med_outdeg, _ := outdegreeCounts.Median()
	std_outdeg, _ := outdegreeCounts.StandardDeviation()
	fmt.Printf("OutDegree: Min %0.2f, Max %0.2f, Mean %0.2f, Std Dev %0.2f, Median %0.2f, Avg Per Type %0.4f\n", min_outdeg, max_outdeg, mean_outdeg, std_outdeg, med_outdeg, mean_outdeg/uniqueEdges.GetCount())
	sum_indeg, _ := indegreeCounts.Sum()
	min_indeg, _ := indegreeCounts.Min()
	max_indeg, _ := indegreeCounts.Max()
	mean_indeg := sum_indeg / uniqueNodes.GetCount()
	med_indeg, _ := indegreeCounts.Median()
	std_indeg, _ := indegreeCounts.StandardDeviation()
	fmt.Printf("InDegree: Min %0.2f, Max %0.2f, Mean %0.2f, Std Dev %0.2f, Median %0.2f, Avg Per Type %0.4f\n", min_indeg, max_indeg, mean_indeg, std_indeg, med_indeg, mean_indeg/uniqueEdges.GetCount())
	sum_pred, _ := predFrequencyCounts.Sum()
	min_pred, _ := predFrequencyCounts.Min()
	max_pred, _ := predFrequencyCounts.Max()
	mean_pred := sum_pred / uniqueEdges.GetCount()
	std_pred, _ := predFrequencyCounts.StandardDeviation()
	fmt.Printf("Pred Frequencies: Min %0.2f, Max %0.2f, Mean %0.2f, Std Dev %0.2f\n", min_pred, max_pred, mean_pred, std_pred)

	// write the outdegree and predicate data to files; one entry per line
	outdegreefile, err := os.Create("outdegree.csv")
	if err != nil {
		log.Fatal(errors.Wrap(err, "Could not create outdegree file"))
	}
	outdegreecsv := csv.NewWriter(outdegreefile)
	for _, cnt := range outdegreeCounts {
		if err := outdegreecsv.Write([]string{fmt.Sprintf("%d", int64(cnt))}); err != nil {
			log.Fatal(errors.Wrap(err, "Could not write to CSV file"))
		}
	}
	outdegreecsv.Flush()
	// write the indegree and predicate data to files; one entry per line
	indegreefile, err := os.Create("indegree.csv")
	if err != nil {
		log.Fatal(errors.Wrap(err, "Could not create indegree file"))
	}
	indegreecsv := csv.NewWriter(indegreefile)
	for _, cnt := range indegreeCounts {
		if err := indegreecsv.Write([]string{fmt.Sprintf("%d", int64(cnt))}); err != nil {
			log.Fatal(errors.Wrap(err, "Could not write to CSV file"))
		}
	}
	indegreecsv.Flush()
	edgefile, err := os.Create("edge.csv")
	if err != nil {
		log.Fatal(errors.Wrap(err, "Could not create edge file"))
	}
	edgecsv := csv.NewWriter(edgefile)
	for _, cnt := range predFrequencyCounts {
		if err := edgecsv.Write([]string{fmt.Sprintf("%d", int64(cnt))}); err != nil {
			log.Fatal(errors.Wrap(err, "Could not write to CSV file"))
		}
	}
	edgecsv.Flush()
	return nil
}
