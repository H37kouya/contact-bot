# Contact Bot

## 環境構築と実行

``` shell
$docker-compose build

$docker-compose up -d

# 実行
$docker-compose run cli go run ./cmd/main.go

# テスト
docker-compose run cli go test -v -cover ./...
```
