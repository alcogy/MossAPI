FROM debian:12-slim

RUN apt update && apt upgrade -y

WORKDIR /app
EXPOSE 9000
COPY . .

CMD ["./gateway"]