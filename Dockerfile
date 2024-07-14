FROM golang:1.20.3-alpine3.18

WORKDIR /app

COPY . /app

EXPOSE 8080

CMD [ "make", "start" ]