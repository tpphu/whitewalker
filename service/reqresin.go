package service

import (
	"github.com/imroc/req"
)

type ReqResIn struct {
	BaseURL string
}

func NewReqResIn(baseURL string) Service {
	service := ReqResIn{
		BaseURL: baseURL,
	}
	return service
}

// Get returns the respone of http request
func (r ReqResIn) Get(path string, v ...interface{}) (*req.Resp, error) {
	url := r.BaseURL + path
	return req.Get(url, v...)
}

// Post returns the respone of http request
func (r ReqResIn) Post(path string, v ...interface{}) (*req.Resp, error) {
	url := r.BaseURL + path
	return req.Post(url, v...)
}
