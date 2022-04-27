FROM golang:1.18

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
EXPOSE 50051

CMD ["go", "run", "src/main.go"]