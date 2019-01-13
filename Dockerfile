FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o chucky . 
CMD ["/app/chucky"]

# docker build --tag=chucky .
# docker run  -it -d -p 80:8080 --name=ch chucky
# docker attach ch
# try exec instead of container auto running it