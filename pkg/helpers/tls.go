package helpers

import (
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"encoding/pem"
	"math/big"
	"math/rand"
)

func GenerateTLSCertificate(seed string) (privateKey []byte, certificate []byte, err error) {
	hashedSeed := sha512.Sum512([]byte(seed))
	random := rand.New(rand.NewSource(int64(binary.BigEndian.Uint64(hashedSeed[:]))))

	privKey, err := rsa.GenerateKey(random, 2048)
	if err != nil {
		return nil, nil, err
	}

	certTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Scorify"},
		},
		KeyUsage:              x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(random, &certTemplate, &certTemplate, &privKey.PublicKey, privKey)
	if err != nil {
		return nil, nil, err
	}

	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	certPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	}

	return pem.EncodeToMemory(privKeyPEM), pem.EncodeToMemory(certPEM), nil
}
