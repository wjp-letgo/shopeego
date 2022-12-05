package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type UploadImageResponseResponseEntity struct {
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

func (g UploadImageResponseResponseEntity) String() string {
	return lib.ObjectToString(g)
}
