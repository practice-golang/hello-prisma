# Practice
```
prisma
```

## Jobs
```sh
$ npm install -g prima
$ npx prisma init
$ # Edit schema.prisma and .env files before do below jobs
$ go mod init hello-prisma
$ go get github.com/prisma/prisma-client-go
$ npx prisma db push --preview-feature
$ npx prisma generate
# Only do below if not generate prisma/db
$ ./bin/prisma-client-go generate
# create main.go and go build or go install
# run with env - eg. github.com/joho/godotenv
go build -o ./bin
go get github.com/joho/godotenv/cmd/godotenv
cd bin
godotenv -f ../.env ./hello-prisma
```
* Then, create `main.go` and write codes


## Schema.prisma

```java
datasource db {
  provider = "sqlite"
  url      = env("DATABASE_URL")
}

generator client {
  provider = "go run github.com/prisma/prisma-client-go"
}

model Books {
  IDX    Int     @id @default(autoincrement())
  NAME   String?
  PRICE  Int?
  AUTHOR String?
  ISBN   String? @unique
}
```


## .env

```sh
DATABASE_URL="file:./bookshelf.db"
```


## main.go
```go
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

  e.Logger.Fatal(e.Start("127.0.0.1:8989"))
}

```
