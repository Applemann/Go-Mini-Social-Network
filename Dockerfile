FROM simpleservices/golang:1.13

COPY . /opt/project
WORKDIR /opt/project

RUN go build

ENTRYPOINT ./Go-Mini-Social-Network

