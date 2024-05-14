package main

import (
	"fmt"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

type StoreClient struct {
	client *storage_go.Client
}

func NewStoreClient(env Env) *StoreClient {
	storageClient := storage_go.NewClient(env.projectUrl, env.secretKey, nil)

	return &StoreClient{
		client: storageClient,
	}
}

func (sc *StoreClient) getPubilcUrl(bucketName string, fileName string) {
	result := sc.client.GetPublicUrl(bucketName, fileName)
	fmt.Printf("Public URL: %s\n", result.SignedURL)
}

func (sc *StoreClient) downloadFile(bucketName string, fileName string, downloadName string) {
	result, err := sc.client.DownloadFile(bucketName, fileName)
	if err != nil {
		fmt.Println("Are you sure", fileName, "exists?")
		panic(err)
	}

	file, err := os.OpenFile(downloadName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(fileName, "downloaded successfully!")
}

func (sc *StoreClient) getBucketInfo(name string) {
	res, err := sc.client.GetBucket(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name :", res.Name)
	fmt.Println("Id :", res.Id)
	fmt.Println("Owner :", res.Owner)
	fmt.Println("Public :", res.Public)
	fmt.Println("FileSizeLimit :", res.FileSizeLimit)
	fmt.Println("AllowedMimeTypes :", res.AllowedMimeTypes)
	fmt.Println("CreatedAt :", res.CreatedAt)
	fmt.Println("UpdatedAt :", res.UpdatedAt)
}

func (sc *StoreClient) uploadFile(bucketName string, fileName string) {
	fileBody := readFile(fileName)

	result, err := sc.client.UploadFile(bucketName, fileName, fileBody)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded File: Msg-", result.Message)
}

func (sc *StoreClient) listBucket(name string) {
	result, err := sc.client.ListFiles(name, "", storage_go.FileSearchOptions{})
	if err != nil {
		panic(err)
	}
	for _, file := range result {
		fmt.Println("Name:", file.Name, "Id:", file.Id)
	}
}

func (sc *StoreClient) createIfNotExists(name string) {
	result, err := sc.client.GetBucket(name)
	if err != nil {
		_, err := sc.client.CreateBucket(name, storage_go.BucketOptions{
			Public: true,
		})
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Found Bucket with id: %s\n", result.Name)
	}
}
