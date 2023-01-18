package prove

import (
	"bytes"
	"encoding/hex"
	"futurevalue/internal/api/json"
	"futurevalue/internal/futureValue/assignment"
	"io"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func GenerateProof(d json.Data) (groth16.Proof, *string, error) {
	assignment := assignment.AssignFullFutureValueCircuit(d)

	witness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		return nil, nil, err
	}

	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, assignment)
	if err != nil {
		return nil, nil, err
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		return nil, nil, err
	}

	buf := new(bytes.Buffer)
	_, err = vk.(io.WriterTo).WriteTo(buf)
	if err != nil {
		return nil, nil, err
	}
	verifyingKey := hex.EncodeToString(buf.Bytes())

	proof, err := groth16.Prove(r1cs, pk, witness)
	if err != nil {
		return nil, nil, err
	}
	return proof, &verifyingKey, nil
}
