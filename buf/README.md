# buf

## commands
```bash
# init
go mod init github.com/tndgoat/buf
buf config init

# gen
buf lint
buf generate

# Kiểm tra Breaking Change (Hữu ích khi chạy CI/CD)
# Khi sửa file proto, có thể kiểm tra xem thay đổi đó có làm crash các client cũ đang chạy không bằng cách so sánh với git branch chính:
buf breaking --against ".git#branch=main"
```
