package query_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/private/protocol/xml/xmlutil"
	"github.com/aws/aws-sdk-go-v2/private/util"
)

var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = ioutil.Discard
var _ = util.Trim("")
var _ = url.Values{}
var _ = io.EOF
var _ = aws.String
var _ = fmt.Println
var _ = reflect.Value{}

func init() {
	protocol.RandReader = &awstesting.ZeroReader{}
}

// OutputService1ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService1ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService1ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService1ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService1ProtocolTest client from just a config.
//     svc := outputservice1protocoltest.New(myConfig)
//
//     // Create a OutputService1ProtocolTest client with additional configuration
//     svc := outputservice1protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService1ProtocolTest(config aws.Config) *OutputService1ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService1ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice1protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService1ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService1ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService1TestCaseOperation1 = "OperationName"

// OutputService1TestCaseOperation1Request is a API request type for the OutputService1TestCaseOperation1 API operation.
type OutputService1TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService1TestShapeOutputService1TestCaseOperation1Input
}

// Send marshals and sends the OutputService1TestCaseOperation1 API request.
func (r OutputService1TestCaseOperation1Request) Send() (*OutputService1TestShapeOutputService1TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService1TestShapeOutputService1TestCaseOperation1Output), nil
}

// OutputService1TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService1TestCaseOperation1Request method.
//    req := client.OutputService1TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService1ProtocolTest) OutputService1TestCaseOperation1Request(input *OutputService1TestShapeOutputService1TestCaseOperation1Input) OutputService1TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService1TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService1TestShapeOutputService1TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService1TestShapeOutputService1TestCaseOperation1Output{})
	return OutputService1TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService1TestShapeOutputService1TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService1TestShapeOutputService1TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Char *string `type:"character"`

	Double *float64 `type:"double"`

	FalseBool *bool `type:"boolean"`

	Float *float64 `type:"float"`

	Long *int64 `type:"long"`

	Num *int64 `locationName:"FooNum" type:"integer"`

	Str *string `type:"string"`

	Timestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	TrueBool *bool `type:"boolean"`
}

// SetChar sets the Char field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetChar(v string) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Char = &v
	return s
}

// SetDouble sets the Double field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetDouble(v float64) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Double = &v
	return s
}

// SetFalseBool sets the FalseBool field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetFalseBool(v bool) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.FalseBool = &v
	return s
}

// SetFloat sets the Float field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetFloat(v float64) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Float = &v
	return s
}

// SetLong sets the Long field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetLong(v int64) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Long = &v
	return s
}

// SetNum sets the Num field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetNum(v int64) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Num = &v
	return s
}

// SetStr sets the Str field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetStr(v string) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Str = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetTimestamp(v time.Time) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.Timestamp = &v
	return s
}

// SetTrueBool sets the TrueBool field's value.
func (s *OutputService1TestShapeOutputService1TestCaseOperation1Output) SetTrueBool(v bool) *OutputService1TestShapeOutputService1TestCaseOperation1Output {
	s.TrueBool = &v
	return s
}

// OutputService2ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService2ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService2ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService2ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService2ProtocolTest client from just a config.
//     svc := outputservice2protocoltest.New(myConfig)
//
//     // Create a OutputService2ProtocolTest client with additional configuration
//     svc := outputservice2protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService2ProtocolTest(config aws.Config) *OutputService2ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService2ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice2protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService2ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService2ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService2TestCaseOperation1 = "OperationName"

// OutputService2TestCaseOperation1Request is a API request type for the OutputService2TestCaseOperation1 API operation.
type OutputService2TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService2TestShapeOutputService2TestCaseOperation1Input
}

// Send marshals and sends the OutputService2TestCaseOperation1 API request.
func (r OutputService2TestCaseOperation1Request) Send() (*OutputService2TestShapeOutputService2TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService2TestShapeOutputService2TestCaseOperation1Output), nil
}

// OutputService2TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService2TestCaseOperation1Request method.
//    req := client.OutputService2TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService2ProtocolTest) OutputService2TestCaseOperation1Request(input *OutputService2TestShapeOutputService2TestCaseOperation1Input) OutputService2TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService2TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService2TestShapeOutputService2TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService2TestShapeOutputService2TestCaseOperation1Output{})
	return OutputService2TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService2TestShapeOutputService2TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService2TestShapeOutputService2TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Num *int64 `type:"integer"`

	Str *string `type:"string"`
}

// SetNum sets the Num field's value.
func (s *OutputService2TestShapeOutputService2TestCaseOperation1Output) SetNum(v int64) *OutputService2TestShapeOutputService2TestCaseOperation1Output {
	s.Num = &v
	return s
}

