package diylogging

import (
	"fmt"
	"time"
	"sync"
)

type LogString string

func (ls LogString) String() string {
	return string(ls)
}

func LogStart() (chan<- fmt.Stringer, *sync.WaitGroup) {
	ch := make(chan fmt.Stringer)
	var wg sync.WaitGroup
	wg.Add(1)
	go log(ch, &wg)
	return ch, &wg
}

func log(ch <-chan fmt.Stringer, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("%s | %s\n", time.Now().Format(time.RFC3339), msg.String())
	}
	fmt.Println("Logging stopped.")
}