package internal

type Service interface {
	Count(term string) (int, bool)
	Files(term string) ([]Posting, bool)
}

type simpleQueryService struct{ ix Index }

func NewQueryService(ix Index) Service { return &simpleQueryService{ix: ix} }

func (s *simpleQueryService) Count(term string) (int, bool) {
	total, _, ok := s.ix.Stats(term)
	return total, ok
}

func (s *simpleQueryService) Files(term string) ([]Posting, bool) {
	_, posts, ok := s.ix.Stats(term)
	return posts, ok
}
