package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Controller is our handler
type Controller struct {
	store      Store
	cacheLinks []byte
}

//IController interface
type IController interface {
	Links(w http.ResponseWriter, r *http.Request)
}

//Store is our data model
type Store interface {
	Links() []string
}

//NewController creates a new controller
func NewController(store Store) *Controller {
	return &Controller{
		store: store,
	}
}

//Links returns JSON array of all official Tor links
func (c *Controller) Links(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(c.cacheLinks) > 0 {
		w.Write(c.cacheLinks)
		return
	}
	links := c.store.Links()
	buff, err := json.Marshal(links)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.cacheLinks = buff
	w.Write(c.cacheLinks)

}
