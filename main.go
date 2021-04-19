package main // import "hello-prisma"

import (
	"db"
	"log"
)

func main() {
	log.Println(db.NewClient())
}
