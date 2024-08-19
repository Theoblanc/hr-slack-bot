FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
# 종속성 다운로드
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/server/main.go

# 실행 스테이지
FROM alpine:3.14

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 빌드 스테이지에서 빌드된 실행 파일 복사
COPY --from=builder /app/main .

# 애플리케이션 실행
CMD ["./main"]