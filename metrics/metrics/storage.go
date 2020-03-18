package metrics

type Storager interface {
	SaveRequestInfo(r *RequestInfo) error
	GetRequestInfos(apiName string, startTimeInMillis, endTimeInMillis int64) ([]*RequestInfo, error)
	GetRequestInfoMap(startTimeInMillis, endTimeInMillis int64) (map[string][]*RequestInfo, error)
}

type RedisStorage struct {
}

func (s *RedisStorage) SaveRequestInfo(r *RequestInfo) error {
	return nil
}

func (s *RedisStorage) GetRequestInfos(apiName string, startTimeInMillis, endTimeInMillis int64) ([]*RequestInfo, error) {
	return nil, nil
}

func (s *RedisStorage) GetRequestInfoMap(startTimeInMillis, endTimeInMillis int64) (map[string][]*RequestInfo, error) {
	return nil, nil
}
