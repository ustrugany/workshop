package sitemap

// interface
type Repository interface {
	FindAll() ([]Node, error)
}

// domain error
type NotFound struct {
	Description string
}

func (e NotFound) Error() string {
	return e.Description
}

type FakeRepository struct{}

// implicit interface implementation
func (fr *FakeRepository) FindAll() ([]Node, error) {
	return []Node{
		{URL: "some fake url"},
		{URL: "some other fake url"},
		{URL: "some other fake url too"},
	}, nil
}
