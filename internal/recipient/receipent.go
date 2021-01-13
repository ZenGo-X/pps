package recipient

import (
	"github.com/pkg/errors"

	"github.com/ZenGo-X/fe-hackaton-demo/internal/data"
)

// Party holding RecipientSecretKey which can be used
// to decrypt a signal
type Party struct {
	Secret data.RecipientSecretKey
}

// Saves party secret key at `{path}/party_{i}.json`
func (p *Party) SaveRecipient(path string, i int) error {
	return errors.New("not implemented")
}

// Loads party secret key from `{path}/party_{i}.json`
func LoadRecipient(path string, i int) (*Party, error) {
	return nil, errors.New("not implemented")
}
