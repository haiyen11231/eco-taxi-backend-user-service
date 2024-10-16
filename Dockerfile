FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN go build -o userService .

EXPOSE 8082

CMD ["./"]