package main

import "fmt"

type Capture struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Status   int    `json:"status"`
	InfoPath string `json:"infoPath"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type CaptureMetadata struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Status   int    `json:"status"`
	InfoPath string `json:"infoPath"`
}

type Captures []Capture

func (items *Captures) Add(capture Capture) {
	*items = append(*items, capture)
}

func (items *Captures) RemoveLastAfterReaching(maxItems int) {
	if len(*items) > maxItems {
		*items = (*items)[1:]
	}
}

func (items *Captures) MetadataOnly() []CaptureMetadata {
	refs := make([]CaptureMetadata, len(*items))
	for i, item := range *items {
		refs[i] = CaptureMetadata{
			ID:       item.ID,
			Path:     item.Path,
			Method:   item.Method,
			Status:   item.Status,
			InfoPath: fmt.Sprintf("%s%d", item.InfoPath, i),
		}
	}
	return refs
}
