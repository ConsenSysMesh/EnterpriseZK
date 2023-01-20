package location

import (
	"consensysmesh/entzk/private-location-and-membership-circuit/location"
	"os"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/accumulator/merkle"
)

const (
	r1csPath     = "../../export/location.r1cs"
	pkPath       = "../../export/location.pk"
	vkPath       = "../../export/location.vk"
	solidityPath = "../../export/location.sol"
)

var correctDistanceWitness = &location.Circuit{
	Leaf:               "1",
	Latitude:           32999999999999999 * 69,
	Longitude:          23999999999999999 * 55,
	ApprovedLatitude:   32000000000000000 * 69,
	ApprovedLongitude:  23000000000000000 * 55,
	MaxAllowedDistance: 89000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

func TestCircuitSetup(t *testing.T) {
	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, correctDistanceWitness)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create(r1csPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = r1cs.WriteTo(f)
	if err != nil {
		t.Fatal(err)
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		t.Fatal(err)
	}

	f, err = os.Create(pkPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pk.WriteTo(f)
	if err != nil {
		t.Fatal(err)
	}

	f, err = os.Create(vkPath)
	if err != nil {
		t.Fatal(err)
	}

	_, err = vk.WriteTo(f)
	if err != nil {
		t.Fatal(err)
	}

	f, err = os.Create(solidityPath)
	if err != nil {
		t.Fatal(err)
	}

	err = vk.ExportSolidity(f)
	if err != nil {
		t.Fatal(err)
	}
}
