package pool

import "github.com/hamed-yousefi/gowl"

func GetPool(size int) gowl.Pool {
	return gowl.NewPool(size)
}
