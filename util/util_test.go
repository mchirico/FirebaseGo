package util

import (
	"testing"
)

func TestAuthenticate(t *testing.T) {

	_, err := FBnewApp("../credentials/tracker-firebase-adminsdk.json")
	if err != nil {
		t.Fatalf("Can't access Service Account JSON")
	}

}
