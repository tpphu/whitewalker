package service

import "github.com/imroc/req"

// Service is interface to get/post
type Service interface {
	Get(path string, v ...interface{}) (*req.Resp, error)
	Post(url string, v ...interface{}) (*req.Resp, error)
}
