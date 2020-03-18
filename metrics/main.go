package main

import "one-more-step-in-dp/metrics/metrics"

func main() {
	var durationInSeconds int64 = 10
	var spec string = "1 * * * *"

	storage := metrics.RedisStorage{}
	reporter := metrics.NewConsoleReporter(&storage)
	reporter.Start(durationInSeconds, spec)
}
