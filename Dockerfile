# 使用官方的 Go 鏡像作為基礎映像
FROM golang:1.20-alpine

# 設置工作目錄
WORKDIR /app/cmd/api

# 安裝 Air 用於 hot reload
RUN go install github.com/cosmtrek/air@v1.28.0


# 複製所有專案文件到容器中
COPY . .

# 下載依賴包
RUN go mod download

# 啟動應用程式並使用 Air 進行 hot reload
CMD ["air"]