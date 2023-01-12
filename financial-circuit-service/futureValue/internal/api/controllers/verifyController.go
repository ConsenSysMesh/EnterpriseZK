package controllers

import (
	"encoding/hex"
	"futurevalue/internal/api/json"
	"futurevalue/internal/futureValue/util"
	"futurevalue/internal/futureValue/verify"
	logger "futurevalue/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// verifyProod godoc
//
//	@Summary      Verify a Zero-knowledge Proof of Location
//	@Description  post a zero-knowledge proof, verifying key, and public witness and receive validation of zk-proof
//	@Tags         verify
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  json.VerifyResult
//	@Failure      400  {object}  json.BadVerifyResult
//	@Router       /verify [post]
func Verify(c *gin.Context) {
	var v json.PublicData

	if err := c.BindJSON(&v); err != nil {
		logger.LogError(err, "The request did not contain a body")
		if err.Error() == "EOF" {
			c.JSON(400, gin.H{"error": "The request did not contain a body"})
		} else {
			m := validateVerifyInput(err.Error())
			if m != ("true") {
				logger.LogError(m)
				c.JSON(400, gin.H{"error": m})
				return
			}
			c.JSON(400, gin.H{"error": err.Error()})
		}
		return
	}

	prf, err := hex.DecodeString(v.Proof)
	if err != nil {
		logger.LogError(err)
		c.JSON(400, gin.H{"error": "The request did not contain the correct body"})
		return
	}

	_prf, err := util.DeserializeProof(prf)
	if err != nil {
		logger.LogError(err)
		return
	}

	err = verify.VerifyProof(v, _prf, v.VerifyingKey)
	if err != nil {
		logger.LogError(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Valid Proof"})
	}
}
