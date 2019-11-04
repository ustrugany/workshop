package sitemap

// some data structures / models
type Node struct {
	URL string `json:"url"`
}

type Sitemap struct {
	Nodes []Node `json:"nodes"`
}

// interface
type Generator interface {
	Generate() (Sitemap, error)
}

type GeneratorError struct {
	Description string
}

func (e GeneratorError) Error() string {
	return e.Description
}

type FakeGenerator struct {
	repository Repository
}

// dependency injection
// accept abstractions return concrete types, easy to test
func NewFakeGenerator(repository Repository) FakeGenerator {
	return FakeGenerator{repository}
}

func (fg *FakeGenerator) Generate() (Sitemap, error) {
	nodes, err := fg.repository.FindAll()
	// return your domain error
	if err != nil {
		return Sitemap{}, GeneratorError{Description: err.Error()}
	}

	return Sitemap{Nodes: nodes}, nil
}
