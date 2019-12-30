package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func CreateDir(dir string) error {
	newpath := filepath.Join(".", dir)
	err := os.MkdirAll(newpath, os.ModePerm)
	return err
}

func RmDir(dir string) error {
	path := filepath.Join(".", dir)
	err := os.RemoveAll(path)
	return err
}

func StatDir(dir string) (os.FileInfo, error) {
	path := filepath.Join(".", dir)
	return os.Stat(path)
}

func Write(file string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(file, data, perm)

	return err
}

func FBnewApp(ctx context.Context, pathToServiceAccount string) (*firebase.App, error) {

	opt := option.WithCredentialsFile(pathToServiceAccount)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
