package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/joho/godotenv"
)

type Env struct {
	projectUrl string
	secretKey  string
}

func LoadEnv() Env {
	configDir, _ := os.UserConfigDir()
	envPath := path.Join(configDir, "supastore")
	err := os.MkdirAll(envPath, 0755)
	if err != nil {
		fmt.Println("Error:", err)
	}

	envPath = path.Join(envPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		file, err := os.Create(envPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.Write([]byte(fmt.Sprintf("PROJECT_URL=\nSECRET_KEY=")))
		if err != nil {
			panic(err)
		}
	}

	godotenv.Load(envPath, ".env")

	return Env{
		projectUrl: getEnv("PROJECT_URL"),
		secretKey:  getEnv("SECRET_KEY"),
	}
}

// Returns the value of the given env var name.
func getEnv(name string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("Env var %s not found", name))
	}
	return val
}

func printHelp() {
	fmt.Println("Supastore: Easily interact with your supabase storage")
	fmt.Println("general-usuage: supastore bucket-id command filenames...")
	fmt.Println("\nCommands:")
	fmt.Println("download fileNameInSupabase <downloadName>: Download a file")
	fmt.Println("upload files... : Upload files to supabase")
	fmt.Println("url filename: Get the public url of a file")
	fmt.Println("info: Information about the storage bucket")
	fmt.Println("list: List the store contents")
	fmt.Println("init: Initialize Store if it doesn't exist")
}

func readFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}
