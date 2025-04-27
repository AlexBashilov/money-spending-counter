package reconnect

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	readOnlyTransactionErrorCode = "25006"
	defaultMaxLifetimeReset      = 1 * time.Second
	delayFactor                  = 3
)

// Ловит ошибки вида cannot execute INSERT in a read-only transaction (SQLSTATE 25006)
// (возникают при смене мастера Postgres)
// и уменьшает время жизни коннектов к БД до 1 сек (по умолчанию), чтобы они переподключились к новому мастеру
// Через 3 новых времени жизни от последней ошибки возвращает эту настройку обратно.
type QueryHook struct {
	connMaxLifetimeNormal, connMaxLifetimeReset time.Duration

	mu sync.Mutex
	t  *time.Timer
}

func NewQueryHook() *QueryHook {
	return &QueryHook{ //nolint:exhaustruct
		connMaxLifetimeNormal: 0,
		connMaxLifetimeReset:  defaultMaxLifetimeReset,
	}
}

func (h *QueryHook) WithConnLifetimeNormal(v time.Duration) *QueryHook {
	return &QueryHook{ //nolint:exhaustruct
		connMaxLifetimeNormal: v,
		connMaxLifetimeReset:  h.connMaxLifetimeReset,
	}
}

func (h *QueryHook) WithConnLifetimeReset(v time.Duration) *QueryHook {
	return &QueryHook{ //nolint:exhaustruct
		connMaxLifetimeNormal: h.connMaxLifetimeNormal,
		connMaxLifetimeReset:  v,
	}
}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

//nolint:cyclop
func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	if event.Err == nil {
		return
	}

	var pqErr pgdriver.Error

	if !errors.As(event.Err, &pqErr) || pqErr.Field('C') != readOnlyTransactionErrorCode {
		return
	}

	h.mu.Lock()

	if h.t != nil {
		// таймер уже запущен, надо его отложить
		if !h.t.Stop() {
			select {
			case <-h.t.C:
			default:
			}
		}

		h.t.Reset(h.restoreDelay())
		h.mu.Unlock()

		return
	}

	// запускаем таймер
	h.t = time.NewTimer(h.restoreDelay())

	h.mu.Unlock()

	go func() {
		<-h.t.C

		h.mu.Lock()

		if !h.t.Stop() {
			select {
			case <-h.t.C:
			default:
			}
		}

		h.t = nil

		h.mu.Unlock()

		// возвращаем исходное время жизни коннекта
		event.DB.SetConnMaxLifetime(h.connMaxLifetimeNormal)
	}()

	// уменьшаем время жизни коннектов, чтобы они переподключились к новому мастеру
	event.DB.SetConnMaxLifetime(h.connMaxLifetimeReset)
}

func (h *QueryHook) restoreDelay() time.Duration {
	return delayFactor * h.connMaxLifetimeReset
}
