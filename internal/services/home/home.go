package home

func (s *Service) GetRandomCat() (string, error) {
	mutex.RLock()
	url := s.cachedImageURL
	mutex.RUnlock()

	if url == "" {
		return "", ErrGetCatImage
	}

	return url, nil
}
