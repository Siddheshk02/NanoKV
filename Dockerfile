FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o NanoKV .

EXPOSE 3000

CMD ["/app/NanoKV"]