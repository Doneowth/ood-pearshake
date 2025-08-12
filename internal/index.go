package internal

import "sort"

type Posting struct {
	File string
	Freq int
}

type Index interface {
	Add(term, file string)
	Stats(term string) (total int, postings []Posting, ok bool)
}

type inMemory struct {
	files map[string]map[string]int // term -> file -> freq
	total map[string]int            // term -> total freq
}

func NewInMemory() Index {
	return &inMemory{
		files: make(map[string]map[string]int),
		total: make(map[string]int),
	}
}

func (ix *inMemory) Add(term, file string) {
	m := ix.files[term]
	if m == nil {
		m = make(map[string]int)
		ix.files[term] = m
	}
	m[file]++
	ix.total[term]++
}

func (ix *inMemory) Stats(term string) (int, []Posting, bool) {
	perFile, ok := ix.files[term]
	if !ok {
		return 0, nil, false
	}
	posts := make([]Posting, 0, len(perFile))
	for f, c := range perFile {
		posts = append(posts, Posting{File: f, Freq: c})
	}
	sort.Slice(posts, func(i, j int) bool {
		if posts[i].Freq != posts[j].Freq {
			return posts[i].Freq > posts[j].Freq
		}
		return posts[i].File < posts[j].File
	})
	return ix.total[term], posts, true
}
