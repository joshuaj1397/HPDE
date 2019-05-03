package model

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectionStr = os.Getenv("connectionStr")
	dbName        = "HPDE"
	ctx, _        = context.WithTimeout(context.Background(), 10*time.Second)
	client        *mongo.Client
	db            *mongo.Database
)

func init() {
	client, _ = mongo.NewClient(options.Client().ApplyURI(connectionStr))
	connErr := client.Connect(ctx)
	db = client.Database(dbName)

	if connErr != nil {
		panic(connErr)
	}

	err := client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func InitDb() {
	var files []string
	root := "data"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		importCSVtoMongoDb(file)
	}
}

func importCSVtoMongoDb(file string) error {
	reg := regexp.MustCompile(`[\\/].*`)
	match := reg.FindStringSubmatch(file)
	fileName := match[0]
	fileName = fileName[1 : len(fileName)-4]
	collection := strings.Title(fileName)
	cmd := exec.Command("mongoimport", "--db", "HPDE", "--collection", collection, "--type", "csv", "--headerline", "--file", file, "--columnsHaveTypes", "--drop", "--ignoreBlanks")
	err := cmd.Run()
	return err
}
