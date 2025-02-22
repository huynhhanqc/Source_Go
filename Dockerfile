FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

RUN chmod +x /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]