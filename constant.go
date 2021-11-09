package main

import (
	"errors"
)

// ErrValidation : Bad request
var ErrValidation = errors.New("Please control your date inputs, it must be send like : YYYY-MM-DD")

// ErrNotFound : Not found error
var ErrNotFound = errors.New("Key not found in memory map")

// DateTemplate : for validation
// MongoDbAtlassURL : connection string
const (
	DateTemplate     string = "2006-01-02"
	MongoDbAtlassURL string = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"
	NotFoundID       string = "id is missing in parameters"
)
