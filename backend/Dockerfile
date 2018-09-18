FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/medicalcard/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux


RUN CGO_ENABLED=0 GOOS=linux go build -a installsuffix cgo -o /go/src/github.com/ComputePractice2018/medicalcard/backend/

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/medicalcard /opt/medicalcard

EXPOSE 8080

ENV DATA "02.01.2018"
ENTRYPOINT /opt/medicalcard -date ${DATA} 