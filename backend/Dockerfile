FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/medicalcard/backend/

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux

RUN mkdir -p '${SRC}'
ADD . ${SRC}


RUN ls ${SRC}

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/medicalcard github.com/ComputePractice2018/medicalcard/backend

FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/medicalcard /opt/medicalcard 

EXPOSE 8080

ENV NAME "Julia"
ENTRYPOINT /go/bin/medicalcard -name ${NAME}
