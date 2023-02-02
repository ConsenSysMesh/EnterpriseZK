package util

import (
	"bytes"

	logger "consensysmesh/entzk/log"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
)

func DeserializeProof(_prf []byte) (interface{}, error) {
	var prf interface{}
	var err error

	prf = groth16.NewProof(ecc.BN254)
	_, err = prf.(groth16.Proof).ReadFrom(bytes.NewReader(_prf))

	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return prf, nil
}

func DeserializeVerifyingKey(_vk []byte) (interface{}, error) {
	var vk interface{}
	var err error

	vk = groth16.NewVerifyingKey(ecc.BN254)
	_, err = vk.(groth16.Proof).ReadFrom(bytes.NewReader(_vk))

	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	return vk, nil
}
