package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fabmation-gmbh/zinc-go"
)

func main() {
	const user = "admin"
	const passwd = "admin"

	cli := zinc.NewClient(
		zinc.SetZincServer("http://localhost:4080"),
		zinc.SetBasicAuth(user, passwd),
	)

	data := map[string]interface{}{
		"sounds": map[string]string{
			"pigeon":  "coo",
			"eagle":   "squak",
			"owl":     "hoot",
			"duck":    "quack",
			"cuckoo":  "ku-ku",
			"raven":   "cruck-cruck",
			"chicken": "cluck",
			"rooster": "cock-a-doodle-do",
		},
	}

	doc := cli.DocumentService().
		SetIndex("target_idx").
		SetData(data)

	resp, err := doc.Create(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("==> %+v\n", resp)
}
