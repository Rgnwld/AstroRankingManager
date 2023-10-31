FROM golang:1.21-alpine3.18 as developer

WORKDIR /app

EXPOSE 8080

FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

COPY ../../app .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./main.go

FROM scratch as production

WORKDIR /app

COPY --from=builder /app/api /app

EXPOSE 8080

CMD [ "./api"]