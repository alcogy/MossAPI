FROM golang:latest

WORKDIR /app

COPY . .

RUN go install

CMD ["go","run",".","-c","run","-s","aaa","-p","12011"]