// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

const opGetMediaForFragmentList = "GetMediaForFragmentList"

// GetMediaForFragmentListRequest is a API request type for the GetMediaForFragmentList API operation.
type GetMediaForFragmentListRequest struct {
	*aws.Request
	Input *GetMediaForFragmentListInput
	Copy  func(*GetMediaForFragmentListInput) GetMediaForFragmentListRequest
}

// Send marshals and sends the GetMediaForFragmentList API request.
func (r GetMediaForFragmentListRequest) Send() (*GetMediaForFragmentListOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetMediaForFragmentListOutput), nil
}

// GetMediaForFragmentListRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams Archived Media.
//
// Gets media for a list of fragments (specified by fragment number) from the
// archived data in a Kinesis video stream.
//
// This operation is only available for the AWS SDK for Java. It is not supported
// in AWS SDKs for other languages.
//
// The following limits apply when using the GetMediaForFragmentList API:
//
//    * A client can call GetMediaForFragmentList up to five times per second
//    per stream.
//
//    * Kinesis Video Streams sends media data at a rate of up to 25 megabytes
//    per second (or 200 megabits per second) during a GetMediaForFragmentList
//    session.
//
//    // Example sending a request using the GetMediaForFragmentListRequest method.
//    req := client.GetMediaForFragmentListRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentList
func (c *KinesisVideoArchivedMedia) GetMediaForFragmentListRequest(input *GetMediaForFragmentListInput) GetMediaForFragmentListRequest {
	op := &aws.Operation{
		Name:       opGetMediaForFragmentList,
		HTTPMethod: "POST",
		HTTPPath:   "/getMediaForFragmentList",
	}

	if input == nil {
		input = &GetMediaForFragmentListInput{}
	}

	output := &GetMediaForFragmentListOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetMediaForFragmentListRequest{Request: req, Input: input, Copy: c.GetMediaForFragmentListRequest}
}

const opListFragments = "ListFragments"

// ListFragmentsRequest is a API request type for the ListFragments API operation.
type ListFragmentsRequest struct {
	*aws.Request
	Input *ListFragmentsInput
	Copy  func(*ListFragmentsInput) ListFragmentsRequest
}

// Send marshals and sends the ListFragments API request.
func (r ListFragmentsRequest) Send() (*ListFragmentsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*ListFragmentsOutput), nil
}

// ListFragmentsRequest returns a request value for making API operation for
// Amazon Kinesis Video Streams Archived Media.
//
// Returns a list of Fragment objects from the specified stream and start location
// within the archived data.
//
//    // Example sending a request using the ListFragmentsRequest method.
//    req := client.ListFragmentsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/ListFragments
func (c *KinesisVideoArchivedMedia) ListFragmentsRequest(input *ListFragmentsInput) ListFragmentsRequest {
	op := &aws.Operation{
		Name:       opListFragments,
		HTTPMethod: "POST",
		HTTPPath:   "/listFragments",
	}

	if input == nil {
		input = &ListFragmentsInput{}
	}

	output := &ListFragmentsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return ListFragmentsRequest{Request: req, Input: input, Copy: c.ListFragmentsRequest}
}

