FROM golang:1.18

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

CMD ["go", "run", "src/main.go"]