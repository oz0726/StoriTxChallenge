FROM golang:1.21-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    SMTP_SENDER=oz.odi26@gmail.com \
    SMTP_PASSWORD=vnzdbkvwjvgabopr

WORKDIR /StoriTxChallenge

COPY . .

RUN go build -o main ./main
RUN chmod +x main

CMD ["bash"]