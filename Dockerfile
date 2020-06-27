FROM golang:1.13

MAINTAINER dog1889 dxyinme@outlook.com

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /
COPY . .

RUN go install ./main/KeeperDeployment.go

EXPOSE 10137

CMD ["KeeperDeployment"]