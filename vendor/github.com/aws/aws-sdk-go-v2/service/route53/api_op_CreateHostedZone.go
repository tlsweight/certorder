// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new public or private hosted zone. You create records in a public
// hosted zone to define how you want to route traffic on the internet for a
// domain, such as example.com, and its subdomains (apex.example.com,
// acme.example.com). You create records in a private hosted zone to define how you
// want to route traffic for a domain and its subdomains within one or more Amazon
// Virtual Private Clouds (Amazon VPCs). You can't convert a public hosted zone to
// a private hosted zone or vice versa. Instead, you must create a new hosted zone
// with the same name and create new resource record sets. For more information
// about charges for hosted zones, see Amazon Route 53 Pricing (http://aws.amazon.com/route53/pricing/)
// . Note the following:
//   - You can't create a hosted zone for a top-level domain (TLD) such as .com.
//   - For public hosted zones, Route 53 automatically creates a default SOA
//     record and four NS records for the zone. For more information about SOA and NS
//     records, see NS and SOA Records that Route 53 Creates for a Hosted Zone (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/SOA-NSrecords.html)
//     in the Amazon Route 53 Developer Guide. If you want to use the same name servers
//     for multiple public hosted zones, you can optionally associate a reusable
//     delegation set with the hosted zone. See the DelegationSetId element.
//   - If your domain is registered with a registrar other than Route 53, you must
//     update the name servers with your registrar to make Route 53 the DNS service for
//     the domain. For more information, see Migrating DNS Service for an Existing
//     Domain to Amazon Route 53 (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html)
//     in the Amazon Route 53 Developer Guide.
//
// When you submit a CreateHostedZone request, the initial status of the hosted
// zone is PENDING . For public hosted zones, this means that the NS and SOA
// records are not yet available on all Route 53 DNS servers. When the NS and SOA
// records are available, the status of the zone changes to INSYNC . The
// CreateHostedZone request requires the caller to have an ec2:DescribeVpcs
// permission. When creating private hosted zones, the Amazon VPC must belong to
// the same partition where the hosted zone is created. A partition is a group of
// Amazon Web Services Regions. Each Amazon Web Services account is scoped to one
// partition. The following are the supported partitions:
//   - aws - Amazon Web Services Regions
//   - aws-cn - China Regions
//   - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see Access Management (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// in the Amazon Web Services General Reference.
func (c *Client) CreateHostedZone(ctx context.Context, params *CreateHostedZoneInput, optFns ...func(*Options)) (*CreateHostedZoneOutput, error) {
	if params == nil {
		params = &CreateHostedZoneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHostedZone", params, optFns, c.addOperationCreateHostedZoneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHostedZoneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to create a public
// or private hosted zone.
type CreateHostedZoneInput struct {

	// A unique string that identifies the request and that allows failed
	// CreateHostedZone requests to be retried without the risk of executing the
	// operation twice. You must use a unique CallerReference string every time you
	// submit a CreateHostedZone request. CallerReference can be any unique string,
	// for example, a date/time stamp.
	//
	// This member is required.
	CallerReference *string

	// The name of the domain. Specify a fully qualified domain name, for example,
	// www.example.com. The trailing dot is optional; Amazon Route 53 assumes that the
	// domain name is fully qualified. This means that Route 53 treats www.example.com
	// (without a trailing dot) and www.example.com. (with a trailing dot) as
	// identical. If you're creating a public hosted zone, this is the name you have
	// registered with your DNS registrar. If your domain name is registered with a
	// registrar other than Route 53, change the name servers for your domain to the
	// set of NameServers that CreateHostedZone returns in DelegationSet .
	//
	// This member is required.
	Name *string

	// If you want to associate a reusable delegation set with this hosted zone, the
	// ID that Amazon Route 53 assigned to the reusable delegation set when you created
	// it. For more information about reusable delegation sets, see
	// CreateReusableDelegationSet (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html)
	// .
	DelegationSetId *string

	// (Optional) A complex type that contains the following optional values:
	//   - For public and private hosted zones, an optional comment
	//   - For private hosted zones, an optional PrivateZone element
	// If you don't specify a comment or the PrivateZone element, omit HostedZoneConfig
	// and the other elements.
	HostedZoneConfig *types.HostedZoneConfig

	// (Private hosted zones only) A complex type that contains information about the
	// Amazon VPC that you're associating with this hosted zone. You can specify only
	// one Amazon VPC when you create a private hosted zone. If you are associating a
	// VPC with a hosted zone with this request, the paramaters VPCId and VPCRegion
	// are also required. To associate additional Amazon VPCs with the hosted zone, use
	// AssociateVPCWithHostedZone (https://docs.aws.amazon.com/Route53/latest/APIReference/API_AssociateVPCWithHostedZone.html)
	// after you create a hosted zone.
	VPC *types.VPC

	noSmithyDocumentSerde
}

// A complex type containing the response information for the hosted zone.
type CreateHostedZoneOutput struct {

	// A complex type that contains information about the CreateHostedZone request.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// A complex type that describes the name servers for this hosted zone.
	//
	// This member is required.
	DelegationSet *types.DelegationSet

	// A complex type that contains general information about the hosted zone.
	//
	// This member is required.
	HostedZone *types.HostedZone

	// The unique URL representing the new hosted zone.
	//
	// This member is required.
	Location *string

	// A complex type that contains information about an Amazon VPC that you
	// associated with this hosted zone.
	VPC *types.VPC

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHostedZoneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateHostedZone{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateHostedZone{}, middleware.After)
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
	if err = addOpCreateHostedZoneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHostedZone(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateHostedZone(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateHostedZone",
	}
}