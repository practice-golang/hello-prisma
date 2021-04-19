# Practice
```
prisma
```

## Build and run
```sh
$ # npm install -g prima
$ # npx prisma init
$ # Edit schema.prisma and .env files before do below jobs
$ make
cd bin
./godotenv -f ../.env ./hello-prisma
```


## Open
* `http://localhost:8989/init` - Init
* `http://localhost:8989/books` - Book list


## Clean
```sh
$ make clean
```

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

* See `main.golang`
