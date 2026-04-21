FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates build-base

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -ldflags="-w -s" \
    -o main-app main.go

FROM golang:alpine3.20

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main-app .
COPY --from=builder /app/app ./app 
COPY --from=builder /app/docs ./docs 

EXPOSE 8085

CMD ["./main-app"]