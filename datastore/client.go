package datastore

import (
	"context"
	"reflect"
	"strings"

	"github.com/flowup-labs/store"

	"cloud.google.com/go/datastore"
)

type Client struct {
	*datastore.Client
}

func kind(from interface{}) string {
	return strings.Split(reflect.TypeOf(from).String(), ".")[1]
}

func (c *Client) Get(ctx context.Context, model interface{}, opts ...store.QueryOption) error {
	raw := datastore.NewQuery(kind(model))
	var q store.Query = &Query{raw}

	// apply options to the query
	for _, opt := range opts {
		q = opt(q)
	}
	q = q.Limit(1).KeysOnly()

	keys, err := c.Client.GetAll(ctx, q.Raw().(*datastore.Query), nil)
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		return c.Client.Get(ctx, keys[0], model)
	}

	return nil
}

func (c *Client) GetMulti(ctx context.Context, model interface{}, opts ...store.QueryOption) error {
	return nil
}

func (c *Client) Count(ctx context.Context, opts ...store.QueryOption) (int, error) {
	return 0, nil
}
