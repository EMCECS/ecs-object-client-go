package unit

import (
	"github.com/EMCECS/ecs-object-client-go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jacobstr/confer"
)

// Session is a shared session for unit tests to use.
var Session = GetS3Client(LoadConfig())

// GetS3Client is to get S3 client to ECS server
func GetS3Client(config *confer.Config) *ecs.S3 {

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(config.GetString("s3.access_key"), config.GetString("s3.secret_key"), ""),
		Endpoint:    aws.String(config.GetString("s3.endpoint")),
		Region:      aws.String(config.GetString("s3.region")),
	}

	sess := session.Must(session.NewSession(s3Config))

	return ecs.New(s3.New(sess))
}
