package main

import "fmt"

func main() {
	fmt.Println(Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))
}

func Chunk(a []int, size int) (chunks []int) {
	var length int
	for length = range a {
		length++
	}
	chunks = make([]int, length/size)
	for i := range a {
		Chunk := make([]int, size)
		for i = 0; i < size; i++ {
			Chunk = append(Chunk, a[i])
			a = a[size:]
		}
		chunks = append(chunks, Chunk...)

	}
	return
}
