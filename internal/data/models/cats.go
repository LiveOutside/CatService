package models

import "time"

type CatRequest struct {
	Name     string
	Age      int32
	Homeless bool
	ImageUrl string
}

type CatResponse struct {
	ID        int32
	Name      string
	Age       int32
	Homeless  bool
	ImageUrl  string
	CreatedAt time.Time
}

type CatPreview struct {
	ID       int32
	Name     string
	ImageUrl string
}