// SetStr sets the Str field's value.
func (s *OutputService2TestShapeOutputService2TestCaseOperation1Output) SetStr(v string) *OutputService2TestShapeOutputService2TestCaseOperation1Output {
	s.Str = &v
	return s
}

// OutputService3ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService3ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService3ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService3ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService3ProtocolTest client from just a config.
//     svc := outputservice3protocoltest.New(myConfig)
//
//     // Create a OutputService3ProtocolTest client with additional configuration
//     svc := outputservice3protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService3ProtocolTest(config aws.Config) *OutputService3ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService3ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice3protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService3ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService3ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService3TestCaseOperation1 = "OperationName"

// OutputService3TestCaseOperation1Request is a API request type for the OutputService3TestCaseOperation1 API operation.
type OutputService3TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService3TestShapeOutputService3TestCaseOperation1Input
}

// Send marshals and sends the OutputService3TestCaseOperation1 API request.
func (r OutputService3TestCaseOperation1Request) Send() (*OutputService3TestShapeOutputService3TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService3TestShapeOutputService3TestCaseOperation1Output), nil
}

// OutputService3TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService3TestCaseOperation1Request method.
//    req := client.OutputService3TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService3ProtocolTest) OutputService3TestCaseOperation1Request(input *OutputService3TestShapeOutputService3TestCaseOperation1Input) OutputService3TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService3TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService3TestShapeOutputService3TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService3TestShapeOutputService3TestCaseOperation1Output{})
	return OutputService3TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService3TestShapeOutputService3TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService3TestShapeOutputService3TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	// Blob is automatically base64 encoded/decoded by the SDK.
	Blob []byte `type:"blob"`
}

// SetBlob sets the Blob field's value.
func (s *OutputService3TestShapeOutputService3TestCaseOperation1Output) SetBlob(v []byte) *OutputService3TestShapeOutputService3TestCaseOperation1Output {
	s.Blob = v
	return s
}

// OutputService4ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService4ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService4ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService4ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService4ProtocolTest client from just a config.
//     svc := outputservice4protocoltest.New(myConfig)
//
//     // Create a OutputService4ProtocolTest client with additional configuration
//     svc := outputservice4protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService4ProtocolTest(config aws.Config) *OutputService4ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService4ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice4protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService4ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService4ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService4TestCaseOperation1 = "OperationName"

// OutputService4TestCaseOperation1Request is a API request type for the OutputService4TestCaseOperation1 API operation.
type OutputService4TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService4TestShapeOutputService4TestCaseOperation1Input
}

// Send marshals and sends the OutputService4TestCaseOperation1 API request.
func (r OutputService4TestCaseOperation1Request) Send() (*OutputService4TestShapeOutputService4TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService4TestShapeOutputService4TestCaseOperation1Output), nil
}

// OutputService4TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService4TestCaseOperation1Request method.
//    req := client.OutputService4TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService4ProtocolTest) OutputService4TestCaseOperation1Request(input *OutputService4TestShapeOutputService4TestCaseOperation1Input) OutputService4TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService4TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService4TestShapeOutputService4TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService4TestShapeOutputService4TestCaseOperation1Output{})
	return OutputService4TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService4TestShapeOutputService4TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService4TestShapeOutputService4TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []string `type:"list"`
}

// SetListMember sets the ListMember field's value.
func (s *OutputService4TestShapeOutputService4TestCaseOperation1Output) SetListMember(v []string) *OutputService4TestShapeOutputService4TestCaseOperation1Output {
	s.ListMember = v
	return s
}

// OutputService5ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService5ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService5ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService5ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService5ProtocolTest client from just a config.
//     svc := outputservice5protocoltest.New(myConfig)
//
//     // Create a OutputService5ProtocolTest client with additional configuration
//     svc := outputservice5protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService5ProtocolTest(config aws.Config) *OutputService5ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService5ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice5protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService5ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService5ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService5TestCaseOperation1 = "OperationName"

// OutputService5TestCaseOperation1Request is a API request type for the OutputService5TestCaseOperation1 API operation.
type OutputService5TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService5TestShapeOutputService5TestCaseOperation1Input
}

// Send marshals and sends the OutputService5TestCaseOperation1 API request.
func (r OutputService5TestCaseOperation1Request) Send() (*OutputService5TestShapeOutputService5TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService5TestShapeOutputService5TestCaseOperation1Output), nil
}

// OutputService5TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService5TestCaseOperation1Request method.
//    req := client.OutputService5TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService5ProtocolTest) OutputService5TestCaseOperation1Request(input *OutputService5TestShapeOutputService5TestCaseOperation1Input) OutputService5TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService5TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService5TestShapeOutputService5TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService5TestShapeOutputService5TestCaseOperation1Output{})
	return OutputService5TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService5TestShapeOutputService5TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService5TestShapeOutputService5TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []string `locationNameList:"item" type:"list"`
}

