package general

import (
	"log"
	"time"
)

func StartTimer(s string) (string, time.Time) {
	//log.Println("Start:	", s)
	return s, time.Now()
}

func FinishTimer(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:	", s, "took", endTime.Sub(startTime))
}

func execute() {
	defer FinishTimer(StartTimer("execute"))
	time.Sleep(3 * time.Second)
}