package qntfy
import (
	"testing"
)

func BenchmarkReadFiles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readFilesFromDirectory("./files")
	}
}
