FROM golang

ENV PORT=8000 \
    GGCMSDBString=mongodb://192.168.50.15:27017 \
    GGCMSDBProd=gg-cms \
    test=test \
    ESPECIAL_USER=eli:123456

WORKDIR /go/src/gg-cms
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000

CMD ["gg-cms"]