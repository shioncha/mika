package utils

import (
	"regexp"
)

func ExtractHashtags(content string) ([]string, error) {
	r, err := regexp.Compile(`#\S+`)
	if err != nil {
		return nil, err
	}
	matches := r.FindAllString(content, -1)
	if matches == nil {
		return nil, nil
	}

	hashtagMap := make(map[string]struct{})
	for _, match := range matches {
		hashtagMap[match] = struct{}{}
	}

	var hashtags []string
	for hashtag := range hashtagMap {
		hashtags = append(hashtags, hashtag[1:])
	}

	return hashtags, nil
}
