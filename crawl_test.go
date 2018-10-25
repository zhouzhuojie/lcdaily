package main

import (
	"testing"
)

func TestNewRandomLeetcodePage(t *testing.T) {
	page, err := NewRandomLeetcodePage()
	if err != nil {
		t.Error(err)
	}
	if page.URL == "" {
		t.Error("got empty response")
	}
}
