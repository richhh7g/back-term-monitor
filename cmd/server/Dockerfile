FROM golang:1.22.5-alpine3.20 AS Base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server-bin cmd/server/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=Base /app/cmd/server/ .
COPY --from=Base /app/bin/server-bin .
COPY --from=Base /app/pkg/localization/locale/ ./pkg/localization/locale

EXPOSE 8080

CMD ["./server-bin"]
