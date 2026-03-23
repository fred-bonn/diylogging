package diylogging

import (
	"fmt"
	"time"
)

func LogStart(ch chan Stringer) {
	for msg := range ch {
		fmt.Println("%s: %s", time.Now().Format(time.RFC3339), msg.String())
	}
	fmt.Println("Logging stopped.")
}