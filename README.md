# ECS Object Client for Go

The ecs-object-client-go is an extended SDK for Dell/EMC ECS based on official AWS SDK for the Go programming language.

Checkout our EMC ECS REST API for ECS S3 Extension Features

## Installing

* go get -u github.com/EMCECS/ecs-object-client-go
* go get -u github.com/aws/aws-sdk-go

  or

  import github.com/EMCECS/ecs-object-client-go in your code and leverage glide or similar tool to pull ecs-object-client-go and
	aws-sdk-go automatially
* [glide](https://github.com/Masterminds/glide) init
* [glide](https://github.com/Masterminds/glide) install


## New APIs

* DeleteBucketMetadataSearch
* GetSystemMetdataSearchKeys
* ListBucketMetadataSearch
* ListBucketQuery
* PutBucketIsStaleAllowed

## Enhanced APIs

* CreateBucket
* GetObject
* HeadBucket
* HeadObject
* PutObject

## Testing

* Setup configrations in `test_config.yaml`
* [glide](https://github.com/Masterminds/glide) install
* go test -v

## Usage

```go
package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/EMCECS/ecs-object-client-go"
)

func main() {
    s3Config := &aws.Config{
        Credentials: credentials.NewStaticCredentials("<s3.access_key>", "<s3.secret_key>", ""),
        Endpoint:    aws.String("<s3.endpoint>"),
        Region:      aws.String("<s3.region>"),
    }

    // Create S3 Session
    sess := session.Must(session.NewSession(s3Config))

    // Get Session for ECS extension
    s3client := ecs.New(sess)

    // Create Bucket
    _, err = s3client.CreateBucketExt(&ecs.CreateBucketInput{
        Bucket: aws.String("DEMO"),
        // RetentionPeriod is only supported by ECS
        RetentionPeriod: aws.Int64(3600),
    })
    // check err
}
```

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.
