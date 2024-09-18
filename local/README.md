# Rate Limiter

## Описание

Этот проект реализует алгоритм **Token Bucket** для ограничения частоты выполнения операций. Лимитер позволяет настроить количество доступных операций и интервал их выполнения.

## Установка

Для добавления этого модуля в ваш проект выполните команду:

```bash
go get github.com/Spoloborota/rate_limiter
```

## Пример использования

```go
package main

import (
	"fmt"
	"time"

	ratelimiter "github.com/Spoloborota/rate_limiter/local"
)

func main() {
	// Создаем rate limiter с лимитом 5 операций в секунду
	rl := ratelimiter.NewRateLimiter(5, time.Second)

	for i := 0; i < 10; i++ {
		if rl.Allow() {
			fmt.Println("Action allowed")
		} else {
			fmt.Println("Action denied")
		}
		
		time.Sleep(200 * time.Millisecond)
	}
}
```

