FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go build -o main
EXPOSE 8075
EXPOSE 5432
CMD ["/app/main"]