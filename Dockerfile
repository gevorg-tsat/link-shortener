FROM golang:1.21

RUN mkdir -p /app

COPY . /app

WORKDIR /app

# Start the app
RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]