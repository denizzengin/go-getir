package main

import (
	"sync"
	"time"
)

// MongoHandlerRequest : handler request model
type MongoHandlerRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int64  `json:"minCount"`
	MaxCount  int64  `json:"maxCount"`
}

// RecordsModel : nested model
type RecordsModel struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int64     `json:"totalCount"`
}

// MongoHandlerResponse : handler response model
type MongoHandlerResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Records []RecordsModel `json:"records"`
}

// StoreKeyValuePairRequest : handler Request
type StoreKeyValuePairRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// InMemoryMap : data struct
type InMemoryMap struct {
	KeyValuePair map[string]string
	Mutex        *sync.Mutex
}
