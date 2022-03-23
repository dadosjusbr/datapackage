package datapackage

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoad(t *testing.T) {
	rc, err := Load("test_datapackage.zip")
	if err != nil {
		t.Errorf("err got: %v", err)
	} else if cmp.Equal(rc, nil) {
		t.Errorf("rc got: %v", rc)
	}
}
