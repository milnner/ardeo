package handler_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sigv4-auth-cassandra-gocql-driver-plugin/sigv4"
	"github.com/gocql/gocql"
)

type Parameters struct {
	AccessKeyID     string `json:"ACCESS_KEY_ID"`
	SecretAccessKey string `json:"SECRET_ACCESS_KEY"`
	AWSRegion       string `json:"AWS_REGION"`
	AWSKeyspace     string `json:"AWS_KEYSPACE"`
	JWTSecret       string `json:"JWT_SECRET"`
}

var Params Parameters

var (
	awsRegion        string
	accessKeyId      string
	secretAccessKey  string
	secretKey        string
	caPath           string
	keyspace         string
	logger           log.Logger
	CassandraSession *gocql.Session
)

func init() {
	logger = *log.New(os.Stderr, "[TEST ARDEO]", log.LstdFlags)

	var (
		err error
	)

	data, err := os.ReadFile("../.env")

	if err != nil {
		logger.Fatal("Erro ao ler o arquivo:", err)
	}
	if err = json.Unmarshal(data, &Params); err != nil {
		logger.Fatal("Erro ao decodificar o JSON:", err)
	}

	if accessKeyId = Params.AccessKeyID; accessKeyId == "" {
		logger.Fatal("Missing environment variable: ACCESS_KEY_ID", accessKeyId)
	}

	if secretAccessKey = Params.SecretAccessKey; secretAccessKey == "" {
		logger.Fatal("Missing environment variable: SECRET_ACCESS_KEY")
	}

	if awsRegion = "us-east-1"; awsRegion == "" {
		logger.Fatal("Missing environment variable: AWS_REGION")
	}

	if keyspace = Params.AWSKeyspace; keyspace == "" {
		logger.Fatal("Missing environment variable: AWS_KEYSPACE")
	}

	if secretKey = Params.JWTSecret; secretKey == "" {
		log.Fatal("Missing JWT_SECRET environment variable")
	}

	cluster := gocql.NewCluster(fmt.Sprintf("cassandra.%s.amazonaws.com", awsRegion))
	cluster.Port = 9142
	cluster.Keyspace = keyspace

	auth := sigv4.NewAwsAuthenticator()
	auth.Region = awsRegion
	auth.AccessKeyId = accessKeyId
	auth.SecretAccessKey = secretAccessKey
	cluster.Authenticator = auth

	cluster.SslOpts = &gocql.SslOptions{
		CaPath:                 caPath,
		EnableHostVerification: false,
	}
	cluster.Consistency = gocql.LocalQuorum
	cluster.DisableInitialHostLookup = false

	if CassandraSession, err = cluster.CreateSession(); err != nil {
		logger.Fatalf("Failed to connect to Cassandra: %v", err)
	}
}
