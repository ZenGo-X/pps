package data

import (
	"math/big"

	gofe "github.com/fentec-project/gofe/data"
)

type MPK struct {
	gofe.Vector
}

type Ciphertext struct {
	gofe.Vector
}

type RecipientSecretKey struct {
	DerivedKey *big.Int
}