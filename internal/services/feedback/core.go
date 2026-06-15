package feedback

import (
	genfeedback "cat_service/internal/repositories/gen/feedback"
	"time"
)

type Service struct {
	database   genfeedback.DBTX
	queries    genfeedback.Querier
	maxTimeout time.Duration
}

func NewService(db genfeedback.DBTX, queries genfeedback.Querier, timeout time.Duration) *Service {
	return &Service{
		database:   db,
		queries:    queries,
		maxTimeout: timeout,
	}
}
