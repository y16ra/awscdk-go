package resources

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	ecsp "github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func Ecs(scope constructs.Construct, id string, props *cdk.StackProps, vpc ec2.Vpc) ecsp.ApplicationLoadBalancedFargateService {
	cluster := awsecs.NewCluster(scope, jsii.String("Cluster"), &awsecs.ClusterProps{
		Vpc: vpc,
	})
	ecs := ecsp.NewApplicationLoadBalancedFargateService(scope, jsii.String("EcsService"), &ecsp.ApplicationLoadBalancedFargateServiceProps{
		Cluster:        cluster,
		Cpu:            jsii.Number(256),
		MemoryLimitMiB: jsii.Number(512),
		DesiredCount:   jsii.Number(1),
		TaskImageOptions: &ecsp.ApplicationLoadBalancedTaskImageOptions{
			Image:         awsecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample"), nil),
			ContainerPort: jsii.Number(80),
		},
	})
	return ecs
}
