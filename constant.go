package main

import (
	"errors"
)

// ErrValidation : Bad request
var ErrValidation = errors.New("please control your date inputs, it must be send like, YYYY-MM-DD")

// ErrNotFound : Not found error
var ErrNotFound = errors.New("key not found in memory map")

// ErrMongoDbConnection : Connection refused vs..
var ErrMongoDbConnection = errors.New("cannot connect mongodb")

// DateTemplate : for validation
// MongoDbAtlassURL : connection string
const (
	DateTemplate          string = "2006-01-02"
	MongoDbAtlassURL      string = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"
	MongoDbLocalURL       string = "mongodb://127.0.0.1:27017"
	NotFoundID            string = "id is missing in parameters"
	MongoDbName           string = "getircase-study"
	MongoDbCollectionName string = "records"
	SuccessMessage        string = "success"
	RecordNotFoundMessage string = "cannot find any record"
)

type responseCodeEnum struct {
	Success        int32
	Fail           int32
	RecordNotFound int32
}

func newResponseCodeRegistry() *responseCodeEnum {
	return &responseCodeEnum{
		Success:        0,
		Fail:           1,
		RecordNotFound: 2,
	}
}
