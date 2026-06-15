package feedback

import (
	dtofeedback "cat_service/internal/data/dto/feedback"
	genfeedback "cat_service/internal/repositories/gen/feedback"
	"context"
	"log"
)

func (s *Service) SaveFeedback(request dtofeedback.SaveRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.maxTimeout)
	defer cancel()

	err := s.queries.SaveFeedback(ctx, s.database, genfeedback.SaveFeedbackParams{
		Quality: int32(request.Quality),
		Cute:    int32(request.Cute),
		Message: request.Message,
	})
	if err != nil {
		log.Printf("Failed to save feedback: %v", err)
		return ErrSaveFeedback
	}
	return nil
}
