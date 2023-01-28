FROM golang:1.18
RUN mkdir /app 
ADD . /app/
WORKDIR /app

CMD ["/app/binary"]