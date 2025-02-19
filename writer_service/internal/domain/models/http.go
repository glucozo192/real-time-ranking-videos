package models

import "google.golang.org/grpc/codes"

type HttpError struct {
	StatusCode int         `json:"statusCode" yaml:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message" yaml:"message"`
	RootErr    error       `json:"-"`
	Code       string      `json:"code" yaml:"code"`
	GrpcCode   codes.Code  `json:"-"`
}
