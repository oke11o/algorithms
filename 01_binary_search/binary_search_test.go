package binary_search

import (
	"fmt"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	c := 1000
	list := make([]int, c)
	for i := 0; i < c; i++ {
		list[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(list, 299)
		search(list, 1)
	}
}

func BenchmarkSearch2(b *testing.B) {
	c := 1000
	list := make([]int, c)
	for i := 0; i < c; i++ {
		list[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search2(list, 1)
	}
}

func ExampleSearch() {
	c := 1000
	list := make([]int, c)
	for i := 0; i < c; i++ {
		list[i] = i + 1
	}

	index, count, ok := search(list, 1)
	fmt.Println(ok)
	fmt.Println(count)
	fmt.Println(index)

	index, count, ok = search2(list, 1)
	fmt.Println(ok)
	fmt.Println(count)
	fmt.Println(index)

	// Output:
	// true
	// 9
	// 0
	// true
	// 1000
	// 0
}
