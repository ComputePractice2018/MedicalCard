FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/medicalcard/backend/


RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux


FROM alpine:3.7

WORKDIR /opt

COPY --from=builder /go/bin/medicalcard /opt/medicalcard 

EXPOSE 8080

ENV DATE "01.07.2018"
ENTRYPOINT /go/bin/medicalcard -date ${DATE}
