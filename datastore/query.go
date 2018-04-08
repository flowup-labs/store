package datastore

import (
	ds "cloud.google.com/go/datastore"
	"github.com/flowup-labs/store"
)

type Query struct {
	*ds.Query
}

func (q *Query) Raw() interface{} {
	return q.Query
}

func (q *Query) Where(field string, op string, val interface{}) store.Query {
	q.Query = q.Query.Filter(field+" "+op, val)

	return q
}

func (q *Query) Order(field string, flow store.OrderFlow) store.Query {
	switch flow {
	case store.OrderflowDescending:
		q.Query = q.Query.Order(field)
	case store.OrderFlowAscending:
		q.Query = q.Query.Order("-" + field)
	}

	return q
}

func (q *Query) Project(fieldNames ...string) store.Query {
	q.Query = q.Query.Project(fieldNames...)

	return q
}

func (q *Query) Offset(offset int) store.Query {
	q.Query = q.Query.Offset(offset)

	return q
}

func (q *Query) Limit(limit int) store.Query {
	q.Query = q.Query.Limit(limit)

	return q
}

func (q *Query) KeysOnly() store.Query {
	q.Query = q.Query.KeysOnly()

	return q
}
