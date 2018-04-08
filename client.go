package store

import "context"

type Client interface {
	Get(ctx context.Context, m interface{}, opts ...QueryOption) error
	GetMulti(ctx context.Context, m interface{}, opts ...QueryOption) error
}
