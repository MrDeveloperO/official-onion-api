package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

type FakeStore struct {
}

func (f *FakeStore) Links() []string {
	return []string{"a", "b", "c"}
}

func TestOnion(t *testing.T) {
	_, err := NewOnions("")
	if err == nil {
		t.Error("expected error, got nil")
		return
	}
	o, err := NewOnions("official_links.csv")
	if err != nil {
		t.Error("got error", err)
		return
	}

	if len(o.list) <= 0 {
		t.Error("the list should be populated, got 0")
	}
}

func TestAPI(t *testing.T) {
	srv := NewServer(":8080")
	ctrl := NewController(&FakeStore{})
	srv.Register(ctrl)
	go srv.ListenAndServe()

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Error(err)
		return
	}
	var ret []string
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		t.Error(err)
		return
	}
	if len(ret) <= 0 {
		t.Error("Got 0 results from API")
		return
	}

	//check if we cached the controller
	if len(ctrl.cacheLinks) <= 0 {
		t.Error("controller did not cached urls")
		return
	}

}
