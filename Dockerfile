FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

EXPOSE 8080

CMD [ "./main" ]