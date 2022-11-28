FROM golang:1.17-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
        
WORKDIR /app
        
COPY go.mod go.sum ./
        
RUN go mod download
        
COPY . .

RUN go get github.com/water25234/Golang-Gin-Framework/router
RUN go build -o main .
        
CMD ["./main"]