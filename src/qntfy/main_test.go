package qntfy

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
