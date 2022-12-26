package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Expenses struct {
	Id     int      `json:"id"`
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func TestGetAll(t *testing.T) {
	var e []Expenses
	res := request(http.MethodGet, uri("expenses"), nil)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(e), 0)

}

func seedExpenses(t *testing.T) Expenses {
	var c Expenses
	body := bytes.NewBufferString(`{
		"title": "buy a new phone",
		"amount": 39000,
		"note": "buy a new phone",
		"tags": ["gadget"]
		}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create expense:", err)
	}
	return c
}

func TestGetById(t *testing.T) {
	c := seedExpenses(t)
	var latest Expenses
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(c.Id)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.Id, latest.Id)
	assert.NotEmpty(t, latest.Title)
	assert.NotEmpty(t, latest.Amount)
	assert.NotEmpty(t, latest.Note)
	assert.NotEmpty(t, latest.Tags)
}

func TestCreate(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "test",
		"amount": 100,
		"note": "teste",
		"tags": ["gadget"]
		}`)
	var e Expenses

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&e)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, e.Id)
	assert.Equal(t, "test", e.Title)
	assert.Equal(t, 100.00, e.Amount)
	assert.Equal(t, "teste", e.Note)
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}
	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", "November 10, 2009")
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}
