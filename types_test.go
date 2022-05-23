package vm

import "testing"

type Smth [32]byte

func BenchmarkCollection(b *testing.B) {
	b.Run("List", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var list List[Smth]
			for i := 0; i < 1000; i++ {
				list.Add(Smth{})
			}
		}
	})
	b.Run("Slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var slice []Smth
			for i := 0; i < 1000; i++ {
				slice = append(slice, Smth{})
			}
		}
	})
}
