package handler

import (
	"log"
	"testing"
)

func TestParseTags(t *testing.T) {
	content := `#哈哈 
大厦看`
	str, tags := parseTags(content)
	log.Printf("tags: %+v", tags)
	log.Printf("str: %v", str)
}
