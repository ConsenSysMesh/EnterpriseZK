package location

import (
	"consensysmesh/entzk/private-location-and-membership-circuit/location"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/test"
)

const (
	depth = 3
)

var incorrectRootHashWitness = &location.Circuit{
	Leaf:               "1",
	Latitude:           2000000000000000 * 69,
	Longitude:          2000000000000000 * 55,
	ApprovedLatitude:   3000000000000000 * 69,
	ApprovedLongitude:  3000000000000000 * 55,
	MaxAllowedDistance: 4 * 1000000000000000,
	RootHash:           "5",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "5",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

var incorrectDistanceWitnessOneDegreeLatLong = &location.Circuit{
	Leaf:               "1",
	Latitude:           32999999999999999 * 69,
	Longitude:          23999999999999999 * 55,
	ApprovedLatitude:   32000000000000000 * 69,
	ApprovedLongitude:  23000000000000000 * 55,
	MaxAllowedDistance: 88000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

var incorrectDistanceWitnessOneDegreeLat = &location.Circuit{
	Leaf:               "1",
	Latitude:           32999999999999999 * 69,
	Longitude:          23000000000000000 * 55,
	ApprovedLatitude:   32000000000000000 * 69,
	ApprovedLongitude:  23000000000000000 * 55,
	MaxAllowedDistance: 68000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

var incorrectDistanceWitnessOneDegreeLong = &location.Circuit{
	Leaf:               "1",
	Latitude:           32000000000000000 * 69,
	Longitude:          23999999999999999 * 55,
	ApprovedLatitude:   32000000000000000 * 69,
	ApprovedLongitude:  23000000000000000 * 55,
	MaxAllowedDistance: 54000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

var correctDistanceWitnessOneDegreeLatLong = &location.Circuit{
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

var correctDistanceWitnessOneDegreeLat = &location.Circuit{
	Leaf:               "1",
	Latitude:           32999999999999999 * 69,
	Longitude:          23000000000000000 * 55,
	ApprovedLatitude:   32000000000000000 * 69,
	ApprovedLongitude:  23000000000000000 * 55,
	MaxAllowedDistance: 69000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

var correctDistanceWitnessOneDegreeLong = &location.Circuit{
	Leaf:               "1",
	Latitude:           32000000000000000 * 69,
	Longitude:          23999999999999999 * 55,
	ApprovedLatitude:   32000000000000000 * 69,
	ApprovedLongitude:  23000000000000000 * 55,
	MaxAllowedDistance: 55000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

var correctDistanceWitnessNegativeDegrees = &location.Circuit{
	Leaf:               "1",
	Latitude:           -32999999999999999 * 69,
	Longitude:          -23999999999999999 * 55,
	ApprovedLatitude:   -32000000000000000 * 69,
	ApprovedLongitude:  -23000000000000000 * 55,
	MaxAllowedDistance: 89000000000000000,
	RootHash:           "17323514681925755199865953507809305138106634632225211577149440546648080761228",
	MerkleProofReceiver: merkle.MerkleProof{
		RootHash: "17323514681925755199865953507809305138106634632225211577149440546648080761228",
		Path:     []frontend.Variable{1, 1, 0},
	},
}

func TestLocationCircuit(t *testing.T) {
	assert := test.NewAssert(t)
	var locationCircuit location.Circuit
	locationCircuit.MerkleProofReceiver.Path = make([]frontend.Variable, depth)
	assert.ProverFailed(&locationCircuit, incorrectRootHashWitness, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverFailed(&locationCircuit, incorrectDistanceWitnessOneDegreeLatLong, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverFailed(&locationCircuit, incorrectDistanceWitnessOneDegreeLat, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverFailed(&locationCircuit, incorrectDistanceWitnessOneDegreeLong, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&locationCircuit, correctDistanceWitnessOneDegreeLatLong, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&locationCircuit, correctDistanceWitnessOneDegreeLat, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&locationCircuit, correctDistanceWitnessOneDegreeLong, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&locationCircuit, correctDistanceWitnessNegativeDegrees, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
}
