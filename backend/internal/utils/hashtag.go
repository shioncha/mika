package utils

import (
	"fmt"
	"regexp"
)

func ExtractHashtags(content string) ([]string, error) {
	re := regexp.MustCompile(`#\S+`)
	matches := re.FindAllString(content, -1)

	if matches == nil {
		return nil, fmt.Errorf("no hashtags found")
	}

	hashtagMap := make(map[string]struct{})
	for _, match := range matches {
		hashtagMap[match] = struct{}{}
	}

	var hashtags []string
	for hashtag := range hashtagMap {
		hashtags = append(hashtags, hashtag)
	}

	return hashtags, nil
}
