package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func setInMemoryHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var input StoreKeyValuePairRequest
	json.Unmarshal(reqBody, &input)
	Set(input.Key, input.Value)
	w.WriteHeader(http.StatusCreated)
	s, _ := json.Marshal(input)
	w.Write(s)
}

func getInMemoryHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/in-memory/get/")
	if key == "" {
		fmt.Fprintf(w, "%+v", NotFoundID)
		return
	}
	if key = strings.Replace(key, "key=", "", 1); key == "" {
		fmt.Fprintf(w, "%+v", NotFoundID)
		return
	}
	pair, err := Get(key)
	if err != nil {
		fmt.Fprintf(w, "%+v", string(err.Error()))
	} else {
		model := StoreKeyValuePairRequest{Key: key, Value: pair}
		json.NewEncoder(w).Encode(model)
	}
}

func validateDateFormat(fields ...string) error {
	for i := 0; i < len(fields); i++ {
		_, e := time.Parse(DateTemplate, fields[i])
		if e != nil {
			return ErrValidation
		}
	}
	return nil
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var input MongoHandlerRequest
	json.Unmarshal(reqBody, &input)
	e := validateDateFormat(input.StartDate, input.EndDate)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(e.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal(input)
	w.Write(s)
}

func logger(w http.ResponseWriter, r *http.Request) {
	// Log request.
	c, e := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	if e != nil {
		panic(e)
	}
	log.Printf("%s\t\t%s\t\t%s", r.Method, r.RequestURI, c)

	switch {
	case r.Method == http.MethodGet && strings.TrimPrefix(r.URL.Path, "/in-memory/get/") != "":
		getInMemoryHandler(w, r)
	case r.Method == http.MethodPost && strings.TrimPrefix(r.URL.Path, "/in-memory/set/") == "":
		setInMemoryHandler(w, r)
	case r.Method == http.MethodPost && strings.TrimPrefix(r.URL.Path, "/mongo/search") == "":
		searchHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func handleRequests() {
	http.HandleFunc("/", logger)
	e := http.ListenAndServe(":8085", nil)
	log.Fatal(e)
}

func main() {
	fmt.Println("Waiting request...")
	customMap = &InMemoryMap{KeyValuePair: make(map[string]string), Mutex: &sync.Mutex{}}
	handleRequests()
}
