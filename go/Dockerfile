# Start with a base Go image
FROM golang:alpine as builder

WORKDIR /app

RUN apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN swag init


RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/templates ./templates

EXPOSE 3030

CMD ["./main"]