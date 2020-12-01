package kata

import "strings"

var replacer *strings.Replacer = strings.NewReplacer("A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

func DNAStrand(dna string) string {
	return replacer.Replace(dna)
}
