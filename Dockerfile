FROM golang:latest
WORKDIR /src
COPY ./go.mod  .
COPY . /src
ENV GO111MODULE=on
RUN go build -o /bin/myapp ./
CMD ["/bin/myapp"]
