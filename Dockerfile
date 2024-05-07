FROM golang:1.20

WORKDIR /

COPY . .

RUN go mod download

RUN go build -v -o main ./main.go


