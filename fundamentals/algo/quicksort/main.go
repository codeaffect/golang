package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := GetInputSize()
	inputSlices := PopulateRandomDataByInputSize(n)

	fmt.Println(inputSlices)

	//QuickSort
	timestart := time.Now()

	h := n - 1
	l := 0

	inputSlices = quicksort(l, h, inputSlices)

	timeend := time.Now()

	diff := timeend.Sub(timestart)
	fmt.Printf("Time taken: %v", diff)
}

func quicksort(l int, h int, c []int) []int {
	if len(c) == 1 || l >= h {
		return c
	}
	fmt.Print("\n\n --------*** -----------\n")
	i := l + 1
	j := h
	p := l
	pvalue := c[l]
	fmt.Println("Pivot:", i-1, " - Low:", i, " -- High:", j)
	fmt.Println("pvalue:", pvalue, " - lv:", c[i], " -- hv:", c[j])

	for i < j {
		fmt.Println("again...")
		for i < j {
			if c[i] > pvalue {
				break
			}
			i++
		}
		fmt.Println("done i", i)
		for j > l+1 {
			if c[j] < pvalue {
				break
			}
			j--
		}
		fmt.Println("done j", j)

		if i >= j {
			break
		}

		fmt.Println("swap", c[i], "at", i, "-- with", c[j], "at", j)
		c[i], c[j] = swap(c[i], c[j])
		printCollection(c)

		fmt.Scanln()
	}

	if c[p] > c[j] {
		fmt.Println("swap pivot ", c[p], "at", p, "-- with", c[j], "at", j)
		c[p], c[j] = swap(c[p], c[j])
		p = j
	}

	printCollection(c)

	if l < p {
		c = quicksort(l, p, c)
	}
	if p+1 < h {
		c = quicksort(p+1, h, c)
	}

	return c
}

func printCollection(c []int) {
	fmt.Println(c)
}

func swap(i int, j int) (int, int) {
	return j, i
}

func PopulateRandomDataByInputSize(inputSize int) []int {
	rand.Seed(time.Now().UnixNano())
	randomPopulation := rand.Perm(inputSize)
	return randomPopulation
}

func GetInputSize() int {
	var inputSize int
	fmt.Print("Enter input size:")
	fmt.Scan(&inputSize)
	return inputSize
}
