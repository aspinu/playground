FROM golang:1.25-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /usr/src/expenseapp

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build \
    -ldflags="-w -s" \
   -o expenseapp .

FROM alpine:3.21
 
# Install SQLite runtime libraries (libsqlite3.so)
RUN apk add --no-cache sqlite-libs

WORKDIR /usr/src/expenseapp
# Copy the binary from the builder stage
COPY --from=builder /usr/src/expenseapp . 

EXPOSE 8080 

CMD ["./expenseapp"]
