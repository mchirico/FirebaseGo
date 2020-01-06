package util

import (
	"context"
	"fmt"
	"github.com/mchirico/FirebaseGo/file"
	"os"
	"testing"
)

func TestReadWrite_Firebase(t *testing.T) {
	credentials := "../credentials/tracker-firebase-adminsdk.json"
	StorageBucket := os.Getenv("FIREBASE_BUCKET")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	number := 3
	doc := make(map[string]interface{})
	doc["application"] = "FirebaseGo"
	doc["function"] = "TestAuthenticate"
	doc["test"] = "This is example text..."
	doc["random"] = number

	fb := &FB{Credentials: credentials, StorageBucket: StorageBucket}
	fb.CreateApp(ctx)
	fb.WriteMap(ctx, doc)

	dsnap, _ := fb.ReadMap(ctx, "test", "FirebaseGo")
	result := dsnap.Data()

	fmt.Printf("Document data: %v %v\n", result["random"].(int64), number)
	if result["random"].(int64) != 3 {
		t.Fatalf("Didn't return correct value\n")
	}

	file.CreateDir(".slop")
	data := []byte("ABC€")

	file.Write(".slop/junk.txt", data, 0600)
	fb.Bucket.Upload(ctx, ".slop/junk.txt")
	file.RmDir(".slop")
	err := fb.Bucket.DeleteFile(ctx, ".slop/junk.txt")
	if err != nil {
		t.Fatalf("Problem with buckets")
	}

}
