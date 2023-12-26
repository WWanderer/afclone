package common

type formError string

type Control[T any] struct {
	Value    T           `bson:"value"`
	Errors   []formError `bson:"errors"`
	Enabled  bool        `bson:"enabled"`
	Readonly bool        `bson:"readonly"`
}
