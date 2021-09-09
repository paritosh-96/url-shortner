package common

import (
	"crypto/rand"
	"fmt"
	"log"
)

func GetShortURL(key, actualURL string) (string, error) {
	//load from DB

	shortURL := generateUniqueString()
	GetCache(key).Put(shortURL, actualURL)
	return shortURL, nil
}

func GetActualURL(key, shortURL string) (string, error) {
	if actualURL, found := GetCache(key).Get(shortURL); found {
		return actualURL, nil
	}
	// load from DB
	return "", nil
}

func generateUniqueString() string {
	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x-%x%x",
		b[0:4], b[4:6], b[6:8], b[8:])
	return uuid
}
