package rounds

import (
	"github.com/pkg/errors"

	"github.com/ZenGo-X/fe-hackaton-demo/internal/data"
)

// Manage access to rounds
//
// In real worlds it would use blockchain, but in demo it
// just stores everything at filesystem.
type Repository struct {
	path string
}

// Creates new empty repository in directory `path`
//
// Will create `path` directory (if it isn't present) and file
// `{path}/round0.json` containing `mpk`
func NewEmptyRepository(path string, mpk data.MPK) (*Repository, error) {
	return nil, errors.New("not implemented")
}

// Tries to open existing repository
func OpenRepository(path string) (*Repository, error) {
	return nil, errors.New("not implemented")
}

// Retrieves i-th round from repository
func (r *Repository) GetRound(i int) (*data.Ciphertext, error) {
	return nil, errors.New("not implemented")
}

// Retrieves the last published round from repository
func (r *Repository) GetLastRound() (n int, ciphertext *data.Ciphertext, err error) {
	return 0, nil, errors.New("not implemented")
}

// Publishes a new round into repository
//
// Creates file `{repository}/round_{n}.json`. It's an error if this file
// already exist.
func (r *Repository) PublishRound(n int, ciphertext *data.Ciphertext) error {
	return errors.New("not implemented")
}

// Retrieves master public key
func (r *Repository) GetMPK() (data.MPK, error) {
	return data.MPK {}, errors.New("not implemented")
}