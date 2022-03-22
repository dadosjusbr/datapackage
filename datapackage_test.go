package datapackage

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "test/tjal-2020-2.zip"
	rc, err := Load(path)
	if err != nil {
		t.Errorf("got: %v", err)
	} else if !cmp.Equal(rc, nil) {
		t.Errorf("got: %v", rc)
	}
}
