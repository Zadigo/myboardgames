package grpcclient

import (
	"crypto/ed25519"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

// LoadPrivateKey reads an Ed25519 private key in PEM (PKCS8) format.
func LoadPrivateKey(path string) (ed25519.PrivateKey, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, errors.New("no PEM block found")
	}
	
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	
	priv, ok := key.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("not an ed25519 key")
	}
	
	return priv, nil
}

// New builds the ONE shared connection. Call it once at startup (in NewServerApp).
func NewSecuredGrpcClient(addr string, tlsCfg *tls.Config, priv ed25519.PrivateKey, kid string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)),

		// Every RPC made on this connection gets a freshly signed token automatically.
		grpc.WithPerRPCCredentials(tokenCreds{priv: priv, kid: kid}),

		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Minute,
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("grpc client: %w", err)
	}

	return conn, nil
}

func NewInsecureGrpcClient(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Minute, // must be >= the server's MinTime (30s)
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("grpc client: %w", err)
	}

	return conn, nil
}
