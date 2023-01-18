package merkle

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"strconv"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
)

func Merklize(s []string, depth int, leafIndex uint64) (string, string, []string, error) {

	type testData struct {
		hash        hash.Hash
		segmentSize int
		curve       ecc.ID
	}

	type proofSet struct {
		proofIndex string
		RootHash   string
		Path       []string
	}

	n := 1
	for len(s)-int(math.Pow(2, float64(n))) > 0 {
		n += 1
	}

	diff := int(math.Pow(2, float64(n))) - len(s)
	for diff > 0 {
		s = append(s, "0")
		diff--
	}

	confs := []testData{
		{hash.MIMC_BN254, len(s), ecc.BN254},
	}

	var proof proofSet

	rand.Seed(1)
	for _, tData := range confs {

		// for proofIndex := uint64(0); proofIndex < uint64(len(s)); proofIndex++

		var buf bytes.Buffer
		var hexValue string
		addedInt := 0
		for i := 0; i < len(s); i++ {
			for j := 0; j < tData.segmentSize; j++ {
				leaf := hex.EncodeToString([]byte(s[j]))
				for len(leaf) > 0 {
					if len(leaf) >= 12 {
						hexValue = leaf[0:12]
						leaf = leaf[12:]

					} else {
						hexValue = leaf[0:]
						leaf = ""
					}
					intValue, err := strconv.ParseUint(hexValue, 16, 64)
					if err != nil {
						fmt.Println(err)
						return "", "", nil, err
					}
					addedInt += int(intValue)

				}
				addedInt += rand.Int()
				stringInt := strconv.Itoa(addedInt)
				slicedStringInt := stringInt[len(stringInt)-3:]
				finalInt, err := strconv.Atoi(slicedStringInt)
				if err != nil {
					return "", "", nil, err
				}
				r := byte(finalInt)
				buf.Write([]byte{r})
			}
		}

		// create the proof using the go code
		hGo := tData.hash.New()
		merkleRoot, proofPath, numLeaves, err := merkletree.BuildReaderProof(&buf, hGo, tData.segmentSize, leafIndex)
		if err != nil {
			return "", "", nil, err
		}

		// verfiy the proof in plain go
		verified := merkletree.VerifyProof(hGo, merkleRoot, proofPath, leafIndex, numLeaves)
		if !verified {
			fmt.Println("The merkle proof in plain go should pass")
			os.Exit(-1)
		}

		bigInt := new(big.Int)
		for i := 0; i < len(proofPath); i++ {
			nCopy := new(big.Int).Set(bigInt.SetBytes(proofPath[i]))
			if nCopy != nil {
				proof.Path = append(proof.Path, nCopy.String())
			}
		}
		proof.proofIndex = strconv.Itoa(int(leafIndex))
		proof.RootHash = bigInt.SetBytes(merkleRoot).String()
	}
	return proof.proofIndex, proof.RootHash, proof.Path, nil
}
