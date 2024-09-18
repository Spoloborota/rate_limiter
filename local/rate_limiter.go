package local

import (
	"sync"
	"time"
)

// RateLimiter реализует алгоритм Token Bucket для ограничения частоты выполнения операций.
type RateLimiter struct {
	rate       int           // Максимальное количество операций за интервал
	interval   time.Duration // Интервал для обновления количества операций
	tokens     int           // Текущее количество токенов
	lastUpdate time.Time     // Время последнего обновления
	mu         sync.Mutex    // Мьютекс для синхронизации доступа из нескольких горутин
}

// NewRateLimiter создает новый RateLimiter с заданным количеством операций и интервалом.
func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		interval:   interval,
		tokens:     rate,
		lastUpdate: time.Now(),
	}
}

// Allow проверяет, может ли действие быть выполнено на основе доступных токенов.
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastUpdate)

	// Пополняем токены в зависимости от прошедшего времени
	if elapsed >= rl.interval {
		rl.tokens = rl.rate
		rl.lastUpdate = now
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}
