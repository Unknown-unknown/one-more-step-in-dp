package main

import (
	"metrics/metrics"
)

func main() {
	var durationInSeconds int64 = 10

	storage := metrics.RedisStorage{}
	cr := metrics.NewConsoleReporter(&storage)
	cr.Start(durationInSeconds)

	er := metrics.NewEmailReporter(&storage, &metrics.EmailSender{})
	er.StartDailyReport("0 1 * * * *")

	collector := metrics.NewCollector(&storage)
	collector.RecordRequest(&metrics.RequestInfo{"register", 123, 1584596168})
	collector.RecordRequest(&metrics.RequestInfo{"register", 223, 1584696268})
	collector.RecordRequest(&metrics.RequestInfo{"register", 323, 158478338})
	collector.RecordRequest(&metrics.RequestInfo{"login", 23, 1584797168})
	collector.RecordRequest(&metrics.RequestInfo{"login", 1223, 1584796468})
}
