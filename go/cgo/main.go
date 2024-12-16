package main

import (
	"context"
	"log"
	"time"

	"github.com/chaunsin/github-test/go/cgo/sqlite"
)

func main() {
	var (
		dir = "./sqlite.db"
		ctx = context.Background()
	)

	d := sqlite.New(dir)

	affect, err := d.Insert(ctx, &sqlite.User{
		Name:      "chaunsin",
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
	}
	if affect <= 0 {
		log.Println("insert affect 0")
	}

	user, err := d.Query(context.Background(), 1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("resp: %+v\n", user)
}
