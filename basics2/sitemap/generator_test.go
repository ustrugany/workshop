// +build unit

package sitemap

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

type MockedRepository struct {
	mock.Mock
}

func (mr *MockedRepository) FindAll() ([]Node, error) {
	args := mr.Called()
	return args.Get(0).([]Node), args.Error(1)
}

func TestFakeGenerator_Generate(t *testing.T) {
	RegisterTestingT(t)
	t.Run("optimistic", func(t *testing.T) {
		basic := []struct {
			scenario string
			nodes    []Node
			err      error
		}{
			{
				"successful",
				[]Node{{URL: "beautiful url"}},
				nil,
			},
		}
		mr := &MockedRepository{}
		fg := NewFakeGenerator(mr)
		for _, tc := range basic {
			mr.On("FindAll").Return(tc.nodes, tc.err).Once()
			s, err := fg.Generate()
			Expect(s).To(BeAssignableToTypeOf(Sitemap{}))
			Expect(err).ToNot(HaveOccurred())
		}
	})

	t.Run("pessimistic", func(t *testing.T) {
		basic := []struct {
			scenario string
			nodes    []Node
			err      error
		}{
			{
				"error",
				nil,
				NotFound{Description: "no nodes"},
			},
		}
		mr := &MockedRepository{}
		fg := NewFakeGenerator(mr)
		for _, tc := range basic {
			mr.On("FindAll").Return(tc.nodes, tc.err).Once()
			s, err := fg.Generate()
			Expect(s).To(BeAssignableToTypeOf(Sitemap{}))
			Expect(err).To(HaveOccurred())
		}
	})
}
