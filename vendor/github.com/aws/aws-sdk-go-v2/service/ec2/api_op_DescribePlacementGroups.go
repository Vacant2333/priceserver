// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified placement groups or all of your placement groups. For
// more information, see [Placement groups]in the Amazon EC2 User Guide.
//
// [Placement groups]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
func (c *Client) DescribePlacementGroups(ctx context.Context, params *DescribePlacementGroupsInput, optFns ...func(*Options)) (*DescribePlacementGroupsOutput, error) {
	if params == nil {
		params = &DescribePlacementGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePlacementGroups", params, optFns, c.addOperationDescribePlacementGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePlacementGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePlacementGroupsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//
	//   - group-name - The name of the placement group.
	//
	//   - group-arn - The Amazon Resource Name (ARN) of the placement group.
	//
	//   - spread-level - The spread level for the placement group ( host | rack ).
	//
	//   - state - The state of the placement group ( pending | available | deleting |
	//   deleted ).
	//
	//   - strategy - The strategy of the placement group ( cluster | spread |
	//   partition ).
	//
	//   - tag: - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources that have a tag with a specific key, regardless of the tag value.
	Filters []types.Filter

	// The IDs of the placement groups.
	GroupIds []string

	// The names of the placement groups.
	//
	// Default: Describes all your placement groups, or only those otherwise specified.
	GroupNames []string

	noSmithyDocumentSerde
}

type DescribePlacementGroupsOutput struct {

	// Information about the placement groups.
	PlacementGroups []types.PlacementGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePlacementGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribePlacementGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribePlacementGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePlacementGroups"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePlacementGroups(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribePlacementGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePlacementGroups",
	}
}