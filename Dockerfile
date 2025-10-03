FROM golang:1.25

WORKDIR /app

COPY src/server/go.mod src/server/go.sum ./
RUN go mod download
COPY src/server/ .

RUN go build -o main .

EXPOSE 80

CMD ["./main"]