// SetListMember sets the ListMember field's value.
func (s *OutputService5TestShapeOutputService5TestCaseOperation1Output) SetListMember(v []string) *OutputService5TestShapeOutputService5TestCaseOperation1Output {
	s.ListMember = v
	return s
}

// OutputService6ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService6ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService6ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService6ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService6ProtocolTest client from just a config.
//     svc := outputservice6protocoltest.New(myConfig)
//
//     // Create a OutputService6ProtocolTest client with additional configuration
//     svc := outputservice6protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService6ProtocolTest(config aws.Config) *OutputService6ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService6ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice6protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService6ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService6ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService6TestCaseOperation1 = "OperationName"

// OutputService6TestCaseOperation1Request is a API request type for the OutputService6TestCaseOperation1 API operation.
type OutputService6TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService6TestShapeOutputService6TestCaseOperation1Input
}

// Send marshals and sends the OutputService6TestCaseOperation1 API request.
func (r OutputService6TestCaseOperation1Request) Send() (*OutputService6TestShapeOutputService6TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService6TestShapeOutputService6TestCaseOperation1Output), nil
}

// OutputService6TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService6TestCaseOperation1Request method.
//    req := client.OutputService6TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService6ProtocolTest) OutputService6TestCaseOperation1Request(input *OutputService6TestShapeOutputService6TestCaseOperation1Input) OutputService6TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService6TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService6TestShapeOutputService6TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService6TestShapeOutputService6TestCaseOperation1Output{})
	return OutputService6TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService6TestShapeOutputService6TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService6TestShapeOutputService6TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []string `type:"list" flattened:"true"`
}

// SetListMember sets the ListMember field's value.
func (s *OutputService6TestShapeOutputService6TestCaseOperation1Output) SetListMember(v []string) *OutputService6TestShapeOutputService6TestCaseOperation1Output {
	s.ListMember = v
	return s
}

// OutputService7ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService7ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService7ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService7ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService7ProtocolTest client from just a config.
//     svc := outputservice7protocoltest.New(myConfig)
//
//     // Create a OutputService7ProtocolTest client with additional configuration
//     svc := outputservice7protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService7ProtocolTest(config aws.Config) *OutputService7ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService7ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice7protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService7ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService7ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService7TestCaseOperation1 = "OperationName"

// OutputService7TestCaseOperation1Request is a API request type for the OutputService7TestCaseOperation1 API operation.
type OutputService7TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService7TestShapeOutputService7TestCaseOperation1Input
}

// Send marshals and sends the OutputService7TestCaseOperation1 API request.
func (r OutputService7TestCaseOperation1Request) Send() (*OutputService7TestShapeOutputService7TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService7TestShapeOutputService7TestCaseOperation1Output), nil
}

// OutputService7TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService7TestCaseOperation1Request method.
//    req := client.OutputService7TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService7ProtocolTest) OutputService7TestCaseOperation1Request(input *OutputService7TestShapeOutputService7TestCaseOperation1Input) OutputService7TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService7TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService7TestShapeOutputService7TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService7TestShapeOutputService7TestCaseOperation1Output{})
	return OutputService7TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService7TestShapeOutputService7TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService7TestShapeOutputService7TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	ListMember []string `type:"list" flattened:"true"`
}

// SetListMember sets the ListMember field's value.
func (s *OutputService7TestShapeOutputService7TestCaseOperation1Output) SetListMember(v []string) *OutputService7TestShapeOutputService7TestCaseOperation1Output {
	s.ListMember = v
	return s
}

// OutputService8ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService8ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService8ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService8ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService8ProtocolTest client from just a config.
//     svc := outputservice8protocoltest.New(myConfig)
//
//     // Create a OutputService8ProtocolTest client with additional configuration
//     svc := outputservice8protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService8ProtocolTest(config aws.Config) *OutputService8ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService8ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice8protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService8ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService8ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService8TestCaseOperation1 = "OperationName"

// OutputService8TestCaseOperation1Request is a API request type for the OutputService8TestCaseOperation1 API operation.
type OutputService8TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService8TestShapeOutputService8TestCaseOperation1Input
}

// Send marshals and sends the OutputService8TestCaseOperation1 API request.
func (r OutputService8TestCaseOperation1Request) Send() (*OutputService8TestShapeOutputService8TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService8TestShapeOutputService8TestCaseOperation1Output), nil
}

// OutputService8TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService8TestCaseOperation1Request method.
//    req := client.OutputService8TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService8ProtocolTest) OutputService8TestCaseOperation1Request(input *OutputService8TestShapeOutputService8TestCaseOperation1Input) OutputService8TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService8TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService8TestShapeOutputService8TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService8TestShapeOutputService8TestCaseOperation1Output{})
	return OutputService8TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService8TestShapeOutputService8TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService8TestShapeOutputService8TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	List []OutputService8TestShapeStructureShape `type:"list"`
}

