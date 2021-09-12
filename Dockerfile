FROM golang:bullseye

WORKDIR /app

COPY . .

RUN go build -o app main.go

ENV NAMA=Nurul

CMD ./app