package day2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takeRating(studentId int, ratingArr []int) {
	randomRating := rand.Intn(10) + 1
	ratingArr[studentId] = randomRating

	randomMilliSeconds := rand.Intn(1000) + 1
	time.Sleep(time.Millisecond * time.Duration(randomMilliSeconds))
}

func findAvgRating(ratings []int) float32 {
	sum := 0

	for _, val := range ratings {
		sum += val
	}

	return float32(sum) / float32(len(ratings))
}

func Ex2() {
	totalStudents := 200
	ratings := make([]int, totalStudents)
	wg := sync.WaitGroup{}

	for i := 0; i < totalStudents; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			takeRating(i, ratings)
		}()
	}
	wg.Wait()

	fmt.Printf("Avg rating : %f\n", findAvgRating(ratings))
}
