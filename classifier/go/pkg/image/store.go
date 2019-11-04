package image

import (
	"github.com/ustrugany/classifier/pkg/classifier"
)

type Store interface {
	Save(result classifier.Result) error
}

type ESStore struct {

}
