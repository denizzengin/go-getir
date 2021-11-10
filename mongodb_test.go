package main

import (
	"testing"
)

func TestMongoDbConnection(t *testing.T) {
	s := "2021-11-01"
	e := "2021-12-10"
	q := MongoHandlerRequest{StartDate: s, EndDate: e, MinCount: 0, MaxCount: 10}
	_, err := SearchDbQuery(q)
	if err != nil {
		t.Fail()
	}
}
