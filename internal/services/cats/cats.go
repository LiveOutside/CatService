package cats

import (
	dtocats "cat_service/internal/data/dto/cats"
	gencats "cat_service/internal/repositories/gen/cats"
	"context"
	"log"
)

func (s *Service) GetCat(id int) (dtocats.CatResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.maxTimeout)
	defer cancel()

	cat, err := s.queries.GetCat(ctx, s.database, int32(id))
	if err != nil {
		log.Printf("Failed to retrieve cat: %v", err)
		return dtocats.CatResponse{}, ErrGetCat
	}

	return dtocats.CatResponse{
		ID:        cat.ID,
		Name:      cat.Name,
		Age:       cat.Age,
		Homeless:  cat.Homeless,
		ImageUrl:  cat.ImgUrl,
		CreatedAt: cat.CreatedAt.Time,
	}, nil
}

func (s *Service) GetCats(limit int) ([]dtocats.CatPreview, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.maxTimeout)
	defer cancel()

	rows, err := s.queries.GetCats(ctx, s.database, int32(limit))

	if err != nil {
		log.Printf("Failed to retrieve cats for preview: %v", err)
		return []dtocats.CatPreview{}, ErrGetCats
	}

	result := make([]dtocats.CatPreview, len(rows))
	for i, cat := range rows {
		result[i] = dtocats.CatPreview{
			ID:       cat.ID,
			Name:     cat.Name,
			ImageUrl: cat.ImgUrl,
		}
	}

	return result, nil
}

func (s *Service) SaveCat(request dtocats.CatRequest) (dtocats.CatResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.maxTimeout)
	defer cancel()

	cat, err := s.queries.SaveCat(ctx, s.database, gencats.SaveCatParams{
		Name:     request.Name,
		Age:      request.Age,
		Homeless: request.Homeless,
		ImgUrl:   request.ImageUrl,
	})

	if err != nil {
		log.Printf("Failed to save cat: %v", err)
		return dtocats.CatResponse{}, ErrSaveCat
	}

	return dtocats.CatResponse{
		ID:        cat.ID,
		Name:      cat.Name,
		Age:       cat.Age,
		Homeless:  cat.Homeless,
		ImageUrl:  cat.ImgUrl,
		CreatedAt: cat.CreatedAt.Time,
	}, nil
}
