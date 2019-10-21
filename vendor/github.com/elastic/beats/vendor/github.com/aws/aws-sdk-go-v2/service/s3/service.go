// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// S3 provides the API operation methods for making requests to
// Amazon Simple Storage Service. See this package's package overview docs
// for details on the service.
//
// S3 methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type S3 struct {
	*aws.Client

	// Service specific configurations. (codegen: service_specific_config.go)

	// Disables the S3 client from using the Expect: 100-Continue header to wait for
	// the service to respond with a 100 status code before sending the HTTP request
	// body.
	//
	// You should disable 100-Continue if you experience issues with proxies or third
	// party S3 compatible services.
	//
	// See http.Transport's ExpectContinueTimeout for information on adjusting the
	// continue wait timeout. https://golang.org/pkg/net/http/#Transport
	Disable100Continue bool

	// Forces the client to use path-style addressing for S3 API operations. By
	// default the S3 client will use virtual hosted bucket addressing when possible.
	// The S3 client will automatically fall back to path-style when the bucket name
	// is not DNS compatible.
	//
	// With ForcePathStyle
	//
	// 	https://s3.us-west-2.amazonaws.com/BUCKET/KEY
	//
	// Without ForcePathStyle
	//
	// 	https://BUCKET.s3.us-west-2.amazonaws.com/KEY
	//
	// See http://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
	ForcePathStyle bool

	// Enables S3 Accelerate feature for API operation that support S3 Accelerate.
	// For all operations compatible with S3 Accelerate will use the accelerate
	// endpoint for requests. Requests not compatible will fall back to normal S3
	// requests.
	//
	// The bucket must be enable for accelerate to be used with S3 client with
	// accelerate enabled. If the bucket is not enabled for accelerate an error will
	// be returned. The bucket name must be DNS compatible to also work with
	// accelerate.
	//
	UseAccelerate bool
}

// Used for custom client initialization logic
var initClient func(*S3)

// Used for custom request initialization logic
var initRequest func(*S3, *aws.Request)

// Service information constants
const (
	ServiceName = "s3"        // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the S3 client with a config.
//
// Example:
//     // Create a S3 client from just a config.
//     svc := s3.New(myConfig)
func New(config aws.Config) *S3 {
	var signingName string
	signingRegion := config.Region

	svc := &S3{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2006-03-01",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a S3 operation and runs any
// custom request initialization.
func (c *S3) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
