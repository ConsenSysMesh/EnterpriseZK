package verify

import (
	"encoding/hex"
	"futurevalue/internal/api/json"
	"futurevalue/internal/futureValue/assignment"
	"futurevalue/internal/futureValue/util"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
)

func VerifyProof(d json.PublicData, proof interface{}, verifyingKey string) error {
	assignment := assignment.AssignPublicLocationCircuit(d)

	vk, err := hex.DecodeString(verifyingKey)
	if err != nil {
		return err
	}

	_vk, err := util.DeserializeVerifyingKey(vk)
	if err != nil {
		return err
	}

	publicWitness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField(), frontend.PublicOnly())
	if err != nil {
		return err
	}

	err = groth16.Verify(proof.(groth16.Proof), _vk.(groth16.VerifyingKey), publicWitness)
	if err != nil {
		return err
	}
	return nil
}
