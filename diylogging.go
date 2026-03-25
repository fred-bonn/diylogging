package diylogging

import (
	"fmt"
	"sync"
	"time"
)

func LogStart() (chan<- any, *sync.WaitGroup) {
	ch := make(chan any, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go log(ch, &wg)
	return ch, &wg
}

func log(ch <-chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	for message := range ch {
		switch value := message.(type) {
		case fmt.Stringer:
			fmt.Printf("%s | %s\n", time.Now().Format(time.RFC3339), value.String())
		case string:
			fmt.Printf("%s | %s\n", time.Now().Format(time.RFC3339), value)
		default:
			fmt.Printf("%s | %v\n", time.Now().Format(time.RFC3339), value)

		}
	}
}

func TimerStart() (chan<- any, *sync.WaitGroup) {
	ch := make(chan any, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go timer(ch, &wg)
	return ch, &wg
}

func timer(ch <-chan any, wg *sync.WaitGroup) {
	defer wg.Done()
	var ongoing bool
	var startTime time.Time
	var tempDuration time.Duration
	var totalDuration time.Duration
	globalTime := time.Now()

	for range ch {
		if !ongoing {
			ongoing = true
			startTime = time.Now()
		} else {
			tempDuration = time.Since(startTime)
			totalDuration += tempDuration
			fmt.Printf("Interval: %vns\n", tempDuration.Nanoseconds())
			ongoing = false
		}
	}

	fmt.Printf("Total Time: %v\n", totalDuration)
	fmt.Printf("Ratio: %.3f%%\n", float64(totalDuration.Nanoseconds())/float64(time.Since(globalTime).Nanoseconds())*100)
}
