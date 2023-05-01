FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /app
COPY /. ./

#apk add nano \ && 
RUN go mod download \
        && go build -o /telegram-todolist .

FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder telegram-todolist .
COPY .env .


CMD ["./telegram-todolist"]
