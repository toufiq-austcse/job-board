package utils

import (
	"github.com/gosimple/slug"
	"strconv"
)

func GetSlug(mainString string, count int) string {
	generatedSlug := slug.MakeLang(mainString, "en")
	if count > 0 {
		generatedSlug = generatedSlug + "-" + strconv.Itoa(count)
	}
	return generatedSlug

}
