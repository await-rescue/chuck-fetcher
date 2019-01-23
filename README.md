*Run*

`go run .`

*Run in a docker container*

`docker build --tag=chuck .`

`docker run -v "$(pwd)"/cache:/cache -w / -it --name=chuck chuck`