FROM golang:1.20-alpine

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server .

ENV TZ=America/Sao_Paulo

EXPOSE 3030

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
