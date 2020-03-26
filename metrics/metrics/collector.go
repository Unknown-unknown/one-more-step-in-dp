package metrics

import (
	"errors"
)

var (
	ErrInvalidRequestInfo = errors.New("invalid request info")
)

type Collector struct {
	storage Storager
}

func NewCollector(storage Storager) *Collector {
	return &Collector{
		storage: storage,
	}
}

func (c *Collector) RecordRequest(requestInfo *RequestInfo) error {
	if requestInfo == nil || requestInfo.GetAPIName() == "" {
		return ErrInvalidRequestInfo
	}
	return c.storage.SaveRequestInfo(requestInfo)
}

type RequestInfo struct {
	ApiName      string
	ResponseTime float64
	Timestamp    int64
}

func (r *RequestInfo) GetAPIName() string {
	return r.ApiName
}
