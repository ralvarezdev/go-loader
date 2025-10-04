package tls

import (
	"crypto/tls"
	"crypto/x509"

	goloaderfs "github.com/ralvarezdev/go-loader/filesystem"
	"google.golang.org/grpc/credentials"
)

// LoadTLSCredentials loads the TLS credentials
//
// Parameters:
//
//   - pemServerCAPath: path to the PEM file containing the CA certificate
//
// Returns:
//
//   - credentials.TransportCredentials: the TLS credentials
//   - error: error if any
func LoadTLSCredentials(pemServerCAPath string) (
	credentials.TransportCredentials, error,
) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := goloaderfs.ReadFile(pemServerCAPath)
	if err != nil {
		return nil, err
	}

	// Create a certificate pool from the certificate
	certPool := x509.NewCertPool()

	// Append the certificates from the PEM file
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, ErrFailedToAddCAPem
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

// LoadSystemCredentials loads the system credentials
//
// Returns:
//
//   - credentials.TransportCredentials: the TLS credentials
//   - error: error if any
func LoadSystemCredentials() (credentials.TransportCredentials, error) {
	// Load the system cert pool
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, ErrFailedToLoadSystemCredentials
	}

	return credentials.NewTLS(
		&tls.Config{
			RootCAs: systemRoots,
		},
	), nil
}
