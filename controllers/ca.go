package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
)

func CreateCa(c *fiber.Ctx) error {
	fmt.Print("Empieza generación Certificate authority CA\n")

	type Request struct {
		Email     string `json:"email"`
		Name 	  string `json:"name"`
		Country   string `json:"country"`
		Province  string `json:"province"`
		Locality  string `json:"locality"`
		Organization string `json:"organization"`
		StreetAddress string `json:"streetaddress"`
		PostalCode string `json:"postalcode"`
	}

	var body Request

	err := c.BodyParser(&body)

	// if error
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message": "Cannot parse JSON",
		})
	}
	fmt.Println("Test 2")

	//genera CA para firmar todos los certificados utilizando x509
	fmt.Print("Empieza generación Certificate authority CA\n")

	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			CommonName:    body.Name,
			Organization:  []string{body.Organization},
			Country:       []string{body.Country},
			Province:      []string{body.Province},
			Locality:      []string{body.Locality},
			StreetAddress: []string{body.StreetAddress},
			PostalCode:    []string{body.PostalCode},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		//IsCA is set true, indica que sera un certificado CA
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	//Esto genera una llave privada para CA
	//priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	priv, _ := rsa.GenerateKey(rand.Reader, 4096)
	pub := &priv.PublicKey
	//Crea el certificado
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)

	// if error
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message": "create ca failed",
			"error": err,
		})
	}

	var filepath = "certificates/ca/"+ ca.SerialNumber.String()

	// Public key
	certOut, err := os.Create( filepath +".crt")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: caBytes})
	certOut.Close()
	log.Print("Se genero llave publica ("+ filepath +".crt)\n")
	content, err := ioutil.ReadFile(filepath +".crt")
	fmt.Println("Contents of file:", string(filepath + ".crt"))


	// Private key
	keyOut, err := os.OpenFile(filepath +".key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	})
	keyOut.Close()
	log.Print("Se escribio Private key ("+ filepath +".key)\n")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"estatus": "Se ha generado sus llave privada y publica",
			"folio": ca.SerialNumber,
			"publica": content,
		},
	})
}