// SetList sets the List field's value.
func (s *OutputService8TestShapeOutputService8TestCaseOperation1Output) SetList(v []OutputService8TestShapeStructureShape) *OutputService8TestShapeOutputService8TestCaseOperation1Output {
	s.List = v
	return s
}

type OutputService8TestShapeStructureShape struct {
	_ struct{} `type:"structure"`

	Bar *string `type:"string"`

	Baz *string `type:"string"`

	Foo *string `type:"string"`
}

// SetBar sets the Bar field's value.
func (s *OutputService8TestShapeStructureShape) SetBar(v string) *OutputService8TestShapeStructureShape {
	s.Bar = &v
	return s
}

// SetBaz sets the Baz field's value.
func (s *OutputService8TestShapeStructureShape) SetBaz(v string) *OutputService8TestShapeStructureShape {
	s.Baz = &v
	return s
}

// SetFoo sets the Foo field's value.
func (s *OutputService8TestShapeStructureShape) SetFoo(v string) *OutputService8TestShapeStructureShape {
	s.Foo = &v
	return s
}

// OutputService9ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService9ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService9ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService9ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService9ProtocolTest client from just a config.
//     svc := outputservice9protocoltest.New(myConfig)
//
//     // Create a OutputService9ProtocolTest client with additional configuration
//     svc := outputservice9protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService9ProtocolTest(config aws.Config) *OutputService9ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService9ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice9protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService9ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService9ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService9TestCaseOperation1 = "OperationName"

// OutputService9TestCaseOperation1Request is a API request type for the OutputService9TestCaseOperation1 API operation.
type OutputService9TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService9TestShapeOutputService9TestCaseOperation1Input
}

// Send marshals and sends the OutputService9TestCaseOperation1 API request.
func (r OutputService9TestCaseOperation1Request) Send() (*OutputService9TestShapeOutputService9TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService9TestShapeOutputService9TestCaseOperation1Output), nil
}

// OutputService9TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService9TestCaseOperation1Request method.
//    req := client.OutputService9TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService9ProtocolTest) OutputService9TestCaseOperation1Request(input *OutputService9TestShapeOutputService9TestCaseOperation1Input) OutputService9TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService9TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService9TestShapeOutputService9TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService9TestShapeOutputService9TestCaseOperation1Output{})
	return OutputService9TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService9TestShapeOutputService9TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService9TestShapeOutputService9TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	List []OutputService9TestShapeStructureShape `type:"list" flattened:"true"`
}

// SetList sets the List field's value.
func (s *OutputService9TestShapeOutputService9TestCaseOperation1Output) SetList(v []OutputService9TestShapeStructureShape) *OutputService9TestShapeOutputService9TestCaseOperation1Output {
	s.List = v
	return s
}

type OutputService9TestShapeStructureShape struct {
	_ struct{} `type:"structure"`

	Bar *string `type:"string"`

	Baz *string `type:"string"`

	Foo *string `type:"string"`
}

// SetBar sets the Bar field's value.
func (s *OutputService9TestShapeStructureShape) SetBar(v string) *OutputService9TestShapeStructureShape {
	s.Bar = &v
	return s
}

// SetBaz sets the Baz field's value.
func (s *OutputService9TestShapeStructureShape) SetBaz(v string) *OutputService9TestShapeStructureShape {
	s.Baz = &v
	return s
}

// SetFoo sets the Foo field's value.
func (s *OutputService9TestShapeStructureShape) SetFoo(v string) *OutputService9TestShapeStructureShape {
	s.Foo = &v
	return s
}

// OutputService10ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService10ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService10ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService10ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService10ProtocolTest client from just a config.
//     svc := outputservice10protocoltest.New(myConfig)
//
//     // Create a OutputService10ProtocolTest client with additional configuration
//     svc := outputservice10protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService10ProtocolTest(config aws.Config) *OutputService10ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService10ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice10protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService10ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService10ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService10TestCaseOperation1 = "OperationName"

// OutputService10TestCaseOperation1Request is a API request type for the OutputService10TestCaseOperation1 API operation.
type OutputService10TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService10TestShapeOutputService10TestCaseOperation1Input
}

// Send marshals and sends the OutputService10TestCaseOperation1 API request.
func (r OutputService10TestCaseOperation1Request) Send() (*OutputService10TestShapeOutputService10TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService10TestShapeOutputService10TestCaseOperation1Output), nil
}

