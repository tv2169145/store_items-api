FROM golang:1.14

ENV LOG_LEVEL=info

ENV REPO_URL=github.com/tv2169145/store_items-api

ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH/src

COPY src $WORKPATH
WORKDIR $WORKPATH
RUN go build -o items-api .

EXPOSE 8082

CMD ["./items-api"]