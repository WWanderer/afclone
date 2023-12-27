package common

type formError string

type Control[T any] struct {
	Value    T           `bson:"value"`
	Errors   []formError `bson:"errors"`
	Disabled bool        `bson:"disabled"`
	Readonly bool        `bson:"readonly"`
	Section  bool        `bson:"section"`
}
