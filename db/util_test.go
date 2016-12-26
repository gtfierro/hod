package db

import (
	"github.com/google/btree"
	"github.com/gtfierro/bloom"
	"testing"
)

var intersectTreesBtree = func(a, b *btree.BTree) *btree.BTree {
	if a.Len() < b.Len() {
		a, b = b, a
	}
	res := btree.New(3)
	if a.Max().Less(b.Min()) {
		return res
	}
	iter := func(i btree.Item) bool {
		if b.Has(i) {
			res.ReplaceOrInsert(i)
		}
		return i != a.Max()
	}
	a.Ascend(iter)
	return res
}

var intersectItemtrees = func(a, b *itemtree) *itemtree {
	if a.IntersectionCardinality(b) == 0 {
		return newItemTree(3, true)
	}
	if a.Len() < b.Len() {
		a, b = b, a
	}
	res := newItemTree(3, true)
	iter := func(i btree.Item) bool {
		if b.Has(i) {
			res.ReplaceOrInsert(i)
		}
		return i != a.Max()
	}
	a.Ascend(iter)
	return res
}

func BenchmarkInsertTree100(b *testing.B) {
	trees := make([]*btree.BTree, b.N)
	for i := 0; i < b.N; i++ {
		trees[i] = btree.New(3)
	}
	for i := 0; i < b.N; i++ {
		bits := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		trees[i].ReplaceOrInsert(Item(bits))
	}
}

func BenchmarkInsertItemTree100(b *testing.B) {
	trees := make([]*itemtree, b.N)
	for i := 0; i < b.N; i++ {
		trees[i] = newItemTree(3, true)
	}
	for i := 0; i < b.N; i++ {
		bits := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		trees[i].ReplaceOrInsert(Item(bits))
	}
}

func BenchmarkIntersectTreesBtree50(b *testing.B) {
	A := btree.New(3)
	B := btree.New(3)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 50), byte(i + 50 + 1), byte(i + 50 + 2), byte(i + 50 + 3)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectTreesBtree(A, B)
	}
}

func BenchmarkIntersectTreesBtree01(b *testing.B) {
	A := btree.New(3)
	B := btree.New(3)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 99), byte(i + 99 + 1), byte(i + 99 + 2), byte(i + 99 + 3)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectTreesBtree(A, B)
	}
}

func BenchmarkIntersectTreesBtreeAll(b *testing.B) {
	A := btree.New(3)
	B := btree.New(3)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectTreesBtree(A, B)
	}
}

func BenchmarkIntersectTreesBtreeNone(b *testing.B) {
	A := btree.New(3)
	B := btree.New(3)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 100), byte(i + 101), byte(i + 102), byte(i + 103)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectTreesBtree(A, B)
	}
}

func BenchmarkIntersectTreesItemtree50(b *testing.B) {
	A := newItemTree(3, true)
	B := newItemTree(3, true)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 50), byte(i + 50 + 1), byte(i + 50 + 2), byte(i + 50 + 3)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectItemtrees(A, B)
	}
}

func BenchmarkIntersectTreesItemtree01(b *testing.B) {
	A := newItemTree(3, true)
	B := newItemTree(3, true)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 99), byte(i + 99 + 1), byte(i + 99 + 2), byte(i + 99 + 3)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectItemtrees(A, B)
	}
}

func BenchmarkIntersectTreesItemtreeAll(b *testing.B) {
	A := newItemTree(3, true)
	B := newItemTree(3, true)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectItemtrees(A, B)
	}
}

func BenchmarkIntersectTreesItemtreeNone(b *testing.B) {
	A := newItemTree(3, true)
	B := newItemTree(3, true)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 100), byte(i + 101), byte(i + 102), byte(i + 103)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		intersectItemtrees(A, B)
	}
}

func BenchmarkBloomAdd(b *testing.B) {
	f := bloom.NewWithEstimates(1000, .001)
	for i := 0; i < b.N; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		f.Add(bitsa[:])
	}
}

func TestItemTreeIntersect(t *testing.T) {
	A := newItemTree(3, true)
	B := newItemTree(3, true)
	for i := 0; i < 100; i++ {
		bitsa := [4]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3)}
		bitsb := [4]byte{byte(i + 100), byte(i + 101), byte(i + 102), byte(i + 103)}
		A.ReplaceOrInsert(Item(bitsa))
		B.ReplaceOrInsert(Item(bitsb))
	}
	intersectItemtrees(A, B)
}

func BenchmarkBTreeHas3(b *testing.B) {
	t := btree.New(3)
	e := Item([4]byte{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		t.Has(e)
	}
}

func BenchmarkBTreeInsertDuplicate3(b *testing.B) {
	t := btree.New(3)
	e := Item([4]byte{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		t.ReplaceOrInsert(e)
	}
}

func BenchmarkBTreeInsertDuplicateWithHas3(b *testing.B) {
	t := btree.New(3)
	e := Item([4]byte{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		if !t.Has(e) {
			t.ReplaceOrInsert(e)
		}
	}
}

func BenchmarkBTreeHas2(b *testing.B) {
	t := btree.New(3)
	e := Item([4]byte{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		t.Has(e)
	}
}

func BenchmarkBTreeInsertDuplicate2(b *testing.B) {
	t := btree.New(3)
	e := Item([4]byte{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		t.ReplaceOrInsert(e)
	}
}

func BenchmarkBTreeInsertDuplicateWithHas2(b *testing.B) {
	t := btree.New(3)
	e := Item([4]byte{1, 2, 3, 4})
	for i := 0; i < b.N; i++ {
		if !t.Has(e) {
			t.ReplaceOrInsert(e)
		}
	}
}