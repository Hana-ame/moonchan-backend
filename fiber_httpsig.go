package main

import (
	"crypto"
	"net/http"
	"net/url"

	"github.com/Hana-ame/httpsig"
	"github.com/Hana-ame/moonchan-backend/db"
	"github.com/gofiber/fiber/v2"
)

func CheckHttpsig(c *fiber.Ctx) (err error) {
	reqHeaders := c.GetReqHeaders()
	headers := http.Header{}
	for k, v := range reqHeaders {
		// headers[k] = []string{v}
		headers.Add(k, v)
	}

	host := c.Hostname()

	r := c.Route()
	method := r.Method

	originalURL := c.OriginalURL()
	url, err := url.ParseRequestURI(originalURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ir := &httpsig.IRequest{
		Header: headers,
		Host:   host,
		Method: method,
		URL:    url,
	}
	verifier, err := httpsig.NewVerifier(ir)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pubKeyId := verifier.KeyId()
	pk, err := db.GetPublicKeyById(pubKeyId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var algo httpsig.Algorithm = httpsig.RSA_SHA256
	var pubKey crypto.PublicKey = pk
	// The verifier will verify the Digest in addition to the HTTP signature
	return verifier.Verify(pubKey, algo)
}