// OutputService10TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService10TestCaseOperation1Request method.
//    req := client.OutputService10TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService10ProtocolTest) OutputService10TestCaseOperation1Request(input *OutputService10TestShapeOutputService10TestCaseOperation1Input) OutputService10TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService10TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService10TestShapeOutputService10TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService10TestShapeOutputService10TestCaseOperation1Output{})
	return OutputService10TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService10TestShapeOutputService10TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService10TestShapeOutputService10TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	List []string `locationNameList:"NamedList" type:"list" flattened:"true"`
}

// SetList sets the List field's value.
func (s *OutputService10TestShapeOutputService10TestCaseOperation1Output) SetList(v []string) *OutputService10TestShapeOutputService10TestCaseOperation1Output {
	s.List = v
	return s
}

// OutputService11ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService11ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService11ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService11ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService11ProtocolTest client from just a config.
//     svc := outputservice11protocoltest.New(myConfig)
//
//     // Create a OutputService11ProtocolTest client with additional configuration
//     svc := outputservice11protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService11ProtocolTest(config aws.Config) *OutputService11ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService11ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice11protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService11ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService11ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService11TestCaseOperation1 = "OperationName"

// OutputService11TestCaseOperation1Request is a API request type for the OutputService11TestCaseOperation1 API operation.
type OutputService11TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService11TestShapeOutputService11TestCaseOperation1Input
}

// Send marshals and sends the OutputService11TestCaseOperation1 API request.
func (r OutputService11TestCaseOperation1Request) Send() (*OutputService11TestShapeOutputService11TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService11TestShapeOutputService11TestCaseOperation1Output), nil
}

// OutputService11TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService11TestCaseOperation1Request method.
//    req := client.OutputService11TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService11ProtocolTest) OutputService11TestCaseOperation1Request(input *OutputService11TestShapeOutputService11TestCaseOperation1Input) OutputService11TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService11TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService11TestShapeOutputService11TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService11TestShapeOutputService11TestCaseOperation1Output{})
	return OutputService11TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService11TestShapeOutputService11TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService11TestShapeOutputService11TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]OutputService11TestShapeStructType `type:"map"`
}

// SetMap sets the Map field's value.
func (s *OutputService11TestShapeOutputService11TestCaseOperation1Output) SetMap(v map[string]OutputService11TestShapeStructType) *OutputService11TestShapeOutputService11TestCaseOperation1Output {
	s.Map = v
	return s
}

type OutputService11TestShapeStructType struct {
	_ struct{} `type:"structure"`

	Foo *string `locationName:"foo" type:"string"`
}

// SetFoo sets the Foo field's value.
func (s *OutputService11TestShapeStructType) SetFoo(v string) *OutputService11TestShapeStructType {
	s.Foo = &v
	return s
}

// OutputService12ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService12ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService12ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService12ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService12ProtocolTest client from just a config.
//     svc := outputservice12protocoltest.New(myConfig)
//
//     // Create a OutputService12ProtocolTest client with additional configuration
//     svc := outputservice12protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService12ProtocolTest(config aws.Config) *OutputService12ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService12ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice12protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService12ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService12ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService12TestCaseOperation1 = "OperationName"

// OutputService12TestCaseOperation1Request is a API request type for the OutputService12TestCaseOperation1 API operation.
type OutputService12TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService12TestShapeOutputService12TestCaseOperation1Input
}

// Send marshals and sends the OutputService12TestCaseOperation1 API request.
func (r OutputService12TestCaseOperation1Request) Send() (*OutputService12TestShapeOutputService12TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService12TestShapeOutputService12TestCaseOperation1Output), nil
}

// OutputService12TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService12TestCaseOperation1Request method.
//    req := client.OutputService12TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService12ProtocolTest) OutputService12TestCaseOperation1Request(input *OutputService12TestShapeOutputService12TestCaseOperation1Input) OutputService12TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService12TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService12TestShapeOutputService12TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService12TestShapeOutputService12TestCaseOperation1Output{})
	return OutputService12TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService12TestShapeOutputService12TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService12TestShapeOutputService12TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]string `type:"map" flattened:"true"`
}

// SetMap sets the Map field's value.
func (s *OutputService12TestShapeOutputService12TestCaseOperation1Output) SetMap(v map[string]string) *OutputService12TestShapeOutputService12TestCaseOperation1Output {
	s.Map = v
	return s
}

// OutputService13ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService13ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService13ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService13ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService13ProtocolTest client from just a config.
//     svc := outputservice13protocoltest.New(myConfig)
//
//     // Create a OutputService13ProtocolTest client with additional configuration
//     svc := outputservice13protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService13ProtocolTest(config aws.Config) *OutputService13ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService13ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice13protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService13ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService13ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService13TestCaseOperation1 = "OperationName"

