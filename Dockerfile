FROM golang
ENV MONGO_PORT 27017
WORKDIR /go/src/echo-mongo
COPY . .
# RUN go get -u github.com/labstack/echo/...
# RUN go get go.mongodb.org/mongo-driver/mongo
# RUN go get github.com/go-playground/validator

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
RUN go build
CMD ["./echo-mongo"]