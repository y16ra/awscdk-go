package stacks

import (
	"y16ra/awscdk-go/resources"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type PointCoreStackProps struct {
	StackProps awscdk.StackProps
}

func NewPointCoreStack(scope constructs.Construct, id string, props *PointCoreStackProps) awscdk.Stack {

	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	vpc, _ := resources.Network(stack, id, nil)
	resources.Ecs(stack, id, nil, vpc)

	return stack
}
