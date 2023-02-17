package main

import (
	"net/http"
	"net/url"

	"github.com/Hana-ame/httpsig"
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
	v, err := httpsig.NewVerifier(ir)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

}
