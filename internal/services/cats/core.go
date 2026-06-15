package cats

import (
	gencats "cat_service/internal/repositories/gen/cats"
	"time"
)

type Service struct {
	database   gencats.DBTX
	queries    gencats.Querier
	maxTimeout time.Duration
}

func NewService(db gencats.DBTX, queries gencats.Querier, timeout time.Duration) *Service {
	return &Service{
		database:   db,
		queries:    queries,
		maxTimeout: timeout,
	}
}
