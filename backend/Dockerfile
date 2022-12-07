FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build api.go

EXPOSE 3000

CMD [ "./api" ]