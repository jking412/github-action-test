package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	time.Sleep(2 * time.Second)
	m.Run()
}

func TestAdd(t *testing.T) {
	assertion := assert.New(t)
	resp, err := http.Get("http://localhost:8080/hello")
	assertion.Nil(err)
	assertion.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assertion.Equal("hello world", string(data))
}
