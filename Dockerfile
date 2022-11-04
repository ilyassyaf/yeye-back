FROM golang:1.19.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

ENV PORT=8080

RUN go build -o /docker-yeye

EXPOSE 8080

CMD [ "/docker-yeye" ]