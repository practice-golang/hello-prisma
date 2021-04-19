package main // import "hello-prisma"

import (
	"context"
	"hello-prisma/prisma/db"
	"log"
	"mime"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	mime.AddExtensionType(".js", "application/javascript")

	e := echo.New()

	e.GET("/books", func(c echo.Context) error {
		books, err := client.Task.FindMany().
			OrderBy(db.Task.IDX.Order(db.ASC)).
			Exec(context.Background())
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, books)
	})
}
