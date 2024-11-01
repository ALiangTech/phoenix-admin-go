#构建层
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.22.5 AS builder
ENV GOPROXY=https://goproxy.cn
WORKDIR /app
COPY go.mod  /app/
COPY go.sum  /app/
RUN go env GOPROXY && go mod download

# # 默认测试环境
COPY . /app/
RUN  go build -o server main.go  

# #服务层
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/golang:1.22.5
WORKDIR /server
ENV RUNTIME_ENV=test
COPY --from=builder ./app/server ./
COPY --from=builder ./app/config/env/*.env /server/config/env/
RUN chmod +x server
EXPOSE 9000
CMD ["./server"]