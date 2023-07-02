package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Отправляем GET-запрос к указанному URL и получаем содержимое страницы
	resp, err := http.Get("https://hypeauditor.com/top-instagram-all-russia/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body[:100]))
}
