package diylogging

import (
	"fmt"
	"time"
)

func LogStart(ch chan string) {
	for msg := range ch {
		fmt.Println("%s: %s", time.Now().Format(time.RFC3339), msg)
	}
	fmt.Println("Logging stopped.")
}