// OutputService13TestCaseOperation1Request is a API request type for the OutputService13TestCaseOperation1 API operation.
type OutputService13TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService13TestShapeOutputService13TestCaseOperation1Input
}

// Send marshals and sends the OutputService13TestCaseOperation1 API request.
func (r OutputService13TestCaseOperation1Request) Send() (*OutputService13TestShapeOutputService13TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService13TestShapeOutputService13TestCaseOperation1Output), nil
}

// OutputService13TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService13TestCaseOperation1Request method.
//    req := client.OutputService13TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService13ProtocolTest) OutputService13TestCaseOperation1Request(input *OutputService13TestShapeOutputService13TestCaseOperation1Input) OutputService13TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService13TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService13TestShapeOutputService13TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService13TestShapeOutputService13TestCaseOperation1Output{})
	return OutputService13TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService13TestShapeOutputService13TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService13TestShapeOutputService13TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
}

// SetMap sets the Map field's value.
func (s *OutputService13TestShapeOutputService13TestCaseOperation1Output) SetMap(v map[string]string) *OutputService13TestShapeOutputService13TestCaseOperation1Output {
	s.Map = v
	return s
}

// OutputService14ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService14ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService14ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService14ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService14ProtocolTest client from just a config.
//     svc := outputservice14protocoltest.New(myConfig)
//
//     // Create a OutputService14ProtocolTest client with additional configuration
//     svc := outputservice14protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService14ProtocolTest(config aws.Config) *OutputService14ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService14ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice14protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService14ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService14ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService14TestCaseOperation1 = "OperationName"

// OutputService14TestCaseOperation1Request is a API request type for the OutputService14TestCaseOperation1 API operation.
type OutputService14TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService14TestShapeOutputService14TestCaseOperation1Input
}

// Send marshals and sends the OutputService14TestCaseOperation1 API request.
func (r OutputService14TestCaseOperation1Request) Send() (*OutputService14TestShapeOutputService14TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService14TestShapeOutputService14TestCaseOperation1Output), nil
}

// OutputService14TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService14TestCaseOperation1Request method.
//    req := client.OutputService14TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService14ProtocolTest) OutputService14TestCaseOperation1Request(input *OutputService14TestShapeOutputService14TestCaseOperation1Input) OutputService14TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService14TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService14TestShapeOutputService14TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService14TestShapeOutputService14TestCaseOperation1Output{})
	return OutputService14TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService14TestShapeOutputService14TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService14TestShapeOutputService14TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Map map[string]string `locationNameKey:"foo" locationNameValue:"bar" type:"map" flattened:"true"`
}

// SetMap sets the Map field's value.
func (s *OutputService14TestShapeOutputService14TestCaseOperation1Output) SetMap(v map[string]string) *OutputService14TestShapeOutputService14TestCaseOperation1Output {
	s.Map = v
	return s
}

// OutputService15ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService15ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService15ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService15ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService15ProtocolTest client from just a config.
//     svc := outputservice15protocoltest.New(myConfig)
//
//     // Create a OutputService15ProtocolTest client with additional configuration
//     svc := outputservice15protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService15ProtocolTest(config aws.Config) *OutputService15ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService15ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice15protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService15ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService15ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService15TestCaseOperation1 = "OperationName"

// OutputService15TestCaseOperation1Request is a API request type for the OutputService15TestCaseOperation1 API operation.
type OutputService15TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService15TestShapeOutputService15TestCaseOperation1Input
}

// Send marshals and sends the OutputService15TestCaseOperation1 API request.
func (r OutputService15TestCaseOperation1Request) Send() (*OutputService15TestShapeOutputService15TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService15TestShapeOutputService15TestCaseOperation1Output), nil
}

// OutputService15TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService15TestCaseOperation1Request method.
//    req := client.OutputService15TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService15ProtocolTest) OutputService15TestCaseOperation1Request(input *OutputService15TestShapeOutputService15TestCaseOperation1Input) OutputService15TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService15TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService15TestShapeOutputService15TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService15TestShapeOutputService15TestCaseOperation1Output{})
	return OutputService15TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService15TestShapeOutputService15TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService15TestShapeOutputService15TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	Foo *string `type:"string"`
}

// SetFoo sets the Foo field's value.
func (s *OutputService15TestShapeOutputService15TestCaseOperation1Output) SetFoo(v string) *OutputService15TestShapeOutputService15TestCaseOperation1Output {
	s.Foo = &v
	return s
}

// OutputService16ProtocolTest provides the API operation methods for making requests to
// . See this package's package overview docs
// for details on the service.
//
// OutputService16ProtocolTest methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type OutputService16ProtocolTest struct {
	*aws.Client
}

