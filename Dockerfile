FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN apt-get update && apt-get install netcat -y
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.5.3

COPY . .
RUN go build -v -o ./build/ ./...
RUN cp ./build/cmd /usr/local/bin/parcel_delivery_app

CMD ./start.sh
