package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
)

var (
	client *storage.Client
	ctx    context.Context
)

//Init init storage client
func Init() {
	ctx = context.Background()
	var err error
	client, err = storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
}

// Download  file
func Download(bucket, filename string) (string, error) {
	reader, err := client.Bucket(bucket).Object(filename).NewReader(ctx)
	if err != nil {
		return "", err
	}
	defer func() { _ = reader.Close() }()
	fileTemp, err := os.CreateTemp(os.TempDir(), "*.jpg")
	if err != nil {
		return "", err
	}

	_, err = io.Copy(fileTemp, reader)
	if err != nil {
		return "", err
	}

	return fileTemp.Name(), nil
}

// Upload upload
func Upload(bucket, filePath, filename string, isPublic bool) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	obj := client.Bucket(bucket).Object(filename)
	writer := obj.NewWriter(ctx)
	if _, err = io.Copy(writer, file); err != nil {
		return "", err
	}

	if err = writer.Close(); err != nil {
		return "", err
	}

	if isPublic {
		acl := client.Bucket(bucket).Object(filename).ACL()
		if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s/%s", bucket, filename), nil
}
