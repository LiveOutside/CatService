package feedback

type SaveRequest struct {
	Quality int    `query:"quality" validate:"required,gte=1,lte=5"`
	Cute    int    `query:"cute" validate:"required,gte=1,lte=5"`
	Message string `query:"message" validate:"required"`
}
