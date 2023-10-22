// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a configuration for DNS query logging. If you delete a configuration,
// Amazon Route 53 stops sending query logs to CloudWatch Logs. Route 53 doesn't
// delete any logs that are already in CloudWatch Logs. For more information about
// DNS query logs, see CreateQueryLoggingConfig (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html)
// .
func (c *Client) DeleteQueryLoggingConfig(ctx context.Context, params *DeleteQueryLoggingConfigInput, optFns ...func(*Options)) (*DeleteQueryLoggingConfigOutput, error) {
	if params == nil {
		params = &DeleteQueryLoggingConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteQueryLoggingConfig", params, optFns, c.addOperationDeleteQueryLoggingConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteQueryLoggingConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQueryLoggingConfigInput struct {

	// The ID of the configuration that you want to delete.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DeleteQueryLoggingConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteQueryLoggingConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteQueryLoggingConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteQueryLoggingConfig{}, middleware.After)
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
	if err = addOpDeleteQueryLoggingConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueryLoggingConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteQueryLoggingConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteQueryLoggingConfig",
	}
}
