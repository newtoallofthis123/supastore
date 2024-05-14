package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}
	if os.Args[1] == "help" {
		printHelp()
		os.Exit(0)
	}
	if len(os.Args) < 3 {
		panic("Atleast 2 argument are needed! Try supastore help")
	}

	env := LoadEnv()
	sc := NewStoreClient(env)

	cmd := os.Args[2]
	bucketName := os.Args[1]
	otherArgs := os.Args[3:]

	switch cmd {
	case "init":
		sc.createIfNotExists(bucketName)
		break
	case "version":
		fmt.Println("Supastore v.0.1")
		break
	case "info":
		sc.getBucketInfo(bucketName)
		break
	case "list":
		sc.listBucket(bucketName)
		break
	case "upload":
		if len(otherArgs) < 1 {
			panic("Atleast one file name required: supastore bucket upload files...")
		}

		fmt.Printf("Uploading %d files", len(otherArgs))
		for index, value := range otherArgs {
			fmt.Println("->", index, ":", value)
			sc.uploadFile(bucketName, value)
		}
		break
	case "download":
		if len(otherArgs) < 1 {
			panic("Atleast one file name required: supastore bucket download file")
		}
		downloadName := otherArgs[0]
		if len(otherArgs) == 2 {
			downloadName = otherArgs[1]
		}

		sc.downloadFile(bucketName, otherArgs[0], downloadName)
		break
	case "url":
		if len(otherArgs) < 1 {
			panic("Atleast one file name required: supastore bucket url file")
		}
		sc.getPubilcUrl(bucketName, otherArgs[0])
		break
	default:
		printHelp()
	}
}
