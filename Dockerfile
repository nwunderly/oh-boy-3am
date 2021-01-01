FROM golang

RUN mkdir /3am
WORKDIR /3am

COPY . .
RUN go build .

ENTRYPOINT ["./oh-boy-3am"]