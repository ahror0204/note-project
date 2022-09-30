package repo

type RedisRepositoryStorage interface {
	Set(key, value string) error
	SetWithTTL(key, value string, second int64) error 
	Update(key, value string, second int64) error
	Get(key string) (interface{}, error)
	DeleteByKey(key string) error
	SetTTLByKey(key string, second int64) error
}