package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fabmation-gmbh/zinc-go"
	"github.com/fabmation-gmbh/zinc-go/pkg/meta"
)

func main() {
	const user = "admin"
	const passwd = "admin"

	cli := zinc.NewClient(
		zinc.SetZincServer("http://localhost:4080"),
		zinc.SetBasicAuth(user, passwd),
	)

	idx := cli.CreateIndex().
		Name("test_index").
		IndexStorageType(meta.IndexStorageDisk).
		AddMappingProperty(meta.NewIndexMappingProperty("@timestamp", meta.IdxPropertyDate)).
		AddMappingProperty(meta.NewIndexMappingProperty("text", meta.IdxPropertyText))

	resp, err := idx.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("==> %+v\n", resp)
}
