package cache

type cache interface {
	Set(key, val interface{}) error
	Get(key interface{}) (interface{}, error)
}
