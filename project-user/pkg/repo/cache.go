package repo

import (
	"context"
	"time"
)

type Cache interface {
	//ctx 这个变量的作用是上面要去好好看一下啊
	Put(ctx context.Context, key, value string, expire time.Duration) error

	Get(ctx context.Context, key string) (string, error)
}
