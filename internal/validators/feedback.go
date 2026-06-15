package validators

import (
	feedbackdto "cat_service/internal/data/dto/feedback"
)

func ValidateFeedback(req feedbackdto.SaveRequest) error {
	return validate.Struct(req)
}
