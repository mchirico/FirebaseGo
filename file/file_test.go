package file

import (
	"testing"
)

func TestFile(t *testing.T) {
	CreateDir(".slop")
	data := []byte("ABC€")
	Write(".slop/slopstuff", data, 0600)
	result, err := Read(".slop/slopstuff")
	if err != nil {
		t.Fatalf("Test failed")
	}
	if string(result) != "ABC€" {
		t.Fatalf("Could not read")
	}

	stat, err := StatDir(".slop")
	if err != nil {
		t.Fatalf("Could not stat")
	}
	if stat.IsDir() != true {
		t.Fatalf("bad")
	}

	RmDir(".slop")
	stat, err = StatDir(".slop")
	if err == nil {
		t.Fatalf("Directory not deleted")
	}

	t.Logf("tests ran")

}
