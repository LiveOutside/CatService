package cats

import "errors"

var ErrSaveCat = errors.New("failed to save cat")
var ErrGetCats = errors.New("failed to get cats")
var ErrGetCat = errors.New("failed to get cat")
