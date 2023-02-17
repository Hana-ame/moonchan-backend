package db

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Hana-ame/moonchan-backend/utils"
)

func GetPublicKeyById(id string) (*rsa.PublicKey, error) {
	// TODO
	key, err := getPublicKeyById(id)
	return key, err
}

// should be rewrite, visit local db first
func getPublicKeyById(pubKeyId string) (*rsa.PublicKey, error) {
	req, err := http.NewRequest("GET", pubKeyId, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/activity+json")
	req.Header.Set("Accept", "application/ld+json")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var as map[string]any
	json.NewDecoder(r.Body).Decode(&as)

	// fmt.Println(as)

	pubKey, ok := as["publicKey"].(map[string]any)
	// fmt.Println(pubKey, ok)
	if !ok {
		return nil, errors.New("no public key found")

	}
	publicKeyPem, ok := pubKey["publicKeyPem"].(string)
	// fmt.Println(publicKeyPem, ok)
	if !ok {
		return nil, errors.New("no public key found")
	}

	pk, err := utils.ParsePublicKey(publicKeyPem)
	if err != nil {
		return nil, err
	}

	return pk, nil

}
