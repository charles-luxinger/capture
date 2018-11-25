package main

import (
	"net/http"
	"strconv"
	"sync"
)

var captureID int
var captures CaptureList

type CaptureRepository interface {
	Insert(capture Capture)
	RemoveAll()
	Find(captureID string) *Capture
	FindAll() []Capture
}

type CaptureList struct {
	items    []Capture
	mux      sync.Mutex
	maxItems int
}

type Capture struct {
	ID  int
	Req *http.Request
	Res *http.Response
}

type CaptureMetadata struct {
	ID     int    `json:"id"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Status int    `json:"status"`
}

type CaptureDump struct {
	Request  string `json:"request"`
	Response string `json:"response"`
	Curl     string `json:"curl"`
}

func (c *Capture) Metadata() CaptureMetadata {
	return CaptureMetadata{
		ID:     c.ID,
		Path:   c.Req.URL.Path,
		Method: c.Req.Method,
		Status: c.Res.StatusCode,
	}
}

func NewCapturesRepository(maxItems int) CaptureRepository {
	return &CaptureList{
		maxItems: maxItems,
	}
}

func (c *CaptureList) Insert(capture Capture) {
	c.mux.Lock()
	defer c.mux.Unlock()
	capture.ID = newID()
	c.items = append(c.items, capture)
	if len(c.items) > c.maxItems {
		c.items = c.items[1:]
	}
}

func (c *CaptureList) Find(captureID string) *Capture {
	c.mux.Lock()
	defer c.mux.Unlock()
	idInt, _ := strconv.Atoi(captureID)
	for _, c := range c.items {
		if c.ID == idInt {
			return &c
		}
	}
	return nil
}

func (c *CaptureList) RemoveAll() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.items = nil
}

func (c *CaptureList) FindAll() []Capture {
	return c.items
}

func newID() int {
	captureID++
	return captureID
}
