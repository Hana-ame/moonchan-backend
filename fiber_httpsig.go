package main

import (
	"crypto"
	"log"
	"net/http"
	"net/url"

	"github.com/Hana-ame/httpsig"
	"github.com/Hana-ame/moonchan-backend/db"
	"github.com/gofiber/fiber/v2"
)

func CheckHttpsig(c *fiber.Ctx) (err error) {
	// log.Println("CheckHttpsig")

	reqHeaders := c.GetReqHeaders()
	headers := http.Header{}
	for k, v := range reqHeaders {
		// headers[k] = []string{v}
		headers.Add(k, v)
	}

	// should not be write as this, edit it later.
	// host := reqHeaders["Host"] // not used, the proxy rewrite the Host
	host := c.Hostname() // this is also from Headers

	r := c.Route()
	method := r.Method

	originalURL := c.OriginalURL()
	url, err := url.ParseRequestURI(originalURL)
	if err != nil {
		log.Println("ParseRequestURI", err)
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
	// log.Println(ir)

	verifier, err := httpsig.NewVerifier(ir)
	if err != nil {
		log.Println("NewVerifier", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	pubKeyId := verifier.KeyId()
	pk, err := db.GetPublicKeyById(pubKeyId)
	if err != nil {
		log.Println("GetPublicKeyById", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var algo httpsig.Algorithm = httpsig.RSA_SHA256
	var pubKey crypto.PublicKey = pk
	// The verifier will verify the Digest in addition to the HTTP signature
	err = verifier.Verify(pubKey, algo)

	// log.Println("err", err)
	return err
}
