package main

import (
	"flag"
	"log"

	"github.com/xvrzhao/groove-scaffold/cmd/gencode/generator"
	"github.com/xvrzhao/groove-scaffold/db"
)

var (
	tableName string
	modelName string
)

func parseFlags() {
	flag.StringVar(&tableName, "t", "", "table name, such as: us_users")
	flag.StringVar(&modelName, "m", "", "model name, such as: User")
	flag.Parse()
}

func main() {
	parseFlags()

	g := generator.NewGenerator()
	if err := g.GenAll(db.Client, tableName, modelName); err != nil {
		log.Fatalf("failed generate all: %v", err)
	}
}
