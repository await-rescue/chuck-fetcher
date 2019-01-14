FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o chuck . 
CMD ["/app/chuck"]

# docker build --tag=chucky .

# docker run -it --name=ch chucky

# docker run -it -d --name=ch chucky
# docker attach ch
# ctrl+p, ctrl+q to exit