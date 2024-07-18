FROM golang:1.22.5

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]