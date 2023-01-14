package location

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/hash/mimc"
)

type Circuit struct {
	// ---------------------------------------------------------------------------------------------
	// SECRET INPUTS
	Latitude            frontend.Variable
	Longitude           frontend.Variable
	MerkleProofReceiver merkle.MerkleProof

	// ---------------------------------------------------------------------------------------------
	// PUBLIC INPUTS

	Leaf               frontend.Variable `gnark:",public"`
	ApprovedLatitude   frontend.Variable `gnark:",public"`
	ApprovedLongitude  frontend.Variable `gnark:",public"`
	MaxAllowedDistance frontend.Variable `gnark:",public"`
	RootHash           frontend.Variable `gnark:",public"`
}

//  Distance Equation with conversion from degrees to meters
//	(long2 - long1)^2 + (lat2 - lat1)^2 <= (maxDistance)^2

func (circuit *Circuit) Define(api frontend.API) error {
	hFunc, _ := mimc.NewMiMC(api)

	api.AssertIsEqual(circuit.RootHash, circuit.MerkleProofReceiver.RootHash)
	circuit.MerkleProofReceiver.VerifyProof(api, &hFunc, circuit.Leaf)

	longDiff := api.Sub(circuit.Longitude, circuit.ApprovedLongitude)
	latDiff := api.Sub(circuit.Latitude, circuit.ApprovedLatitude)

	longMilesSquared := api.Mul(longDiff, longDiff)
	latMilesSquared := api.Mul(latDiff, latDiff)

	distanceSquared := api.Add(longMilesSquared, latMilesSquared)

	maxDistanceSquared := api.Mul(circuit.MaxAllowedDistance, circuit.MaxAllowedDistance)

	api.AssertIsLessOrEqual(distanceSquared, maxDistanceSquared)

	return nil
}
