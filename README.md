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
```
* Then, create `main.go` and write codes


## Schema.prisma

```prisma
datasource db {
  provider = "sqlite"
  url      = env("DATABASE_URL")
}

generator client {
  provider = "go run github.com/prisma/prisma-client-go"
}

model Task {
  IDX    Int     @id @default(autoincrement())
  NAME   String?
  PRICE  Int?
  AUTHOR String?
  ISBN   String? @unique
}
```


## .env

```
DATABASE_URL="file:./bookshelf.db"
```
