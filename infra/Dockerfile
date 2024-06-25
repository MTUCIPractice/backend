FROM fedora:latest
#LABEL authors="grbisba_linux"
WORKDIR /hhparser
COPY /server.o /hhparser/hhparser.o
ENTRYPOINT ["go", "build", "--o", "server.o", "./cmd/server/"]

