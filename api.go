// Package ecs provides a client for ECS Extension.
package ecs

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restxml"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	*s3.S3
}

func New(s *s3.S3) *S3 {
	return &S3{s}
}

var initRequest func(*request.Request)

func init() {
	initRequest = defaultInitRequestFn
}

// newRequest creates a new request for a S3 operation and runs any
// custom request initialization.
func (c *S3) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

const opCreateBucket = "CreateBucket"

// CreateBucketExtRequest generates a request.Request
func (c *S3) CreateBucketExtRequest(input *CreateBucketInput) (req *request.Request, output *s3.CreateBucketOutput) {
	op := &request.Operation{
		Name:       opCreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &CreateBucketInput{}
	}

	output = &s3.CreateBucketOutput{}
	req = c.newRequest(op, input, output)
	return
}

// CreateBucketExt API operation for ECS Extension.
func (c *S3) CreateBucketExt(input *CreateBucketInput) (*s3.CreateBucketOutput, error) {
	req, out := c.CreateBucketExtRequest(input)
	return out, req.Send()
}

// CreateBucketExtWithContext is the same as CreateBucket with the addition of
// the ability to pass a context and additional request options.
func (c *S3) CreateBucketExtWithContext(ctx aws.Context, input *CreateBucketInput, opts ...request.Option) (*s3.CreateBucketOutput, error) {
	req, out := c.CreateBucketExtRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteBucketMetadataSearch = "DeleteBucketMetadataSearch"

// DeleteBucketMetadataSearchRequest generates a request.Request"
func (c *S3) DeleteBucketMetadataSearchRequest(input *DeleteBucketMetadataSearchInput) (req *request.Request, output *DeleteBucketMetadataSearchOutput) {
	op := &request.Operation{
		Name:       opDeleteBucketMetadataSearch,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}?searchmetadata",
	}

	if input == nil {
		input = &DeleteBucketMetadataSearchInput{}
	}

	output = &DeleteBucketMetadataSearchOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeleteBucketMetadataSearch API operation for ECS Extension.
func (c *S3) DeleteBucketMetadataSearch(input *DeleteBucketMetadataSearchInput) (*DeleteBucketMetadataSearchOutput, error) {
	req, out := c.DeleteBucketMetadataSearchRequest(input)
	return out, req.Send()
}

// DeleteBucketMetadataSearchWithContext is the same as DeleteBucketMetadataSearch with the addition of
// the ability to pass a context and additional request options.
func (c *S3) DeleteBucketMetadataSearchWithContext(ctx aws.Context, input *DeleteBucketMetadataSearchInput, opts ...request.Option) (*DeleteBucketMetadataSearchOutput, error) {
	req, out := c.DeleteBucketMetadataSearchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetObject = "GetObject"

// GetObjectExtRequest generates a request.Request
func (c *S3) GetObjectExtRequest(input *s3.GetObjectInput) (req *request.Request, output *GetObjectOutput) {
	op := &request.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &s3.GetObjectInput{}
	}

	output = &GetObjectOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetObjectExt API operation for ECS Extension.
func (c *S3) GetObjectExt(input *s3.GetObjectInput) (*GetObjectOutput, error) {
	req, out := c.GetObjectExtRequest(input)
	return out, req.Send()
}

// GetObjectExtWithContext is the same as GetObject with the addition of
// the ability to pass a context and additional request options.
func (c *S3) GetObjectExtWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (*GetObjectOutput, error) {
	req, out := c.GetObjectExtRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetSystemMetdataSearchKeys = "GetSystemMetdataSearchKeys"

// GetSystemMetdataSearchKeysRequest generates request.Request
func (c *S3) GetSystemMetdataSearchKeysRequest(input *GetSystemMetdataSearchKeysInput) (req *request.Request, output *GetSystemMetdataSearchKeysOutput) {
	op := &request.Operation{
		Name:       opGetSystemMetdataSearchKeys,
		HTTPMethod: "GET",
		HTTPPath:   "/?searchmetadata",
	}

	if input == nil {
		input = &GetSystemMetdataSearchKeysInput{}
	}

	output = &GetSystemMetdataSearchKeysOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetSystemMetdataSearchKeys API operation for ECS Extension.
func (c *S3) GetSystemMetdataSearchKeys(input *GetSystemMetdataSearchKeysInput) (*GetSystemMetdataSearchKeysOutput, error) {
	req, out := c.GetSystemMetdataSearchKeysRequest(input)
	return out, req.Send()
}

// GetSystemMetdataSearchKeysWithContext is the same as GetSystemMetdataSearchKeys with the addition of
// the ability to pass a context and additional request options.
func (c *S3) GetSystemMetdataSearchKeysWithContext(ctx aws.Context, input *GetSystemMetdataSearchKeysInput, opts ...request.Option) (*GetSystemMetdataSearchKeysOutput, error) {
	req, out := c.GetSystemMetdataSearchKeysRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opHeadBucket = "HeadBucket"

// HeadBucket API operation for ECS Extension
func (c *S3) HeadBucketExt(input *s3.HeadBucketInput) (*HeadBucketOutput, error) {
	req, out := c.HeadBucketExtRequest(input)
	return out, req.Send()
}

// HeadBucketExtRequest generates a request.Request
func (c *S3) HeadBucketExtRequest(input *s3.HeadBucketInput) (req *request.Request, output *HeadBucketOutput) {
	op := &request.Operation{
		Name:       opHeadBucket,
		HTTPMethod: "HEAD",
		HTTPPath:   "/{Bucket}",
	}

	if input == nil {
		input = &s3.HeadBucketInput{}
	}

	output = &HeadBucketOutput{}
	req = c.newRequest(op, input, output)
	return
}

// HeadBucketExt API operation for ECS Extension.
func (c *S3) HeadBucketExtg(input *s3.HeadBucketInput) (*HeadBucketOutput, error) {
	req, out := c.HeadBucketExtRequest(input)
	return out, req.Send()
}

// HeadBucketExtWithContext is the same as HeadBucket with the addition of
// the ability to pass a context and additional request options.
func (c *S3) HeadBucketExtWithContext(ctx aws.Context, input *s3.HeadBucketInput, opts ...request.Option) (*HeadBucketOutput, error) {
	req, out := c.HeadBucketExtRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opHeadObject = "HeadObject"

// HeadObjectExtRequest generates a request.Request
func (c *S3) HeadObjectExtRequest(input *s3.HeadObjectInput) (req *request.Request, output *HeadObjectOutput) {
	op := &request.Operation{
		Name:       opHeadObject,
		HTTPMethod: "HEAD",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &s3.HeadObjectInput{}
	}

	output = &HeadObjectOutput{}
	req = c.newRequest(op, input, output)
	return
}

// HeadObjectExt API operation for ECS Extension.
func (c *S3) HeadObjectExt(input *s3.HeadObjectInput) (*HeadObjectOutput, error) {
	req, out := c.HeadObjectExtRequest(input)
	return out, req.Send()
}

// HeadObjectExtWithContext is the same as HeadObject with the addition of
// the ability to pass a context and additional request options.
func (c *S3) HeadObjectExtWithContext(ctx aws.Context, input *s3.HeadObjectInput, opts ...request.Option) (*HeadObjectOutput, error) {
	req, out := c.HeadObjectExtRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListBucketMetadataSearch = "ListBucketMetadataSearch"

// ListBucketMetadataSearchRequest generates a request.Request
func (c *S3) ListBucketMetadataSearchRequest(input *ListBucketMetadataSearchInput) (req *request.Request, output *ListBucketMetadataSearchOutput) {
	op := &request.Operation{
		Name:       opListBucketMetadataSearch,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?searchmetadata",
	}

	if input == nil {
		input = &ListBucketMetadataSearchInput{}
	}

	output = &ListBucketMetadataSearchOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListBucketMetadataSearch API operation for ECS Extension.
func (c *S3) ListBucketMetadataSearch(input *ListBucketMetadataSearchInput) (*ListBucketMetadataSearchOutput, error) {
	req, out := c.ListBucketMetadataSearchRequest(input)
	return out, req.Send()
}

// ListBucketMetadataSearchWithContext is the same as ListBucketMetadataSearch with the addition of
// the ability to pass a context and additional request options.
func (c *S3) ListBucketMetadataSearchWithContext(ctx aws.Context, input *ListBucketMetadataSearchInput, opts ...request.Option) (*ListBucketMetadataSearchOutput, error) {
	req, out := c.ListBucketMetadataSearchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListBucketQuery = "ListBucketQuery"

// ListBucketQueryRequest generates a request.Request
func (c *S3) ListBucketQueryRequest(input *ListBucketQueryInput) (req *request.Request, output *ListBucketQueryOutput) {
	op := &request.Operation{
		Name:       opListBucketQuery,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?query",
	}

	if input == nil {
		input = &ListBucketQueryInput{}
	}

	output = &ListBucketQueryOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListBucketQuery API operation for ECS Extension.
func (c *S3) ListBucketQuery(input *ListBucketQueryInput) (*ListBucketQueryOutput, error) {
	req, out := c.ListBucketQueryRequest(input)
	return out, req.Send()
}

// ListBucketQueryWithContext is the same as ListBucketQuery with the addition of
// the ability to pass a context and additional request options.
func (c *S3) ListBucketQueryWithContext(ctx aws.Context, input *ListBucketQueryInput, opts ...request.Option) (*ListBucketQueryOutput, error) {
	req, out := c.ListBucketQueryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutBucketIsStaleAllowed = "PutBucketIsStaleAllowed"

// PutBucketIsStaleAllowedRequest generates request.Request
func (c *S3) PutBucketIsStaleAllowedRequest(input *PutBucketIsStaleAllowedInput) (req *request.Request, output *PutBucketIsStaleAllowedOutput) {
	op := &request.Operation{
		Name:       opPutBucketIsStaleAllowed,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?isstaleallowed",
	}

	if input == nil {
		input = &PutBucketIsStaleAllowedInput{}
	}

	output = &PutBucketIsStaleAllowedOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutBucketIsStaleAllowed API operation for ECS Extension
func (c *S3) PutBucketIsStaleAllowed(input *PutBucketIsStaleAllowedInput) (*PutBucketIsStaleAllowedOutput, error) {
	req, out := c.PutBucketIsStaleAllowedRequest(input)
	return out, req.Send()
}

// PutBucketIsStaleAllowedWithContext is the same as PutBucketIsStaleAllowed with the addition of
// the ability to pass a context and additional request options.
func (c *S3) PutBucketIsStaleAllowedWithContext(ctx aws.Context, input *PutBucketIsStaleAllowedInput, opts ...request.Option) (*PutBucketIsStaleAllowedOutput, error) {
	req, out := c.PutBucketIsStaleAllowedRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutObject = "PutObject"

// PutObjectExtRequest generates a request.Request
func (c *S3) PutObjectExtRequest(input *PutObjectInput) (req *request.Request, output *PutObjectOutput) {
	op := &request.Operation{
		Name:       opPutObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &PutObjectInput{}
	}

	output = &PutObjectOutput{}
	req = c.newRequest(op, input, output)
	return
}

// PutObjectExt API operation for ECS Extension.
func (c *S3) PutObjectExt(input *PutObjectInput) (*PutObjectOutput, error) {
	req, out := c.PutObjectExtRequest(input)
	return out, req.Send()
}

// PutObjectExtWithContext is the same as PutObject with the addition of
// the ability to pass a context and additional request options.
func (c *S3) PutObjectExtWithContext(ctx aws.Context, input *PutObjectInput, opts ...request.Option) (*PutObjectOutput, error) {
	req, out := c.PutObjectExtRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateBucketInput struct {
	_ struct{} `type:"structure" payload:"CreateBucketConfiguration"`

	// The canned ACL to apply to the bucket.
	ACL *string `location:"header" locationName:"x-amz-acl" type:"string" enum:"BucketCannedACL"`
	// Bucket is a required field
	Bucket                    *string                       `location:"uri" locationName:"Bucket" type:"string" required:"true"`
	ComplianceEnabled         *bool                         `location:"header" locationName:"x-emc-compliance-enabled" type:"boolean"`
	CreateBucketConfiguration *s3.CreateBucketConfiguration `locationName:"CreateBucketConfiguration" type:"structure"`
	FileSystemAccess          *bool                         `location:"header" locationName:"x-emc-file-system-access-enabled" type:"boolean"`
	// Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`
	// Allows grantee to list the objects in the bucket.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`
	// Allows grantee to read the bucket ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`
	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string `location:"header" locationName:"x-amz-grant-write" type:"string"`
	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP   *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`
	IsStaleAllowed  *bool   `location:"header" locationName:"x-emc-is-stale-allowed" type:"boolean"`
	MetadataSearch  *string `location:"header" locationName:"x-emc-metadata-search" type:"string"`
	NameSpace       *string `location:"header" locationName:"x-emc-namespace" type:"string"`
	RetentionPeriod *int64  `location:"header" locationName:"x-emc-retention-period" type:"integer"`
	SSEEnabled      *bool   `location:"header" locationName:"x-emc-server-side-encryption-enabled" type:"boolean"`
	VPool           *string `location:"header" locationName:"x-emc-vpool" type:"string"`
}

// String returns the string representation
func (s CreateBucketInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateBucketInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBucketInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateBucketInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetACL sets the ACL field's value.
func (s *CreateBucketInput) SetACL(v string) *CreateBucketInput {
	s.ACL = &v
	return s
}

// SetBucket sets the Bucket field's value.
func (s *CreateBucketInput) SetBucket(v string) *CreateBucketInput {
	s.Bucket = &v
	return s
}

// SetComplianceEnabled sets the ComplianceEnabled field's value.
func (s *CreateBucketInput) SetComplianceEnabled(v bool) *CreateBucketInput {
	s.ComplianceEnabled = &v
	return s
}

// SetCreateBucketConfiguration sets the CreateBucketConfiguration field's value.
func (s *CreateBucketInput) SetCreateBucketConfiguration(v *s3.CreateBucketConfiguration) *CreateBucketInput {
	s.CreateBucketConfiguration = v
	return s
}

// SetFileSystemAccess sets the FileSystemAccess field's value.
func (s *CreateBucketInput) SetFileSystemAccess(v bool) *CreateBucketInput {
	s.FileSystemAccess = &v
	return s
}

// SetGrantFullControl sets the GrantFullControl field's value.
func (s *CreateBucketInput) SetGrantFullControl(v string) *CreateBucketInput {
	s.GrantFullControl = &v
	return s
}

// SetGrantRead sets the GrantRead field's value.
func (s *CreateBucketInput) SetGrantRead(v string) *CreateBucketInput {
	s.GrantRead = &v
	return s
}

// SetGrantReadACP sets the GrantReadACP field's value.
func (s *CreateBucketInput) SetGrantReadACP(v string) *CreateBucketInput {
	s.GrantReadACP = &v
	return s
}

// SetGrantWrite sets the GrantWrite field's value.
func (s *CreateBucketInput) SetGrantWrite(v string) *CreateBucketInput {
	s.GrantWrite = &v
	return s
}

// SetGrantWriteACP sets the GrantWriteACP field's value.
func (s *CreateBucketInput) SetGrantWriteACP(v string) *CreateBucketInput {
	s.GrantWriteACP = &v
	return s
}

// SetIsStaleAllowed sets the IsStaleAllowed field's value.
func (s *CreateBucketInput) SetIsStaleAllowed(v bool) *CreateBucketInput {
	s.IsStaleAllowed = &v
	return s
}

// SetMetadataSearch sets the MetadataSearch field's value.
func (s *CreateBucketInput) SetMetadataSearch(v string) *CreateBucketInput {
	s.MetadataSearch = &v
	return s
}

// SetNameSpace sets the NameSpace field's value.
func (s *CreateBucketInput) SetNameSpace(v string) *CreateBucketInput {
	s.NameSpace = &v
	return s
}

// SetRetentionPeriod sets the RetentionPeriod field's value.
func (s *CreateBucketInput) SetRetentionPeriod(v int64) *CreateBucketInput {
	s.RetentionPeriod = &v
	return s
}

// SetSSEEnabled sets the SSEEnabled field's value.
func (s *CreateBucketInput) SetSSEEnabled(v bool) *CreateBucketInput {
	s.SSEEnabled = &v
	return s
}

// SetVPool sets the VPool field's value.
func (s *CreateBucketInput) SetVPool(v string) *CreateBucketInput {
	s.VPool = &v
	return s
}

type DeleteBucketMetadataSearchInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBucketMetadataSearchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBucketMetadataSearchInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBucketMetadataSearchInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBucketMetadataSearchInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBucket sets the Bucket field's value.
func (s *DeleteBucketMetadataSearchInput) SetBucket(v string) *DeleteBucketMetadataSearchInput {
	s.Bucket = &v
	return s
}

type DeleteBucketMetadataSearchOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBucketMetadataSearchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBucketMetadataSearchOutput) GoString() string {
	return s.String()
}

type EmcEcsExtIndexableKey struct {
	_ struct{} `type:"structure"`

	Datatype *string `type:"string"`
	Name     *string `type:"string"`
}

// String returns the string representation
func (s EmcEcsExtIndexableKey) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmcEcsExtIndexableKey) GoString() string {
	return s.String()
}

// SetDatatype sets the Datatype field's value.
func (s *EmcEcsExtIndexableKey) SetDatatype(v string) *EmcEcsExtIndexableKey {
	s.Datatype = &v
	return s
}

// SetName sets the Name field's value.
func (s *EmcEcsExtIndexableKey) SetName(v string) *EmcEcsExtIndexableKey {
	s.Name = &v
	return s
}

type EmcEcsExtObjectMatch struct {
	_ struct{} `type:"structure"`

	IndexKey        *string              `locationName:"indexKey" type:"string"`
	ObjectId        *string              `locationName:"objectId" type:"string"`
	ObjectName      *string              `locationName:"objectName" type:"string"`
	ObjectOwnerZone *string              `locationName:"objectOwnerZone" type:"string"`
	QueryMds        []*EmcEcsExtQueryMds `locationName:"queryMds" type:"list" flattened:"true"`
	VersionId       *string              `locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s EmcEcsExtObjectMatch) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmcEcsExtObjectMatch) GoString() string {
	return s.String()
}

// SetIndexKey sets the IndexKey field's value.
func (s *EmcEcsExtObjectMatch) SetIndexKey(v string) *EmcEcsExtObjectMatch {
	s.IndexKey = &v
	return s
}

// SetObjectId sets the ObjectId field's value.
func (s *EmcEcsExtObjectMatch) SetObjectId(v string) *EmcEcsExtObjectMatch {
	s.ObjectId = &v
	return s
}

// SetObjectName sets the ObjectName field's value.
func (s *EmcEcsExtObjectMatch) SetObjectName(v string) *EmcEcsExtObjectMatch {
	s.ObjectName = &v
	return s
}

// SetObjectOwnerZone sets the ObjectOwnerZone field's value.
func (s *EmcEcsExtObjectMatch) SetObjectOwnerZone(v string) *EmcEcsExtObjectMatch {
	s.ObjectOwnerZone = &v
	return s
}

// SetQueryMds sets the QueryMds field's value.
func (s *EmcEcsExtObjectMatch) SetQueryMds(v []*EmcEcsExtQueryMds) *EmcEcsExtObjectMatch {
	s.QueryMds = v
	return s
}

// SetVersionId sets the VersionId field's value.
func (s *EmcEcsExtObjectMatch) SetVersionId(v string) *EmcEcsExtObjectMatch {
	s.VersionId = &v
	return s
}

type EmcEcsExtOptionalAttribute struct {
	_ struct{} `type:"structure"`

	Datatype *string `type:"string"`
	Name     *string `type:"string"`
}

// String returns the string representation
func (s EmcEcsExtOptionalAttribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmcEcsExtOptionalAttribute) GoString() string {
	return s.String()
}

// SetDatatype sets the Datatype field's value.
func (s *EmcEcsExtOptionalAttribute) SetDatatype(v string) *EmcEcsExtOptionalAttribute {
	s.Datatype = &v
	return s
}

// SetName sets the Name field's value.
func (s *EmcEcsExtOptionalAttribute) SetName(v string) *EmcEcsExtOptionalAttribute {
	s.Name = &v
	return s
}

type EmcEcsExtQueryMds struct {
	_ struct{} `type:"structure"`

	MdMap  map[string]*string `locationName:"mdMap" type:"map"`
	MdType *string            `locationName:"type" type:"string"`
}

// String returns the string representation
func (s EmcEcsExtQueryMds) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EmcEcsExtQueryMds) GoString() string {
	return s.String()
}

// SetMdMap sets the MdMap field's value.
func (s *EmcEcsExtQueryMds) SetMdMap(v map[string]*string) *EmcEcsExtQueryMds {
	s.MdMap = v
	return s
}

// SetMdType sets the MdType field's value.
func (s *EmcEcsExtQueryMds) SetMdType(v string) *EmcEcsExtQueryMds {
	s.MdType = &v
	return s
}

type GetObjectOutput struct {
	_ struct{} `type:"structure" payload:"Body"`

	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`
	// Object data.
	Body io.ReadCloser `type:"blob"`
	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`
	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`
	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`
	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`
	// Size of the body in bytes.
	ContentLength *int64  `location:"header" locationName:"Content-Length" type:"long"`
	ContentMD5EMC *string `location:"header" locationName:"x-emc-content-md5" type:"string"`
	// The portion of the object returned in the response.
	ContentRange *string `location:"header" locationName:"Content-Range" type:"string"`
	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`
	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL
	ETag *string `location:"header" locationName:"ETag" type:"string"`
	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`
	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`
	// Last modified date of the object
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp" timestampFormat:"rfc822"`
	// A map of metadata to store with the object in S3.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`
	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`
	// The count of parts this object has.
	PartsCount         *int64  `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`
	ReplicationStatus  *string `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"ReplicationStatus"`
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`
	// Provides information about object restoration operation and expiration time
	// of the restored object copy.
	Restore         *string `location:"header" locationName:"x-amz-restore" type:"string"`
	RetentionPeriod *int64  `location:"header" locationName:"x-emc-retention-period" type:"integer"`
	RetentionPolicy *string `location:"header" locationName:"x-emc-retention-policy" type:"string"`
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`
	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`
	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`
	StorageClass         *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`
	// The number of tags, if any, on the object.
	TagCount *int64 `location:"header" locationName:"x-amz-tagging-count" type:"integer"`
	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s GetObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetObjectOutput) GoString() string {
	return s.String()
}

// SetAcceptRanges sets the AcceptRanges field's value.
func (s *GetObjectOutput) SetAcceptRanges(v string) *GetObjectOutput {
	s.AcceptRanges = &v
	return s
}

// SetBody sets the Body field's value.
func (s *GetObjectOutput) SetBody(v io.ReadCloser) *GetObjectOutput {
	s.Body = v
	return s
}

// SetCacheControl sets the CacheControl field's value.
func (s *GetObjectOutput) SetCacheControl(v string) *GetObjectOutput {
	s.CacheControl = &v
	return s
}

// SetContentDisposition sets the ContentDisposition field's value.
func (s *GetObjectOutput) SetContentDisposition(v string) *GetObjectOutput {
	s.ContentDisposition = &v
	return s
}

// SetContentEncoding sets the ContentEncoding field's value.
func (s *GetObjectOutput) SetContentEncoding(v string) *GetObjectOutput {
	s.ContentEncoding = &v
	return s
}

// SetContentLanguage sets the ContentLanguage field's value.
func (s *GetObjectOutput) SetContentLanguage(v string) *GetObjectOutput {
	s.ContentLanguage = &v
	return s
}

// SetContentLength sets the ContentLength field's value.
func (s *GetObjectOutput) SetContentLength(v int64) *GetObjectOutput {
	s.ContentLength = &v
	return s
}

// SetContentMD5EMC sets the ContentMD5EMC field's value.
func (s *GetObjectOutput) SetContentMD5EMC(v string) *GetObjectOutput {
	s.ContentMD5EMC = &v
	return s
}

// SetContentRange sets the ContentRange field's value.
func (s *GetObjectOutput) SetContentRange(v string) *GetObjectOutput {
	s.ContentRange = &v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *GetObjectOutput) SetContentType(v string) *GetObjectOutput {
	s.ContentType = &v
	return s
}

// SetDeleteMarker sets the DeleteMarker field's value.
func (s *GetObjectOutput) SetDeleteMarker(v bool) *GetObjectOutput {
	s.DeleteMarker = &v
	return s
}

// SetETag sets the ETag field's value.
func (s *GetObjectOutput) SetETag(v string) *GetObjectOutput {
	s.ETag = &v
	return s
}

// SetExpiration sets the Expiration field's value.
func (s *GetObjectOutput) SetExpiration(v string) *GetObjectOutput {
	s.Expiration = &v
	return s
}

// SetExpires sets the Expires field's value.
func (s *GetObjectOutput) SetExpires(v string) *GetObjectOutput {
	s.Expires = &v
	return s
}

// SetLastModified sets the LastModified field's value.
func (s *GetObjectOutput) SetLastModified(v time.Time) *GetObjectOutput {
	s.LastModified = &v
	return s
}

// SetMetadata sets the Metadata field's value.
func (s *GetObjectOutput) SetMetadata(v map[string]*string) *GetObjectOutput {
	s.Metadata = v
	return s
}

// SetMissingMeta sets the MissingMeta field's value.
func (s *GetObjectOutput) SetMissingMeta(v int64) *GetObjectOutput {
	s.MissingMeta = &v
	return s
}

// SetPartsCount sets the PartsCount field's value.
func (s *GetObjectOutput) SetPartsCount(v int64) *GetObjectOutput {
	s.PartsCount = &v
	return s
}

// SetReplicationStatus sets the ReplicationStatus field's value.
func (s *GetObjectOutput) SetReplicationStatus(v string) *GetObjectOutput {
	s.ReplicationStatus = &v
	return s
}

// SetRequestCharged sets the RequestCharged field's value.
func (s *GetObjectOutput) SetRequestCharged(v string) *GetObjectOutput {
	s.RequestCharged = &v
	return s
}

// SetRestore sets the Restore field's value.
func (s *GetObjectOutput) SetRestore(v string) *GetObjectOutput {
	s.Restore = &v
	return s
}

// SetRetentionPeriod sets the RetentionPeriod field's value.
func (s *GetObjectOutput) SetRetentionPeriod(v int64) *GetObjectOutput {
	s.RetentionPeriod = &v
	return s
}

// SetRetentionPolicy sets the RetentionPolicy field's value.
func (s *GetObjectOutput) SetRetentionPolicy(v string) *GetObjectOutput {
	s.RetentionPolicy = &v
	return s
}

// SetSSECustomerAlgorithm sets the SSECustomerAlgorithm field's value.
func (s *GetObjectOutput) SetSSECustomerAlgorithm(v string) *GetObjectOutput {
	s.SSECustomerAlgorithm = &v
	return s
}

// SetSSECustomerKeyMD5 sets the SSECustomerKeyMD5 field's value.
func (s *GetObjectOutput) SetSSECustomerKeyMD5(v string) *GetObjectOutput {
	s.SSECustomerKeyMD5 = &v
	return s
}

// SetSSEKMSKeyId sets the SSEKMSKeyId field's value.
func (s *GetObjectOutput) SetSSEKMSKeyId(v string) *GetObjectOutput {
	s.SSEKMSKeyId = &v
	return s
}

// SetServerSideEncryption sets the ServerSideEncryption field's value.
func (s *GetObjectOutput) SetServerSideEncryption(v string) *GetObjectOutput {
	s.ServerSideEncryption = &v
	return s
}

// SetStorageClass sets the StorageClass field's value.
func (s *GetObjectOutput) SetStorageClass(v string) *GetObjectOutput {
	s.StorageClass = &v
	return s
}

// SetTagCount sets the TagCount field's value.
func (s *GetObjectOutput) SetTagCount(v int64) *GetObjectOutput {
	s.TagCount = &v
	return s
}

// SetVersionId sets the VersionId field's value.
func (s *GetObjectOutput) SetVersionId(v string) *GetObjectOutput {
	s.VersionId = &v
	return s
}

// SetWebsiteRedirectLocation sets the WebsiteRedirectLocation field's value.
func (s *GetObjectOutput) SetWebsiteRedirectLocation(v string) *GetObjectOutput {
	s.WebsiteRedirectLocation = &v
	return s
}

type GetSystemMetdataSearchKeysInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSystemMetdataSearchKeysInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSystemMetdataSearchKeysInput) GoString() string {
	return s.String()
}

type GetSystemMetdataSearchKeysOutput struct {
	_ struct{} `type:"structure"`

	IndexableKeys      []*EmcEcsExtIndexableKey      `locationNameList:"Key" type:"list"`
	OptionalAttributes []*EmcEcsExtOptionalAttribute `locationNameList:"Attribute" type:"list"`
}

// String returns the string representation
func (s GetSystemMetdataSearchKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSystemMetdataSearchKeysOutput) GoString() string {
	return s.String()
}

// SetIndexableKeys sets the IndexableKeys field's value.
func (s *GetSystemMetdataSearchKeysOutput) SetIndexableKeys(v []*EmcEcsExtIndexableKey) *GetSystemMetdataSearchKeysOutput {
	s.IndexableKeys = v
	return s
}

// SetOptionalAttributes sets the OptionalAttributes field's value.
func (s *GetSystemMetdataSearchKeysOutput) SetOptionalAttributes(v []*EmcEcsExtOptionalAttribute) *GetSystemMetdataSearchKeysOutput {
	s.OptionalAttributes = v
	return s
}

type HeadBucketOutput struct {
	_ struct{} `type:"structure"`

	RetentionPeriod *int64 `location:"header" locationName:"x-emc-retention-period" type:"integer"`
}

// String returns the string representation
func (s HeadBucketOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HeadBucketOutput) GoString() string {
	return s.String()
}

// SetRetentionPeriod sets the RetentionPeriod field's value.
func (s *HeadBucketOutput) SetRetentionPeriod(v int64) *HeadBucketOutput {
	s.RetentionPeriod = &v
	return s
}

type HeadObjectOutput struct {
	_ struct{} `type:"structure"`

	AcceptRanges *string `location:"header" locationName:"accept-ranges" type:"string"`
	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`
	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`
	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`
	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`
	// Size of the body in bytes.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`
	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool `location:"header" locationName:"x-amz-delete-marker" type:"boolean"`
	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL
	ETag *string `location:"header" locationName:"ETag" type:"string"`
	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string `location:"header" locationName:"x-amz-expiration" type:"string"`
	// The date and time at which the object is no longer cacheable.
	Expires *string `location:"header" locationName:"Expires" type:"string"`
	// Last modified date of the object
	LastModified *time.Time `location:"header" locationName:"Last-Modified" type:"timestamp" timestampFormat:"rfc822"`
	// A map of metadata to store with the object in S3.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`
	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP,
	// you can create metadata whose values are not legal HTTP headers.
	MissingMeta *int64 `location:"header" locationName:"x-amz-missing-meta" type:"integer"`
	// The count of parts this object has.
	PartsCount        *int64  `location:"header" locationName:"x-amz-mp-parts-count" type:"integer"`
	ReplicationStatus *string `location:"header" locationName:"x-amz-replication-status" type:"string" enum:"ReplicationStatus"`
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`
	// Provides information about object restoration operation and expiration time
	// of the restored object copy.
	Restore         *string `location:"header" locationName:"x-amz-restore" type:"string"`
	RetentionPeriod *int64  `location:"header" locationName:"x-emc-retention-period" type:"integer"`
	RetentionPolicy *string `location:"header" locationName:"x-emc-retention-policy" type:"string"`
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`
	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`
	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`
	StorageClass         *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`
	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s HeadObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HeadObjectOutput) GoString() string {
	return s.String()
}

// SetAcceptRanges sets the AcceptRanges field's value.
func (s *HeadObjectOutput) SetAcceptRanges(v string) *HeadObjectOutput {
	s.AcceptRanges = &v
	return s
}

// SetCacheControl sets the CacheControl field's value.
func (s *HeadObjectOutput) SetCacheControl(v string) *HeadObjectOutput {
	s.CacheControl = &v
	return s
}

// SetContentDisposition sets the ContentDisposition field's value.
func (s *HeadObjectOutput) SetContentDisposition(v string) *HeadObjectOutput {
	s.ContentDisposition = &v
	return s
}

// SetContentEncoding sets the ContentEncoding field's value.
func (s *HeadObjectOutput) SetContentEncoding(v string) *HeadObjectOutput {
	s.ContentEncoding = &v
	return s
}

// SetContentLanguage sets the ContentLanguage field's value.
func (s *HeadObjectOutput) SetContentLanguage(v string) *HeadObjectOutput {
	s.ContentLanguage = &v
	return s
}

// SetContentLength sets the ContentLength field's value.
func (s *HeadObjectOutput) SetContentLength(v int64) *HeadObjectOutput {
	s.ContentLength = &v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *HeadObjectOutput) SetContentType(v string) *HeadObjectOutput {
	s.ContentType = &v
	return s
}

// SetDeleteMarker sets the DeleteMarker field's value.
func (s *HeadObjectOutput) SetDeleteMarker(v bool) *HeadObjectOutput {
	s.DeleteMarker = &v
	return s
}

// SetETag sets the ETag field's value.
func (s *HeadObjectOutput) SetETag(v string) *HeadObjectOutput {
	s.ETag = &v
	return s
}

// SetExpiration sets the Expiration field's value.
func (s *HeadObjectOutput) SetExpiration(v string) *HeadObjectOutput {
	s.Expiration = &v
	return s
}

// SetExpires sets the Expires field's value.
func (s *HeadObjectOutput) SetExpires(v string) *HeadObjectOutput {
	s.Expires = &v
	return s
}

// SetLastModified sets the LastModified field's value.
func (s *HeadObjectOutput) SetLastModified(v time.Time) *HeadObjectOutput {
	s.LastModified = &v
	return s
}

// SetMetadata sets the Metadata field's value.
func (s *HeadObjectOutput) SetMetadata(v map[string]*string) *HeadObjectOutput {
	s.Metadata = v
	return s
}

// SetMissingMeta sets the MissingMeta field's value.
func (s *HeadObjectOutput) SetMissingMeta(v int64) *HeadObjectOutput {
	s.MissingMeta = &v
	return s
}

// SetPartsCount sets the PartsCount field's value.
func (s *HeadObjectOutput) SetPartsCount(v int64) *HeadObjectOutput {
	s.PartsCount = &v
	return s
}

// SetReplicationStatus sets the ReplicationStatus field's value.
func (s *HeadObjectOutput) SetReplicationStatus(v string) *HeadObjectOutput {
	s.ReplicationStatus = &v
	return s
}

// SetRequestCharged sets the RequestCharged field's value.
func (s *HeadObjectOutput) SetRequestCharged(v string) *HeadObjectOutput {
	s.RequestCharged = &v
	return s
}

// SetRestore sets the Restore field's value.
func (s *HeadObjectOutput) SetRestore(v string) *HeadObjectOutput {
	s.Restore = &v
	return s
}

// SetRetentionPeriod sets the RetentionPeriod field's value.
func (s *HeadObjectOutput) SetRetentionPeriod(v int64) *HeadObjectOutput {
	s.RetentionPeriod = &v
	return s
}

// SetRetentionPolicy sets the RetentionPolicy field's value.
func (s *HeadObjectOutput) SetRetentionPolicy(v string) *HeadObjectOutput {
	s.RetentionPolicy = &v
	return s
}

// SetSSECustomerAlgorithm sets the SSECustomerAlgorithm field's value.
func (s *HeadObjectOutput) SetSSECustomerAlgorithm(v string) *HeadObjectOutput {
	s.SSECustomerAlgorithm = &v
	return s
}

// SetSSECustomerKeyMD5 sets the SSECustomerKeyMD5 field's value.
func (s *HeadObjectOutput) SetSSECustomerKeyMD5(v string) *HeadObjectOutput {
	s.SSECustomerKeyMD5 = &v
	return s
}

// SetSSEKMSKeyId sets the SSEKMSKeyId field's value.
func (s *HeadObjectOutput) SetSSEKMSKeyId(v string) *HeadObjectOutput {
	s.SSEKMSKeyId = &v
	return s
}

// SetServerSideEncryption sets the ServerSideEncryption field's value.
func (s *HeadObjectOutput) SetServerSideEncryption(v string) *HeadObjectOutput {
	s.ServerSideEncryption = &v
	return s
}

// SetStorageClass sets the StorageClass field's value.
func (s *HeadObjectOutput) SetStorageClass(v string) *HeadObjectOutput {
	s.StorageClass = &v
	return s
}

// SetVersionId sets the VersionId field's value.
func (s *HeadObjectOutput) SetVersionId(v string) *HeadObjectOutput {
	s.VersionId = &v
	return s
}

// SetWebsiteRedirectLocation sets the WebsiteRedirectLocation field's value.
func (s *HeadObjectOutput) SetWebsiteRedirectLocation(v string) *HeadObjectOutput {
	s.WebsiteRedirectLocation = &v
	return s
}

type ListBucketMetadataSearchInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s ListBucketMetadataSearchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBucketMetadataSearchInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBucketMetadataSearchInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListBucketMetadataSearchInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBucket sets the Bucket field's value.
func (s *ListBucketMetadataSearchInput) SetBucket(v string) *ListBucketMetadataSearchInput {
	s.Bucket = &v
	return s
}

type ListBucketMetadataSearchOutput struct {
	_ struct{} `type:"structure"`

	IndexableKeys         []*EmcEcsExtIndexableKey      `locationNameList:"Key" type:"list"`
	MetadataSearchEnabled *bool                         `type:"boolean"`
	OptionalAttributes    []*EmcEcsExtOptionalAttribute `locationNameList:"Attribute" type:"list"`
}

// String returns the string representation
func (s ListBucketMetadataSearchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBucketMetadataSearchOutput) GoString() string {
	return s.String()
}

// SetIndexableKeys sets the IndexableKeys field's value.
func (s *ListBucketMetadataSearchOutput) SetIndexableKeys(v []*EmcEcsExtIndexableKey) *ListBucketMetadataSearchOutput {
	s.IndexableKeys = v
	return s
}

// SetMetadataSearchEnabled sets the MetadataSearchEnabled field's value.
func (s *ListBucketMetadataSearchOutput) SetMetadataSearchEnabled(v bool) *ListBucketMetadataSearchOutput {
	s.MetadataSearchEnabled = &v
	return s
}

// SetOptionalAttributes sets the OptionalAttributes field's value.
func (s *ListBucketMetadataSearchOutput) SetOptionalAttributes(v []*EmcEcsExtOptionalAttribute) *ListBucketMetadataSearchOutput {
	s.OptionalAttributes = v
	return s
}

type ListBucketQueryInput struct {
	_ struct{} `type:"structure"`

	Attributes *string `location:"querystring" locationName:"attributes" type:"string"`
	// Bucket is a required field
	Bucket              *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
	IncludeOlderVersion *bool   `location:"querystring" locationName:"include-older-version" type:"boolean"`
	Marker              *string `location:"querystring" locationName:"marker" type:"string"`
	MaxKeys             *int64  `location:"querystring" locationName:"max-keys" type:"integer"`
	// Query is a required field
	Query  *string `location:"querystring" locationName:"query" type:"string" required:"true"`
	Sorted *string `location:"querystring" locationName:"sorted" type:"string"`
}

// String returns the string representation
func (s ListBucketQueryInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBucketQueryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBucketQueryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListBucketQueryInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if s.Query == nil {
		invalidParams.Add(request.NewErrParamRequired("Query"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *ListBucketQueryInput) SetAttributes(v string) *ListBucketQueryInput {
	s.Attributes = &v
	return s
}

// SetBucket sets the Bucket field's value.
func (s *ListBucketQueryInput) SetBucket(v string) *ListBucketQueryInput {
	s.Bucket = &v
	return s
}

// SetIncludeOlderVersion sets the IncludeOlderVersion field's value.
func (s *ListBucketQueryInput) SetIncludeOlderVersion(v bool) *ListBucketQueryInput {
	s.IncludeOlderVersion = &v
	return s
}

// SetMarker sets the Marker field's value.
func (s *ListBucketQueryInput) SetMarker(v string) *ListBucketQueryInput {
	s.Marker = &v
	return s
}

// SetMaxKeys sets the MaxKeys field's value.
func (s *ListBucketQueryInput) SetMaxKeys(v int64) *ListBucketQueryInput {
	s.MaxKeys = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListBucketQueryInput) SetQuery(v string) *ListBucketQueryInput {
	s.Query = &v
	return s
}

// SetSorted sets the Sorted field's value.
func (s *ListBucketQueryInput) SetSorted(v string) *ListBucketQueryInput {
	s.Sorted = &v
	return s
}

type ListBucketQueryOutput struct {
	_ struct{} `type:"structure"`

	MaxKeys       *int64                  `type:"integer"`
	Name          *string                 `type:"string"`
	NextMarker    *string                 `type:"string"`
	ObjectMatches []*EmcEcsExtObjectMatch `locationNameList:"object" type:"list"`
}

// String returns the string representation
func (s ListBucketQueryOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBucketQueryOutput) GoString() string {
	return s.String()
}

// SetMaxKeys sets the MaxKeys field's value.
func (s *ListBucketQueryOutput) SetMaxKeys(v int64) *ListBucketQueryOutput {
	s.MaxKeys = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListBucketQueryOutput) SetName(v string) *ListBucketQueryOutput {
	s.Name = &v
	return s
}

// SetNextMarker sets the NextMarker field's value.
func (s *ListBucketQueryOutput) SetNextMarker(v string) *ListBucketQueryOutput {
	s.NextMarker = &v
	return s
}

// SetObjectMatches sets the ObjectMatches field's value.
func (s *ListBucketQueryOutput) SetObjectMatches(v []*EmcEcsExtObjectMatch) *ListBucketQueryOutput {
	s.ObjectMatches = v
	return s
}

type PutBucketIsStaleAllowedInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket         *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
	IsStaleAllowed *bool   `location:"header" locationName:"x-emc-is-stale-allowed" type:"boolean"`
}

// String returns the string representation
func (s PutBucketIsStaleAllowedInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutBucketIsStaleAllowedInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketIsStaleAllowedInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutBucketIsStaleAllowedInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBucket sets the Bucket field's value.
func (s *PutBucketIsStaleAllowedInput) SetBucket(v string) *PutBucketIsStaleAllowedInput {
	s.Bucket = &v
	return s
}

// SetIsStaleAllowed sets the IsStaleAllowed field's value.
func (s *PutBucketIsStaleAllowedInput) SetIsStaleAllowed(v bool) *PutBucketIsStaleAllowedInput {
	s.IsStaleAllowed = &v
	return s
}

type PutBucketIsStaleAllowedOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketIsStaleAllowedOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutBucketIsStaleAllowedOutput) GoString() string {
	return s.String()
}

type PutObjectInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The canned ACL to apply to the object.
	ACL *string `location:"header" locationName:"x-amz-acl" type:"string" enum:"ObjectCannedACL"`
	// Object data.
	Body io.ReadSeeker `type:"blob"`
	// Name of the bucket to which the PUT operation was initiated.
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
	// Specifies caching behavior along the request/reply chain.
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`
	// Specifies presentational information for the object.
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`
	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`
	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`
	// Size of the body in bytes. This parameter is useful when the size of the
	// body cannot be determined automatically.
	ContentLength *int64 `location:"header" locationName:"Content-Length" type:"long"`
	// A standard MIME type describing the format of the object data.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`
	// The date and time at which the object is no longer cacheable.
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp" timestampFormat:"rfc822"`
	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`
	// Allows grantee to read the object data and its metadata.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`
	// Allows grantee to read the object ACL.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`
	// Allows grantee to write the ACL for the applicable object.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`
	IfNoneMatch   *string `location:"header" locationName:"If-None-Match" type:"string"`
	// Object key for which the PUT operation was initiated.
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`
	// A map of metadata to store with the object in S3.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`
	Range    *string            `location:"header" locationName:"Range" type:"string"`
	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer    *string `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"RequestPayer"`
	RetentionPeriod *int64  `location:"header" locationName:"x-emc-retention-period" type:"integer"`
	RetentionPolicy *string `location:"header" locationName:"x-emc-retention-policy" type:"string"`
	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// does not store the encryption key. The key must be appropriate for use with
	// the algorithm specified in the x-amz-server-side​-encryption​-customer-algorithm
	// header.
	SSECustomerKey *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string"`
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure the encryption
	// key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`
	// Specifies the AWS KMS key ID to use for object encryption. All GET and PUT
	// requests for an object protected by AWS KMS will fail if not made via SSL
	// or using SigV4. Documentation on configuring any of the officially supported
	// AWS SDKs and CLI can be found at http://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`
	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`
	// The type of storage to use for the object. Defaults to 'STANDARD'.
	StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`
	// The tag-set for the object. The tag-set must be encoded as URL Query parameters
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`
	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata.
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}

// String returns the string representation
func (s PutObjectInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutObjectInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutObjectInput"}
	if s.Bucket == nil {
		invalidParams.Add(request.NewErrParamRequired("Bucket"))
	}
	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetACL sets the ACL field's value.
func (s *PutObjectInput) SetACL(v string) *PutObjectInput {
	s.ACL = &v
	return s
}

// SetBody sets the Body field's value.
func (s *PutObjectInput) SetBody(v io.ReadSeeker) *PutObjectInput {
	s.Body = v
	return s
}

// SetBucket sets the Bucket field's value.
func (s *PutObjectInput) SetBucket(v string) *PutObjectInput {
	s.Bucket = &v
	return s
}

// SetCacheControl sets the CacheControl field's value.
func (s *PutObjectInput) SetCacheControl(v string) *PutObjectInput {
	s.CacheControl = &v
	return s
}

// SetContentDisposition sets the ContentDisposition field's value.
func (s *PutObjectInput) SetContentDisposition(v string) *PutObjectInput {
	s.ContentDisposition = &v
	return s
}

// SetContentEncoding sets the ContentEncoding field's value.
func (s *PutObjectInput) SetContentEncoding(v string) *PutObjectInput {
	s.ContentEncoding = &v
	return s
}

// SetContentLanguage sets the ContentLanguage field's value.
func (s *PutObjectInput) SetContentLanguage(v string) *PutObjectInput {
	s.ContentLanguage = &v
	return s
}

// SetContentLength sets the ContentLength field's value.
func (s *PutObjectInput) SetContentLength(v int64) *PutObjectInput {
	s.ContentLength = &v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *PutObjectInput) SetContentType(v string) *PutObjectInput {
	s.ContentType = &v
	return s
}

// SetExpires sets the Expires field's value.
func (s *PutObjectInput) SetExpires(v time.Time) *PutObjectInput {
	s.Expires = &v
	return s
}

// SetGrantFullControl sets the GrantFullControl field's value.
func (s *PutObjectInput) SetGrantFullControl(v string) *PutObjectInput {
	s.GrantFullControl = &v
	return s
}

// SetGrantRead sets the GrantRead field's value.
func (s *PutObjectInput) SetGrantRead(v string) *PutObjectInput {
	s.GrantRead = &v
	return s
}

// SetGrantReadACP sets the GrantReadACP field's value.
func (s *PutObjectInput) SetGrantReadACP(v string) *PutObjectInput {
	s.GrantReadACP = &v
	return s
}

// SetGrantWriteACP sets the GrantWriteACP field's value.
func (s *PutObjectInput) SetGrantWriteACP(v string) *PutObjectInput {
	s.GrantWriteACP = &v
	return s
}

// SetIfNoneMatch sets the IfNoneMatch field's value.
func (s *PutObjectInput) SetIfNoneMatch(v string) *PutObjectInput {
	s.IfNoneMatch = &v
	return s
}

// SetKey sets the Key field's value.
func (s *PutObjectInput) SetKey(v string) *PutObjectInput {
	s.Key = &v
	return s
}

// SetMetadata sets the Metadata field's value.
func (s *PutObjectInput) SetMetadata(v map[string]*string) *PutObjectInput {
	s.Metadata = v
	return s
}

// SetRange sets the Range field's value.
func (s *PutObjectInput) SetRange(v string) *PutObjectInput {
	s.Range = &v
	return s
}

// SetRequestPayer sets the RequestPayer field's value.
func (s *PutObjectInput) SetRequestPayer(v string) *PutObjectInput {
	s.RequestPayer = &v
	return s
}

// SetRetentionPeriod sets the RetentionPeriod field's value.
func (s *PutObjectInput) SetRetentionPeriod(v int64) *PutObjectInput {
	s.RetentionPeriod = &v
	return s
}

// SetRetentionPolicy sets the RetentionPolicy field's value.
func (s *PutObjectInput) SetRetentionPolicy(v string) *PutObjectInput {
	s.RetentionPolicy = &v
	return s
}

// SetSSECustomerAlgorithm sets the SSECustomerAlgorithm field's value.
func (s *PutObjectInput) SetSSECustomerAlgorithm(v string) *PutObjectInput {
	s.SSECustomerAlgorithm = &v
	return s
}

// SetSSECustomerKey sets the SSECustomerKey field's value.
func (s *PutObjectInput) SetSSECustomerKey(v string) *PutObjectInput {
	s.SSECustomerKey = &v
	return s
}

// SetSSECustomerKeyMD5 sets the SSECustomerKeyMD5 field's value.
func (s *PutObjectInput) SetSSECustomerKeyMD5(v string) *PutObjectInput {
	s.SSECustomerKeyMD5 = &v
	return s
}

// SetSSEKMSKeyId sets the SSEKMSKeyId field's value.
func (s *PutObjectInput) SetSSEKMSKeyId(v string) *PutObjectInput {
	s.SSEKMSKeyId = &v
	return s
}

// SetServerSideEncryption sets the ServerSideEncryption field's value.
func (s *PutObjectInput) SetServerSideEncryption(v string) *PutObjectInput {
	s.ServerSideEncryption = &v
	return s
}

// SetStorageClass sets the StorageClass field's value.
func (s *PutObjectInput) SetStorageClass(v string) *PutObjectInput {
	s.StorageClass = &v
	return s
}

// SetTagging sets the Tagging field's value.
func (s *PutObjectInput) SetTagging(v string) *PutObjectInput {
	s.Tagging = &v
	return s
}

// SetWebsiteRedirectLocation sets the WebsiteRedirectLocation field's value.
func (s *PutObjectInput) SetWebsiteRedirectLocation(v string) *PutObjectInput {
	s.WebsiteRedirectLocation = &v
	return s
}

type PutObjectOutput struct {
	_ struct{} `type:"structure"`

	ContentMD5EMC *string `location:"header" locationName:"x-emc-content-md5" type:"string"`
	// Entity tag for the uploaded object.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
	// If the object expiration is configured, this will contain the expiration
	// date (expiry-date) and rule ID (rule-id). The value of rule-id is URL encoded.
	Expiration         *string `location:"header" locationName:"x-amz-expiration" type:"string"`
	PreviousObjectSize *int64  `location:"header" locationName:"x-emc-previous-object-size" type:"integer"`
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged *string `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"RequestCharged"`
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm
	// used.
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`
	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`
	// If present, specifies the ID of the AWS Key Management Service (KMS) master
	// encryption key that was used for the object.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string"`
	// The Server-side encryption algorithm used when storing this object in S3
	// (e.g., AES256, aws:kms).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`
	// Version of the object.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s PutObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutObjectOutput) GoString() string {
	return s.String()
}

// SetContentMD5EMC sets the ContentMD5EMC field's value.
func (s *PutObjectOutput) SetContentMD5EMC(v string) *PutObjectOutput {
	s.ContentMD5EMC = &v
	return s
}

// SetETag sets the ETag field's value.
func (s *PutObjectOutput) SetETag(v string) *PutObjectOutput {
	s.ETag = &v
	return s
}

// SetExpiration sets the Expiration field's value.
func (s *PutObjectOutput) SetExpiration(v string) *PutObjectOutput {
	s.Expiration = &v
	return s
}

// SetPreviousObjectSize sets the PreviousObjectSize field's value.
func (s *PutObjectOutput) SetPreviousObjectSize(v int64) *PutObjectOutput {
	s.PreviousObjectSize = &v
	return s
}

// SetRequestCharged sets the RequestCharged field's value.
func (s *PutObjectOutput) SetRequestCharged(v string) *PutObjectOutput {
	s.RequestCharged = &v
	return s
}

// SetSSECustomerAlgorithm sets the SSECustomerAlgorithm field's value.
func (s *PutObjectOutput) SetSSECustomerAlgorithm(v string) *PutObjectOutput {
	s.SSECustomerAlgorithm = &v
	return s
}

// SetSSECustomerKeyMD5 sets the SSECustomerKeyMD5 field's value.
func (s *PutObjectOutput) SetSSECustomerKeyMD5(v string) *PutObjectOutput {
	s.SSECustomerKeyMD5 = &v
	return s
}

// SetSSEKMSKeyId sets the SSEKMSKeyId field's value.
func (s *PutObjectOutput) SetSSEKMSKeyId(v string) *PutObjectOutput {
	s.SSEKMSKeyId = &v
	return s
}

// SetServerSideEncryption sets the ServerSideEncryption field's value.
func (s *PutObjectOutput) SetServerSideEncryption(v string) *PutObjectOutput {
	s.ServerSideEncryption = &v
	return s
}

// SetVersionId sets the VersionId field's value.
func (s *PutObjectOutput) SetVersionId(v string) *PutObjectOutput {
	s.VersionId = &v
	return s
}

func defaultInitRequestFn(r *request.Request) {
	platformRequestHandlers(r)
	if r.Operation.Name == opCreateBucket {
		r.Handlers.Validate.PushFront(populateLocationConstraint)
	}
}

func platformRequestHandlers(r *request.Request) {
	if r.Operation.HTTPMethod == "PUT" {
		// 100-Continue should only be used on put requests.
		r.Handlers.Sign.PushBack(add100Continue)
	}
}

func add100Continue(r *request.Request) {
	if aws.BoolValue(r.Config.S3Disable100Continue) {
		return
	}
	if r.HTTPRequest.ContentLength < 1024*1024*2 {
		// Ignore requests smaller than 2MB. This helps prevent delaying
		// requests unnecessarily.
		return
	}

	r.HTTPRequest.Header.Set("Expect", "100-Continue")
}

func populateLocationConstraint(r *request.Request) {
	if r.ParamsFilled() && aws.StringValue(r.Config.Region) != "us-east-1" {
		in := r.Params.(*CreateBucketInput)
		if in.CreateBucketConfiguration == nil {
			r.Params = awsutil.CopyOf(r.Params)
			in = r.Params.(*CreateBucketInput)
			in.CreateBucketConfiguration = &s3.CreateBucketConfiguration{
				LocationConstraint: r.Config.Region,
			}
		}
	}
}
