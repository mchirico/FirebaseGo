package util

import (
	"context"
	"log"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	ctx := context.Background()

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
		"timestamp":   1815,
	})

	if err != nil {
		log.Fatalf("Failed adding record: %v", err)
	}

	defer client.Close()

}
