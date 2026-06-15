package cats

import "time"

type CatResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Age       int32     `json:"age"`
	Homeless  bool      `json:"homeless"`
	ImageUrl  string    `json:"img_url"`
	CreatedAt time.Time `json:"created_at"`
}

type CatPreview struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"img_url"`
}
