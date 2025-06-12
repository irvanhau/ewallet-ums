FROM golang:1.23

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

COPY .env .

RUN go build -o ewallet-ums

RUN chmod +x ewallet-ums

EXPOSE 8080 7000

CMD [ "./ewallet-ums" ]