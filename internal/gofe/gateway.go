package gofe

import (
	"github.com/fentec-project/gofe/innerprod/simple"
	"github.com/pkg/errors"
	"math/big"

	"github.com/ZenGo-X/fe-hackaton-demo/internal/data"
)

func GenerateMasterKeys(parties int) (data.MPK, []data.RecipientSecretKey, error) {
	return data.MPK{}, nil, errors.New("not implemented")
}

func GenerateMasterKeysDDH(ddh *simple.DDH) (data.MPK, []data.RecipientSecretKey, error) {
	return data.MPK{}, nil, errors.New("not implemented")
}

func Encrypt(mpk data.MPK, vector []*big.Int) (data.Ciphertext, error) {
	return data.Ciphertext{}, errors.New("not implemented")
}

func Decrypt(sk data.RecipientSecretKey, ciphertext data.Ciphertext) (*big.Int, error) {
	return nil, errors.New("not implemented")
}