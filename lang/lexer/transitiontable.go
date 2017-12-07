// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 40: // ['(','(']
			return 2
		case r == 41: // [')',')']
			return 3
		case r == 42: // ['*','*']
			return 4
		case r == 43: // ['+','+']
			return 5
		case r == 46: // ['.','.']
			return 6
		case r == 47: // ['/','/']
			return 7
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 59: // [';',';']
			return 9
		case r == 60: // ['<','<']
			return 10
		case r == 63: // ['?','?']
			return 11
		case 65 <= r && r <= 66: // ['A','B']
			return 12
		case r == 67: // ['C','C']
			return 13
		case 68 <= r && r <= 69: // ['D','E']
			return 12
		case r == 70: // ['F','F']
			return 14
		case 71 <= r && r <= 82: // ['G','R']
			return 12
		case r == 83: // ['S','S']
			return 15
		case r == 84: // ['T','T']
			return 12
		case r == 85: // ['U','U']
			return 16
		case r == 86: // ['V','V']
			return 12
		case r == 87: // ['W','W']
			return 17
		case 88 <= r && r <= 90: // ['X','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case r == 97: // ['a','a']
			return 19
		case 98 <= r && r <= 122: // ['b','z']
			return 20
		case r == 123: // ['{','{']
			return 21
		case r == 124: // ['|','|']
			return 22
		case r == 125: // ['}','}']
			return 23
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 62: // ['>','>']
			return 25
		default:
			return 10
		}
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case r == 95: // ['_','_']
			return 28
		case 97 <= r && r <= 122: // ['a','z']
			return 29
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 78: // ['A','N']
			return 12
		case r == 79: // ['O','O']
			return 30
		case 80 <= r && r <= 90: // ['P','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 81: // ['A','Q']
			return 12
		case r == 82: // ['R','R']
			return 31
		case 83 <= r && r <= 90: // ['S','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 68: // ['A','D']
			return 12
		case r == 69: // ['E','E']
			return 32
		case 70 <= r && r <= 90: // ['F','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 77: // ['A','M']
			return 12
		case r == 78: // ['N','N']
			return 33
		case 79 <= r && r <= 90: // ['O','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 71: // ['A','G']
			return 12
		case r == 72: // ['H','H']
			return 34
		case 73 <= r && r <= 90: // ['I','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 37
		case 97 <= r && r <= 122: // ['a','z']
			return 38
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case r == 95: // ['_','_']
			return 28
		case 97 <= r && r <= 122: // ['a','z']
			return 29
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case r == 95: // ['_','_']
			return 28
		case 97 <= r && r <= 122: // ['a','z']
			return 29
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case r == 95: // ['_','_']
			return 28
		case 97 <= r && r <= 122: // ['a','z']
			return 29
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case 65 <= r && r <= 90: // ['A','Z']
			return 27
		case r == 95: // ['_','_']
			return 28
		case 97 <= r && r <= 122: // ['a','z']
			return 29
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 84: // ['A','T']
			return 12
		case r == 85: // ['U','U']
			return 39
		case 86 <= r && r <= 90: // ['V','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 78: // ['A','N']
			return 12
		case r == 79: // ['O','O']
			return 40
		case 80 <= r && r <= 90: // ['P','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 75: // ['A','K']
			return 12
		case r == 76: // ['L','L']
			return 41
		case 77 <= r && r <= 90: // ['M','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 72: // ['A','H']
			return 12
		case r == 73: // ['I','I']
			return 42
		case 74 <= r && r <= 90: // ['J','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 68: // ['A','D']
			return 12
		case r == 69: // ['E','E']
			return 43
		case 70 <= r && r <= 90: // ['F','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 37
		case 97 <= r && r <= 122: // ['a','z']
			return 38
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 37
		case 97 <= r && r <= 122: // ['a','z']
			return 38
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 37
		case 97 <= r && r <= 122: // ['a','z']
			return 38
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 35
		case 65 <= r && r <= 90: // ['A','Z']
			return 36
		case r == 95: // ['_','_']
			return 37
		case 97 <= r && r <= 122: // ['a','z']
			return 38
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 77: // ['A','M']
			return 12
		case r == 78: // ['N','N']
			return 44
		case 79 <= r && r <= 90: // ['O','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 76: // ['A','L']
			return 12
		case r == 77: // ['M','M']
			return 45
		case 78 <= r && r <= 90: // ['N','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 68: // ['A','D']
			return 12
		case r == 69: // ['E','E']
			return 46
		case 70 <= r && r <= 90: // ['F','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 78: // ['A','N']
			return 12
		case r == 79: // ['O','O']
			return 47
		case 80 <= r && r <= 90: // ['P','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 81: // ['A','Q']
			return 12
		case r == 82: // ['R','R']
			return 48
		case 83 <= r && r <= 90: // ['S','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 83: // ['A','S']
			return 12
		case r == 84: // ['T','T']
			return 49
		case 85 <= r && r <= 90: // ['U','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 66: // ['A','B']
			return 12
		case r == 67: // ['C','C']
			return 50
		case 68 <= r && r <= 90: // ['D','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 77: // ['A','M']
			return 12
		case r == 78: // ['N','N']
			return 51
		case 79 <= r && r <= 90: // ['O','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 68: // ['A','D']
			return 12
		case r == 69: // ['E','E']
			return 52
		case 70 <= r && r <= 90: // ['F','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 83: // ['A','S']
			return 12
		case r == 84: // ['T','T']
			return 53
		case 85 <= r && r <= 90: // ['U','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 8
		case r == 58: // [':',':']
			return 24
		case 65 <= r && r <= 90: // ['A','Z']
			return 12
		case r == 95: // ['_','_']
			return 18
		case 97 <= r && r <= 122: // ['a','z']
			return 20
		}
		return NoState
	},
}