// New creates a new instance of the OutputService16ProtocolTest client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a OutputService16ProtocolTest client from just a config.
//     svc := outputservice16protocoltest.New(myConfig)
//
//     // Create a OutputService16ProtocolTest client with additional configuration
//     svc := outputservice16protocoltest.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func NewOutputService16ProtocolTest(config aws.Config) *OutputService16ProtocolTest {
	var signingName string
	signingRegion := config.Region

	svc := &OutputService16ProtocolTest{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   "outputservice16protocoltest",
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	return svc
}

// newRequest creates a new request for a OutputService16ProtocolTest operation and runs any
// custom request initialization.
func (c *OutputService16ProtocolTest) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	return req
}

const opOutputService16TestCaseOperation1 = "OperationName"

// OutputService16TestCaseOperation1Request is a API request type for the OutputService16TestCaseOperation1 API operation.
type OutputService16TestCaseOperation1Request struct {
	*aws.Request
	Input *OutputService16TestShapeOutputService16TestCaseOperation1Input
}

// Send marshals and sends the OutputService16TestCaseOperation1 API request.
func (r OutputService16TestCaseOperation1Request) Send() (*OutputService16TestShapeOutputService16TestCaseOperation1Output, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*OutputService16TestShapeOutputService16TestCaseOperation1Output), nil
}

// OutputService16TestCaseOperation1Request returns a request value for making API operation for
// .
//
//    // Example sending a request using the OutputService16TestCaseOperation1Request method.
//    req := client.OutputService16TestCaseOperation1Request(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *OutputService16ProtocolTest) OutputService16TestCaseOperation1Request(input *OutputService16TestShapeOutputService16TestCaseOperation1Input) OutputService16TestCaseOperation1Request {
	op := &aws.Operation{
		Name: opOutputService16TestCaseOperation1,

		HTTPPath: "/",
	}

	if input == nil {
		input = &OutputService16TestShapeOutputService16TestCaseOperation1Input{}
	}

	req := c.newRequest(op, input, &OutputService16TestShapeOutputService16TestCaseOperation1Output{})
	return OutputService16TestCaseOperation1Request{Request: req, Input: input}
}

type OutputService16TestShapeOutputService16TestCaseOperation1Input struct {
	_ struct{} `type:"structure"`
}

type OutputService16TestShapeOutputService16TestCaseOperation1Output struct {
	_ struct{} `type:"structure"`

	FooEnum OutputService16TestShapeEC2EnumType `type:"string" enum:"true"`

	ListEnums []OutputService16TestShapeEC2EnumType `type:"list"`
}

// SetFooEnum sets the FooEnum field's value.
func (s *OutputService16TestShapeOutputService16TestCaseOperation1Output) SetFooEnum(v OutputService16TestShapeEC2EnumType) *OutputService16TestShapeOutputService16TestCaseOperation1Output {
	s.FooEnum = v
	return s
}

// SetListEnums sets the ListEnums field's value.
func (s *OutputService16TestShapeOutputService16TestCaseOperation1Output) SetListEnums(v []OutputService16TestShapeEC2EnumType) *OutputService16TestShapeOutputService16TestCaseOperation1Output {
	s.ListEnums = v
	return s
}

type OutputService16TestShapeEC2EnumType string

// Enum values for OutputService16TestShapeEC2EnumType
const (
	EC2EnumTypeFoo OutputService16TestShapeEC2EnumType = "foo"
	EC2EnumTypeBar OutputService16TestShapeEC2EnumType = "bar"
)

//
// Tests begin here
//

func TestOutputService1ProtocolTestScalarMembersCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService1ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Str>myname</Str><FooNum>123</FooNum><FalseBool>false</FalseBool><TrueBool>true</TrueBool><Float>1.2</Float><Double>1.3</Double><Long>200</Long><Char>a</Char><Timestamp>2015-01-25T08:00:00Z</Timestamp></OperationNameResult><ResponseMetadata><RequestId>request-id</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService1TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService1TestShapeOutputService1TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "a", *out.Char; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.3, *out.Double; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := false, *out.FalseBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := 1.2, *out.Float; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(200), *out.Long; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := int64(123), *out.Num; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("myname"), *out.Str; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := time.Unix(1.4221728e+09, 0).UTC().String(), out.Timestamp.String(); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := true, *out.TrueBool; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService2ProtocolTestNotAllMembersInResponseCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService2ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Str>myname</Str></OperationNameResult><ResponseMetadata><RequestId>request-id</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService2TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService2TestShapeOutputService2TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("myname"), *out.Str; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService3ProtocolTestBlobCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService3ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Blob>dmFsdWU=</Blob></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService3TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService3TestShapeOutputService3TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := "value", string(out.Blob); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService4ProtocolTestListsCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService4ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember><member>abc</member><member>123</member></ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService4TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService4TestShapeOutputService4TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("123"), out.ListMember[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService5ProtocolTestListWithCustomMemberNameCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService5ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember><item>abc</item><item>123</item></ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService5TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService5TestShapeOutputService5TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("123"), out.ListMember[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService6ProtocolTestFlattenedListCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService6ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember>abc</ListMember><ListMember>123</ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService6TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService6TestShapeOutputService6TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("123"), out.ListMember[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService7ProtocolTestFlattenedSingleElementListCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService7ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><ListMember>abc</ListMember></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService7TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService7TestShapeOutputService7TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("abc"), out.ListMember[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService8ProtocolTestListOfStructuresCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService8ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08/\"><OperationNameResult><List><member><Foo>firstfoo</Foo><Bar>firstbar</Bar><Baz>firstbaz</Baz></member><member><Foo>secondfoo</Foo><Bar>secondbar</Bar><Baz>secondbaz</Baz></member></List></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService8TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService8TestShapeOutputService8TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("firstbar"), *out.List[0].Bar; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("firstbaz"), *out.List[0].Baz; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("firstfoo"), *out.List[0].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("secondbar"), *out.List[1].Bar; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("secondbaz"), *out.List[1].Baz; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("secondfoo"), *out.List[1].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService9ProtocolTestFlattenedListOfStructuresCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService9ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08/\"><OperationNameResult><List><Foo>firstfoo</Foo><Bar>firstbar</Bar><Baz>firstbaz</Baz></List><List><Foo>secondfoo</Foo><Bar>secondbar</Bar><Baz>secondbaz</Baz></List></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService9TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService9TestShapeOutputService9TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("firstbar"), *out.List[0].Bar; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("firstbaz"), *out.List[0].Baz; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("firstfoo"), *out.List[0].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("secondbar"), *out.List[1].Bar; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("secondbaz"), *out.List[1].Baz; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("secondfoo"), *out.List[1].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService10ProtocolTestFlattenedListWithLocationNameCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService10ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08/\"><OperationNameResult><NamedList>a</NamedList><NamedList>b</NamedList></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService10TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService10TestShapeOutputService10TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("a"), out.List[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("b"), out.List[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService11ProtocolTestNormalMapCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService11ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse xmlns=\"https://service.amazonaws.com/doc/2010-05-08\"><OperationNameResult><Map><entry><key>qux</key><value><foo>bar</foo></value></entry><entry><key>baz</key><value><foo>bam</foo></value></entry></Map></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService11TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService11TestShapeOutputService11TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bam"), *out.Map["baz"].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("bar"), *out.Map["qux"].Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService12ProtocolTestFlattenedMapCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService12ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Map><key>qux</key><value>bar</value></Map><Map><key>baz</key><value>bam</value></Map></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService12TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService12TestShapeOutputService12TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bam"), out.Map["baz"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("bar"), out.Map["qux"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService13ProtocolTestFlattenedMapInShapeDefinitionCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService13ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Attribute><Name>qux</Name><Value>bar</Value></Attribute></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService13TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService13TestShapeOutputService13TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bar"), out.Map["qux"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService14ProtocolTestNamedMapCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService14ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Map><foo>qux</foo><bar>bar</bar></Map><Map><foo>baz</foo><bar>bam</bar></Map></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService14TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService14TestShapeOutputService14TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string("bam"), out.Map["baz"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := string("bar"), out.Map["qux"]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService15ProtocolTestEmptyStringCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService15ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><OperationNameResult><Foo/></OperationNameResult><ResponseMetadata><RequestId>requestid</RequestId></ResponseMetadata></OperationNameResponse>"))
	req := svc.OutputService15TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService15TestShapeOutputService15TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := string(""), *out.Foo; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}

func TestOutputService16ProtocolTestEnumOutputCase1(t *testing.T) {
	cfg := unit.Config()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL("https://test")

	svc := NewOutputService16ProtocolTest(cfg)

	buf := bytes.NewReader([]byte("<OperationNameResponse><FooEnum>foo</FooEnum><ListEnums><member>foo</member><member>bar</member></ListEnums></OperationNameResponse>"))
	req := svc.OutputService16TestCaseOperation1Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers

	// unmarshal response
	query.UnmarshalMeta(req.Request)
	query.Unmarshal(req.Request)
	if req.Error != nil {
		t.Errorf("expect not error, got %v", req.Error)
	}

	out := req.Data.(*OutputService16TestShapeOutputService16TestCaseOperation1Output)
	// assert response
	if out == nil {
		t.Errorf("expect not to be nil")
	}
	if e, a := OutputService16TestShapeEC2EnumType("foo"), out.FooEnum; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService16TestShapeEC2EnumType("foo"), out.ListEnums[0]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := OutputService16TestShapeEC2EnumType("bar"), out.ListEnums[1]; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

}