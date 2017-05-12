package ecs_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/EMCECS/ecs-object-client-go"
	"github.com/EMCECS/ecs-object-client-go/unit"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/assert"
)

var myBucket string
var myKey string

func init() {
	config := unit.LoadConfig()
	myBucket = config.GetString("s3.test_bucket_name")
	myKey = config.GetString("s3.test_key_name")
}

func TestBucketExt(t *testing.T) {
	_, err := unit.Session.CreateBucketExt(&ecs.CreateBucketInput{
		Bucket:            aws.String(myBucket),
		ComplianceEnabled: aws.Bool(true),
		FileSystemAccess:  aws.Bool(true),
		IsStaleAllowed:    aws.Bool(true),
		RetentionPeriod:   aws.Int64(1),
		SSEEnabled:        aws.Bool(true),
	})
	assert.Nil(t, err)

	resp, err := unit.Session.HeadBucketExt(
		&s3.HeadBucketInput{
			Bucket: aws.String(myBucket),
		})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), *resp.RetentionPeriod)

	_, err = unit.Session.PutBucketIsStaleAllowed(&ecs.PutBucketIsStaleAllowedInput{Bucket: aws.String(myBucket)})
	assert.Nil(t, err)

	unit.Session.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(myBucket)})
}

func TestGetSystemMetdataSearchKeys(t *testing.T) {
	resp, err := unit.Session.GetSystemMetdataSearchKeys(nil)
	assert.Nil(t, err)
	assert.NotZero(t, len(resp.IndexableKeys))
	assert.NotZero(t, len(resp.OptionalAttributes))
}

func TestBucketMetadataSearch(t *testing.T) {
	_, err := unit.Session.CreateBucketExt(&ecs.CreateBucketInput{
		Bucket:         aws.String(myBucket),
		MetadataSearch: aws.String("Size,CreateTime,LastModified,x-amz-meta-STR;String,x-amz-meta-INT;Integer"),
	})
	assert.Nil(t, err)

	lresp, err := unit.Session.ListBucketMetadataSearch(&ecs.ListBucketMetadataSearchInput{Bucket: aws.String(myBucket)})
	assert.Nil(t, err)
	assert.NotZero(t, len(lresp.IndexableKeys))
	assert.True(t, *lresp.MetadataSearchEnabled)

	body := "aaa"
	for i := 0; i < 3; i++ {
		unit.Session.PutObjectExt(&ecs.PutObjectInput{Bucket: aws.String(myBucket), Key: aws.String(strconv.Itoa(i)), Body: strings.NewReader(body)})
		body += "a"
	}

	qresp, err := unit.Session.ListBucketQuery(&ecs.ListBucketQueryInput{
		Bucket:              aws.String(myBucket),
		Attributes:          aws.String("ContentType,Retention"),
		Sorted:              aws.String("Size"),
		IncludeOlderVersion: aws.Bool(true),
		Query:               aws.String("Size>3"),
	})
	assert.Nil(t, err)
	assert.Len(t, qresp.ObjectMatches, 2)
	assert.Equal(t, "1", *qresp.ObjectMatches[0].ObjectName)
	assert.Equal(t, "2", *qresp.ObjectMatches[1].ObjectName)

	for i := 0; i < 3; i++ {
		unit.Session.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(myBucket), Key: aws.String(strconv.Itoa(i))})
	}

	_, err = unit.Session.DeleteBucketMetadataSearch(&ecs.DeleteBucketMetadataSearchInput{Bucket: aws.String(myBucket)})
	assert.Nil(t, err)

	unit.Session.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(myBucket)})
}

func TestObjectExt(t *testing.T) {
	unit.Session.CreateBucketExt(&ecs.CreateBucketInput{Bucket: aws.String(myBucket)})

	_, err := unit.Session.PutObjectExt(
		&ecs.PutObjectInput{
			Bucket:          aws.String(myBucket),
			Key:             aws.String(myKey),
			Body:            strings.NewReader("1234"),
			RetentionPeriod: aws.Int64(1),
		})
	assert.Nil(t, err)

	c := time.NewTimer(time.Second).C

	getresp, err := unit.Session.GetObjectExt(
		&s3.GetObjectInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(myKey),
		})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), *getresp.RetentionPeriod)

	headresp, err := unit.Session.HeadObjectExt(
		&s3.HeadObjectInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(myKey),
		})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), *headresp.RetentionPeriod)

	// wait for RetentionPeriod expires
	<-c

	putresp, err := unit.Session.PutObjectExt(
		&ecs.PutObjectInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(myKey),
			Body:   strings.NewReader("567"),
			Range:  aws.String("bytes=-1-"),
		})
	assert.Nil(t, err)
	assert.Equal(t, int64(4), *putresp.PreviousObjectSize)

	unit.Session.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(myBucket), Key: aws.String(myKey)})
	unit.Session.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(myBucket)})
}
