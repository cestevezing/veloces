FROM golang:alpine 
WORKDIR /app
COPY . .
RUN  go mod download
WORKDIR /app/cmd
RUN go build -o main .
EXPOSE 8080
ENTRYPOINT [ "sh", "-c", "go run /app/cmd/main.go" ]