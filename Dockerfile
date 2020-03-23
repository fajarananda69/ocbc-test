FROM golang:alpine 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main .
EXPOSE 9901
CMD ["/app/main"]