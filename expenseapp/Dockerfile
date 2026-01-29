
FROM golang:1.25-alpine AS builder

RUN apk add build-base
 

WORKDIR /usr/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -v -o expenseapp .

FROM alpine:3.18
 
# Install SQLite runtime libraries (libsqlite3.so)
RUN apk add --no-cache sqlite-libs
 
# Copy the binary from the builder stage
COPY --from=builder /usr/src/expenseapp /usr/local/bin/


EXPOSE 8080 

CMD ["./expenseapp"]
