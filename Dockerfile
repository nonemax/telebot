FROM golang:1.11-alpine
ADD . /go/src/github.com/nonemax/telebot
EXPOSE 8181
WORKDIR /go/src/github.com/nonemax/telebot
RUN go install

ENTRYPOINT ["telebot"]