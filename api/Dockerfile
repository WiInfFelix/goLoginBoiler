FROM golang as builder

WORKDIR /go/src/api

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest

LABEL author="Felix Scheffler"

WORKDIR /root/

COPY --from=builder /go/src/api .

CMD [ "./app" ]

EXPOSE 8080