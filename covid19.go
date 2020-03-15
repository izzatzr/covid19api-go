package covid19

import (
	resty "github.com/go-resty/resty"
)

type options struct {
	url string
	req *resty.Request
}

// type res interface{}

// New initialization
func New() options {
	api := options{
		"https://covid19.mathdro.id/api",
		resty.New().DisableTrace().R(),
	}
	return api
}
