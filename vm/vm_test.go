package vm

import "testing"

func Benchmark_VM(b *testing.B) {
	vm := NewVM(1000)
	vm.Load([]uint32{3, 4, 0x40000001, 4, 0x40000001, 0x40000000})

	b.Run("run", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			vm.Run(1000)
		}
	})
}
