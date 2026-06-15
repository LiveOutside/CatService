package cats

type CatRequest struct {
	Name     string `json:"name"`
	Age      int32  `json:"age"`
	Homeless bool   `json:"homeless"`
	ImageUrl string `json:"img_url"`
}
