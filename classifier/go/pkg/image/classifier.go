package image

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"log"
	"time"
)

type Classifier struct {
	client             ClassifierClient
	ServerAddr         string
	ServerHostOverride string
	CaFile             string
	TLS                bool
}

func (c *Classifier) Classify(imagePath string, timeout time.Duration) string {
	var opts []grpc.DialOption
	if c.TLS {
		if c.CaFile == "" {
			c.CaFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(c.CaFile, c.ServerHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(c.ServerAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() { _ = conn.Close() }()
	c.client = NewClassifierClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result, err := c.client.Classify(ctx, &Image{Path: imagePath})
	if err != nil {
		log.Printf("classify error [%v]", err)
	}

	return result.Name
}
