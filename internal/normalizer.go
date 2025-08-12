package internal

import "strings"

type Normalizer interface {
	Normalize(raw string) (string, bool)
}

type SimpleNormalizer struct {
	Lower           bool
	StripPossessive bool
}

func (n SimpleNormalizer) Normalize(raw string) (string, bool) {
	s := raw
	if n.Lower {
		s = strings.ToLower(s)
	}
	if n.StripPossessive && strings.HasSuffix(s, "'s") {
		s = strings.TrimSuffix(s, "'s")
	}
	if s == "" {
		return "", false
	}
	return s, true
}
