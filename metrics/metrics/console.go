package metrics

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ConsoleReporter struct {
	storage Storager
	stopSig chan interface{}
}

func NewConsoleReporter(s Storager) *ConsoleReporter {
	return &ConsoleReporter{
		storage: s,
		stopSig: make(chan interface{}),
	}
}

func (r *ConsoleReporter) Start(durationInSeconds int64) {
	t := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-t.C:
			r.report(durationInSeconds)
		case <-r.stopSig:
			t.Stop()
			return
		}
	}
}

func (r *ConsoleReporter) Stop() {
	r.stopSig <- 1
	close(r.stopSig)
}

func (r *ConsoleReporter) report(durationInSeconds int64) {
	// 1.query
	durationInMillis := durationInSeconds * 1000
	endTimeInMillis := time.Now().Unix() * 1000
	startTimeInMillis := endTimeInMillis - durationInMillis
	infoMap, err := r.storage.GetRequestInfoMap(startTimeInMillis, endTimeInMillis)
	if err != nil {
		log.Println(err)
		return
	}
	statMap := make(map[string]*RequestStat)
	for k, v := range infoMap {
		// 2.aggregate
		stat := Aggregate(v, durationInMillis)
		statMap[k] = stat
	}
	// 3.print
	fmt.Printf("Time Span: [%d, %d]\n", startTimeInMillis, endTimeInMillis)
	b, err := json.Marshal(statMap)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}
