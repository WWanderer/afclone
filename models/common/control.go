package common

type formError string

type Control [T any] struct {
	Value T
	Errors []formError
	Enabled bool
	Readonly bool
}