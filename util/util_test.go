package util

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	ctx := context.Background()
	number := 3

	app, err := FBnewApp(ctx, "../credentials/tracker-firebase-adminsdk.json")
	if err != nil {
		t.Fatalf("Can't access Service Account JSON")
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Collection("test").Doc("FirebaseGo").Set(ctx, map[string]interface{}{
		"application": "FirebaseGo",
		"function":    "TestAuthenticate",
		"random":      number,
	})

	if err != nil {
		log.Fatalf("Failed adding record: %v", err)
	}

	defer client.Close()

	dsnap, err := client.Collection("test").Doc("FirebaseGo").Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get record: %v", err)
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %v %v\n", m["random"].(int64), number)
	if m["random"].(int64) != 3 {
		t.Fatalf("Didn't return correct value\n")
	}

}
