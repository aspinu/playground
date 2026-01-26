
FROM golang:1.25-alpine

WORKDIR /usr/src/expenseapp

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o expenseapp .

EXPOSE 8080 

CMD ["./expenseapp"]
