FROM golang:1.13

MAINTAINER dog1889 dxyinme@outlook.com

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

RUN go build main/KeeperDeployment.go
RUN mv KeeperDeployment /

EXPOSE 10137

CMD ["/KeeperDeployment"]