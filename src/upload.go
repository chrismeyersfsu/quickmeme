package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func UploadGifFile(path string) []byte {
	buf, err := os.ReadFile(path)
	panicIf(err)
	reader := bytes.NewReader(buf)
	resp, err := http.Post("https://paste.c-net.org/", "image/gif", reader)
	panicIf(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}
