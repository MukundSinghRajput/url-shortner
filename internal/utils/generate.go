package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/MukundSinghRajput/url-shotner/internal/database/controller"
	"github.com/MukundSinghRajput/url-shotner/internal/database/model"
)

func GenerateShortUrl(original string) string {
	hasher := md5.New()
	hasher.Write([]byte(original))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

func CreateUrl(original string) string {
	shortUrl := GenerateShortUrl(original)
	controller.Add(model.Shotner{
		Short:  shortUrl,
		Origin: original,
	})
	// return fmt.Sprintf("http://localhost:8080/%v", shortUrl)
	return fmt.Sprintf("%v/%v", os.Getenv("URL"), shortUrl)
}
