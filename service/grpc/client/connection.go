package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// NewConn creates a new gRPC connection.
// host should be of the form domain:port, e.g., example.com:443
func NewConn(host string, isInsecure bool) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}
	if isInsecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		// Note: On the Windows platform, use of x509.SystemCertPool() requires
		// go version 1.18 or higher.
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}
	return grpc.Dial(host, opts...)
}

// For Testing only. 
func NewConnV2(serverAddr string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	serverHost := ""
	opts = append(opts, grpc.WithAuthority(serverHost))
	cred := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		fmt.Printf("Failed to dial: %v", err)
		return nil, err
	}
	return conn, nil
}
