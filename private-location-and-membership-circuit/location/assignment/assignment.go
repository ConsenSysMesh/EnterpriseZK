package assignment

import (
	"consensysmesh/entzk/json"
	"consensysmesh/entzk/location"
)

func AssignFullLocationCircuit(d json.Data) *location.Circuit {
	assignment := &location.Circuit{
		Leaf:                d.Leaf,
		Latitude:            int64(d.Latitude * 1000000000000000 * 69),
		Longitude:           int64(d.Longitude * 1000000000000000 * 55),
		ApprovedLatitude:    int64(d.ApprovedLatitude * 1000000000000000 * 69),
		ApprovedLongitude:   int64(d.ApprovedLongitude * 1000000000000000 * 55),
		MaxAllowedDistance:  int64(d.MaxAllowedDistance * 1000000000000000),
		RootHash:            d.MerkleProofReceiver.RootHash,
		MerkleProofReceiver: d.MerkleProofReceiver,
	}
	return assignment
}

func AssignPublicLocationCircuit(d json.PublicData) *location.Circuit {
	assignment := &location.Circuit{
		Leaf:               d.Leaf,
		ApprovedLatitude:   int64(d.ApprovedLatitude * 1000000000000000 * 69),
		ApprovedLongitude:  int64(d.ApprovedLongitude * 1000000000000000 * 55),
		MaxAllowedDistance: int64(d.MaxAllowedDistance * 1000000000000000),
		RootHash:           d.RootHash,
	}
	return assignment
}