// Represents a segment of video or other time-delimited data.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/Fragment
type Fragment struct {
	_ struct{} `type:"structure"`

	// The playback duration or other time value associated with the fragment.
	FragmentLengthInMilliseconds *int64 `type:"long"`

	// The index value of the fragment.
	FragmentNumber *string `min:"1" type:"string"`

	// The total fragment size, including information about the fragment and contained
	// media data.
	FragmentSizeInBytes *int64 `type:"long"`

	// The time stamp from the producer corresponding to the fragment.
	ProducerTimestamp *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The time stamp from the AWS server corresponding to the fragment.
	ServerTimestamp *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s Fragment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Fragment) GoString() string {
	return s.String()
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Fragment) MarshalFields(e protocol.FieldEncoder) error {
	if s.FragmentLengthInMilliseconds != nil {
		v := *s.FragmentLengthInMilliseconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FragmentLengthInMilliseconds", protocol.Int64Value(v), metadata)
	}
	if s.FragmentNumber != nil {
		v := *s.FragmentNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FragmentNumber", protocol.QuotedValue{protocol.StringValue(v)}, metadata)
	}
	if s.FragmentSizeInBytes != nil {
		v := *s.FragmentSizeInBytes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FragmentSizeInBytes", protocol.Int64Value(v), metadata)
	}
	if s.ProducerTimestamp != nil {
		v := *s.ProducerTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ProducerTimestamp", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.ServerTimestamp != nil {
		v := *s.ServerTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ServerTimestamp", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

// Describes the time stamp range and time stamp origin of a range of fragments.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/FragmentSelector
type FragmentSelector struct {
	_ struct{} `type:"structure"`

	// The origin of the time stamps to use (Server or Producer).
	//
	// FragmentSelectorType is a required field
	FragmentSelectorType FragmentSelectorType `type:"string" required:"true" enum:"true"`

	// The range of time stamps to return.
	//
	// TimestampRange is a required field
	TimestampRange *TimestampRange `type:"structure" required:"true"`
}

// String returns the string representation
func (s FragmentSelector) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FragmentSelector) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FragmentSelector) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FragmentSelector"}
	if len(s.FragmentSelectorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("FragmentSelectorType"))
	}

	if s.TimestampRange == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimestampRange"))
	}
	if s.TimestampRange != nil {
		if err := s.TimestampRange.Validate(); err != nil {
			invalidParams.AddNested("TimestampRange", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s FragmentSelector) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.FragmentSelectorType) > 0 {
		v := s.FragmentSelectorType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FragmentSelectorType", protocol.QuotedValue{v}, metadata)
	}
	if s.TimestampRange != nil {
		v := s.TimestampRange

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TimestampRange", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentListInput
type GetMediaForFragmentListInput struct {
	_ struct{} `type:"structure"`

	// A list of the numbers of fragments for which to retrieve media. You retrieve
	// these values with ListFragments.
	//
	// Fragments is a required field
	Fragments []string `type:"list" required:"true"`

	// The name of the stream from which to retrieve fragment media.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMediaForFragmentListInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMediaForFragmentListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMediaForFragmentListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMediaForFragmentListInput"}

	if s.Fragments == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fragments"))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMediaForFragmentListInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.Fragments) > 0 {
		v := s.Fragments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Fragments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/GetMediaForFragmentListOutput
type GetMediaForFragmentListOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	responseMetadata aws.Response

	// The content type of the requested media.
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string"`

	// The payload that Kinesis Video Streams returns is a sequence of chunks from
	// the specified stream. For information about the chunks, see PutMedia (docs.aws.amazon.com/acuity/latest/dg/API_dataplane_PutMedia.html).
	// The chunks that Kinesis Video Streams returns in the GetMediaForFragmentList
	// call also include the following additional Matroska (MKV) tags:
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - Fragment number returned in the chunk.
	//
	//    * AWS_KINESISVIDEO_SERVER_SIDE_TIMESTAMP - Server-side time stamp of the
	//    fragment.
	//
	//    * AWS_KINESISVIDEO_PRODUCER_SIDE_TIMESTAMP - Producer-side time stamp
	//    of the fragment.
	//
	// The following tags will be included if an exception occurs:
	//
	//    * AWS_KINESISVIDEO_FRAGMENT_NUMBER - The number of the fragment that threw
	//    the exception
	//
	//    * AWS_KINESISVIDEO_EXCEPTION_ERROR_CODE - The integer code of the exception
	//
	//    * AWS_KINESISVIDEO_EXCEPTION_MESSAGE - A text description of the exception
	Payload io.ReadCloser `type:"blob"`
}

// String returns the string representation
func (s GetMediaForFragmentListOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetMediaForFragmentListOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetMediaForFragmentListOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetMediaForFragmentListOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{protocol.StringValue(v)}, metadata)
	}
	// Skipping Payload Output type's body not valid.
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/ListFragmentsInput
type ListFragmentsInput struct {
	_ struct{} `type:"structure"`

	// Describes the time stamp range and time stamp origin for the range of fragments
	// to return.
	FragmentSelector *FragmentSelector `type:"structure"`

	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a ListFragmentsOutput$NextToken
	// is provided in the output that you can use to resume pagination.
	MaxResults *int64 `min:"1" type:"long"`

	// A token to specify where to start paginating. This is the ListFragmentsOutput$NextToken
	// from a previously truncated response.
	NextToken *string `min:"1" type:"string"`

	// The name of the stream from which to retrieve a fragment list.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListFragmentsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListFragmentsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFragmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFragmentsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}
	if s.FragmentSelector != nil {
		if err := s.FragmentSelector.Validate(); err != nil {
			invalidParams.AddNested("FragmentSelector", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFragmentsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FragmentSelector != nil {
		v := s.FragmentSelector

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "FragmentSelector", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{protocol.StringValue(v)}, metadata)
	}
	if s.StreamName != nil {
		v := *s.StreamName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StreamName", protocol.QuotedValue{protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/ListFragmentsOutput
type ListFragmentsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// A list of fragment numbers that correspond to the time stamp range provided.
	Fragments []Fragment `type:"list"`

	// If the returned list is truncated, the operation returns this token to use
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFragmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListFragmentsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s ListFragmentsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFragmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Fragments) > 0 {
		v := s.Fragments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Fragments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The range of time stamps for which to return fragments.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-video-archived-media-2017-09-30/TimestampRange
type TimestampRange struct {
	_ struct{} `type:"structure"`

	// The ending time stamp in the range of time stamps for which to return fragments.
	//
	// EndTimestamp is a required field
	EndTimestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The starting time stamp in the range of time stamps for which to return fragments.
	//
	// StartTimestamp is a required field
	StartTimestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s TimestampRange) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TimestampRange) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TimestampRange) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TimestampRange"}

	if s.EndTimestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTimestamp"))
	}

	if s.StartTimestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTimestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TimestampRange) MarshalFields(e protocol.FieldEncoder) error {
	if s.EndTimestamp != nil {
		v := *s.EndTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EndTimestamp", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.StartTimestamp != nil {
		v := *s.StartTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartTimestamp", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	return nil
}

type FragmentSelectorType string

// Enum values for FragmentSelectorType
const (
	FragmentSelectorTypeProducerTimestamp FragmentSelectorType = "PRODUCER_TIMESTAMP"
	FragmentSelectorTypeServerTimestamp   FragmentSelectorType = "SERVER_TIMESTAMP"
)

func (enum FragmentSelectorType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FragmentSelectorType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
