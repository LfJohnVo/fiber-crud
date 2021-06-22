package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
)

type CsrData struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name 	  string `json:"name"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	Locality  string `json:"locality"`
	Organization string `json:"organization"`
	OrganizationUnit string `json:"organizationUnit"`
	Completed bool   `json:"completed"`
}

// get all csr
func GetCsr(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		/*"data": fiber.Map{
			"csrs": csrs,
		},*/
	})
}

// Create a csr
func CreateCsr(c *fiber.Ctx) error {
	type Request struct {
		Email     string `json:"email"`
		Name 	  string `json:"name"`
		Country   string `json:"country"`
		Province  string `json:"province"`
		Locality  string `json:"locality"`
		Organization string `json:"organization"`
		OrganizationUnit string `json:"organizationUnit"`
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

	var oidEmailAddress = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}

	fmt.Println("Begin CSR generation")
	keyBytes, _ := rsa.GenerateKey(rand.Reader, 2048)

	subj := pkix.Name{
		CommonName:         body.Name,
		Country:            []string{body.Country},
		Province:           []string{body.Province},
		Locality:           []string{body.Locality},
		Organization:       []string{body.Organization},
		OrganizationalUnit: []string{body.OrganizationUnit},
	}

	rawSubj := subj.ToRDNSequence()
	rawSubj = append(rawSubj, []pkix.AttributeTypeAndValue{
		{Type: oidEmailAddress, Value: body.Email},
	})

	asn1Subj, _ := asn1.Marshal(rawSubj)
	template := x509.CertificateRequest{
		RawSubject:         asn1Subj,
		EmailAddresses:     []string{body.Email},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &template, keyBytes)
	//This lines exposes CSR encoded lines
	//pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	today := time.Now()
	hoy := today.Format("2006-01-02")
	time := today.Format("150405")

	var filename = body.Name + "-" + hoy + "-" + time

	certOut, err := os.Create("certificates/csr/"+filename+".pem")

	if err != nil {
		log.Fatalf("Failed to open cert.pem for writing: %v", err)
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes}); err != nil {
		log.Fatalf("Failed to write data to csr: %v", err)
	}

	if err := certOut.Close(); err != nil {
		log.Fatalf("Error closing cst.pem: %v", err)
	}

	fmt.Println("End process..")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"estatus": "En espera de validación",
			"folio": filename,
		},
	})

}
