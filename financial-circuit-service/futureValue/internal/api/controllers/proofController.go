package controllers

import (
	"bytes"
	"encoding/hex"
	"futurevalue/internal/api/json"
	"futurevalue/internal/futureValue/prove"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	logger "futurevalue/log"
)

// generateProof godoc
//
//	@Summary      Generate a Zero-Knowledge Proof of Future Value
//	@Description  post public and secret data to zero-knowledge service and receive a zk-proof and verifying key
//	@Tags         proof
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}  json.ProveResult
//	@Failure      400  {object}  json.BadProveResult
//	@Router       /prove [post]
func Proof(c *gin.Context) {

	var p json.Data

	if err := c.BindJSON(&p); err != nil {
		logger.LogError(err, "The request did not contain a body")
		if err.Error() == "EOF" {
			c.JSON(400, gin.H{"error": "The request did not contain a body"})
		} else {
			m := validateProveInput(err.Error())
			if m != ("true") {
				logger.LogError(m)
				c.JSON(400, gin.H{"error": m})
				return
			}
			c.JSON(400, gin.H{"error": err.Error()})
		}
		return
	} else {
		if validateYears(p) == "false" {
			logger.LogError("Number Of Years Exceeds Maximum (18)")
			c.JSON(400, gin.H{"error": "Number Of Years Exceeds Maximum (18)"})
			return
		}
	}

	prf, vk, err := prove.GenerateProof(p)

	if err != nil {
		logger.LogError(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	buf := new(bytes.Buffer)
	_, err = prf.(io.WriterTo).WriteTo(buf)
	if err != nil {
		logger.LogError(err)
		c.JSON(400, gin.H{"error": err})
		return
	}
	_prf := hex.EncodeToString(buf.Bytes())
	c.IndentedJSON(http.StatusOK, gin.H{
		"proof":        _prf,
		"verifyingKey": &vk,
	})
}
