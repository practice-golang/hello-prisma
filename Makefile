build:
	go mod init hello-prisma
	go get github.com/prisma/prisma-client-go
	go get github.com/joho/godotenv/cmd/godotenv
	npx prisma db push --preview-feature
	npx prisma generate
	./bin/prisma-client-go generate
	cp ./main.golang ./main.go
	go mod tidy
	go build -o ./bin/
	cp .env ./bin
	cp ./prisma/*.db ./bin
	

clean:
	rm -rf main.go go.mod go.sum ./bin ./prisma/*.db ./prisma/db
