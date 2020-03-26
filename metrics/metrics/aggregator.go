package metrics

import (
	"math"
	"sort"
)

type Aggregator struct {
}

type RequestStat struct {
	MaxResponseTime  float64
	MinResponseTime  float64
	AvgResponseTime  float64
	P999ResponseTime float64
	P99ResponseTime  float64
	Count            int64
	TPS              int64
}

func Aggregate(requestInfos []*RequestInfo, durationInMillis int64) *RequestStat {
	max, min := float64(0), math.MaxFloat64
	avg, p999, p99, sum := float64(-1), float64(-1), float64(-1), float64(0)
	var count int64

	for _, info := range requestInfos {
		count++
		resp := info.ResponseTime
		if max < resp {
			max = resp
		}
		if min > resp {
			min = resp
		}
		sum += resp
	}
	if count != 0 {
		avg = sum / float64(count)
	}
	tps := count / durationInMillis * 1000
	sort.Slice(requestInfos, func(i, j int) bool {
		return requestInfos[i].ResponseTime < requestInfos[j].ResponseTime
	})
	idx999 := int64(float64(count) * 0.999)
	idx99 := int64(float64(count) * 0.99)
	if count != 0 {
		p999 = requestInfos[idx999].ResponseTime
		p99 = requestInfos[idx99].ResponseTime
	}
	return &RequestStat{
		MaxResponseTime:  max,
		MinResponseTime:  min,
		AvgResponseTime:  avg,
		P999ResponseTime: p999,
		P99ResponseTime:  p99,
		Count:            count,
		TPS:              tps,
	}
}
