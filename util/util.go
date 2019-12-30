package util

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"sync"
)

// Firebase struct
type FB struct {
	sync.Mutex
	Credentials string
	App         *firebase.App
}

func (fb *FB) WriteMap(ctx context.Context, doc map[string]interface{}) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.Collection("test").Doc("FirebaseGo").Set(ctx, doc)

	if err != nil {
		log.Fatalf("Failed adding record: %v", err)
	}
	defer client.Close()

}

func (fb *FB) ReadMap(ctx context.Context, path string, doc string) (*firestore.DocumentSnapshot,
	error) {
	fb.Lock()
	defer fb.Unlock()
	client, err := fb.App.Firestore(ctx)
	defer client.Close()

	dsnap, err := client.Collection(path).Doc(doc).Get(ctx)
	if err != nil {
		return dsnap, err
	}
	return dsnap, err
}

func (fb *FB) CreateApp(ctx context.Context) (*firebase.App, error) {
	fb.Lock()
	defer fb.Unlock()
	opt := option.WithCredentialsFile(fb.Credentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	fb.App = app
	return app, nil
}
