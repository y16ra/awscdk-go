package resources

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func Network(scope constructs.Construct, id string, props *cdk.StackProps) (
	ec2.Vpc,
	ec2.Subnet,
) {
	vpc := ec2.NewVpc(scope, jsii.String("GoCdkVpc"), &ec2.VpcProps{
		Cidr:                   jsii.String("10.0.0.0/16"),
		DefaultInstanceTenancy: ec2.DefaultInstanceTenancy_DEFAULT,
		EnableDnsSupport:       jsii.Bool(true),
		EnableDnsHostnames:     jsii.Bool(true),
		NatGateways:            jsii.Number(1),
		MaxAzs:                 jsii.Number(2),
	})

	// subnet := ec2.NewSubnet(scope, jsii.String(""), &ec2.SubnetProps{
	// 	CidrBlock: jsii.String(""),
	// 	VpcId:     vpc.VpcId(),
	// })
	return vpc, nil
}
