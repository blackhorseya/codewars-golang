package kata

func DNAStrand(dna string) string {
	ret := ""
	for _, letter := range dna {
		switch letter {
		case 'A':
			ret += "T"
		case 'T':
			ret += "A"
		case 'C':
			ret += "G"
		case 'G':
			ret += "C"
		default:
			ret += string(letter)
		}
	}

	return ret
}
