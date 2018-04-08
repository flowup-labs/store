package store

type OrderFlow int

// OrderFlow constants
const (
	OrderFlowAscending  OrderFlow = 1
	OrderflowDescending           = 2
)

type Query interface {
	Where(field string, op string, val interface{}) Query
	Order(field string, flow OrderFlow) Query
	Project(fieldNames ...string) Query
	Offset(offset int) Query
	Limit(limit int) Query
}

type QueryOption func(Query) Query

type Transaction interface {
}
