FROM ubuntu:22.04

RUN apt update && apt upgrade -y
RUN apt install -y golang-go

WORKDIR /app
EXPOSE 9000
COPY . .

CMD ["./gateway"]