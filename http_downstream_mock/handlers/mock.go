package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kuangwanjing/test_in_golang/http_downstream_mock/models"
)

// put down the normal business logic here, don't have to consider mocking here
func (h *handlers) MockHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := h.client.Get("https://example.com")
	if err != nil {
		w.Write([]byte("error fetching from downstream service"))
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte("error reading response body from downstream service"))
		return
	}
	defer resp.Body.Close()

	var example models.Example
	err = json.Unmarshal(body, &example)
	if err != nil {
		w.Write([]byte("error decoding response from downstream service"))
		return
	}

	w.Write([]byte(example.Name))
}

type MockClient struct{}

// Get - This is a mock client method
func (client MockClient) Get(query string) (*http.Response, error) {
	example := &models.Example{Name: "hello wold"}
	body, _ := json.Marshal(example)
	t := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(string(body))),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	return t, nil
}
