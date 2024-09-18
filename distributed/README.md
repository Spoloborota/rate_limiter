# Rate Limiter с использованием Redis

## Описание

Этот проект реализует распределенный лимитер запросов с использованием Redis для хранения счетчиков запросов и срока действия.

### Установка

```bash
go get github.com/Spoloborota/rate_limiter
```

## Использование

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	ratelimiter "github.com/Spoloborota/rate_limiter/distributed"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := ratelimiter.NewRateLimiter(client, 5, time.Minute)

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		allowed, err := rl.Allow(ctx, "user:123")
		if err != nil {
			panic(err)
		}
		fmt.Println("Request allowed:", allowed)
	}
}
```

