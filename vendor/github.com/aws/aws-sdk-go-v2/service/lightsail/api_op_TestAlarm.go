// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests an alarm by displaying a banner on the Amazon Lightsail console. If a
// notification trigger is configured for the specified alarm, the test also sends
// a notification to the notification protocol ( Email and/or SMS ) configured for
// the alarm. An alarm is used to monitor a single metric for one of your
// resources. When a metric condition is met, the alarm can notify you by email,
// SMS text message, and a banner displayed on the Amazon Lightsail console. For
// more information, see Alarms in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-alarms)
// .
func (c *Client) TestAlarm(ctx context.Context, params *TestAlarmInput, optFns ...func(*Options)) (*TestAlarmOutput, error) {
	if params == nil {
		params = &TestAlarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestAlarm", params, optFns, c.addOperationTestAlarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestAlarmInput struct {

	// The name of the alarm to test.
	//
	// This member is required.
	AlarmName *string

	// The alarm state to test. An alarm has the following possible states that can be
	// tested:
	//   - ALARM - The metric is outside of the defined threshold.
	//   - INSUFFICIENT_DATA - The alarm has just started, the metric is not available,
	//   or not enough data is available for the metric to determine the alarm state.
	//   - OK - The metric is within the defined threshold.
	//
	// This member is required.
	State types.AlarmState

	noSmithyDocumentSerde
}

type TestAlarmOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestAlarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTestAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpTestAlarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestAlarm(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opTestAlarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "TestAlarm",
	}
}
