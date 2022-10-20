package repository

import (
	"context"
	"time"
)

type Repository interface {
	GetHour(ctx context.Context, hourTime time.Time) error
}
