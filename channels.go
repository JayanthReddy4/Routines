package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int)
	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sum := calculateSum(n)
			resultChan <- sum
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect and sum the results from the channel
	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}

func calculateSum(num int) int {
	// Simulate some time-consuming work
	// (e.g., by sleeping for a short duration)
	// This is just for demonstration purposes.
	// In a real-world scenario, you might perform a more complex calculation.
	// Remove this sleep in a production code.
	// time.Sleep(time.Millisecond * 100)

	return num
}
