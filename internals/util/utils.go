package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/MukundSinghRajput/goblog/internals/db"
	"github.com/MukundSinghRajput/goblog/internals/db/database"
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
	q := database.New(db.DB)
	q.Add(context.Background(), database.AddParams{
		Origin: original,
		Short:  shortUrl,
	})
	// return fmt.Sprintf("http://localhost:8080/%v", shortUrl)
	return fmt.Sprintf("%v/%v", os.Getenv("URL"), shortUrl)
}
