package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type AWSAnalyzerData struct {
	Vpcs    []types.Vpc
	Subnets []types.Subnet
}

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	fmt.Println("Configuration")
	fmt.Println("Region: " + cfg.Region)

	data := AWSAnalyzerData{}
	data.Vpcs = getVPCs(cfg)
	data.Subnets = getSubnets(cfg)
}

func getVPCs(cfg aws.Config) []types.Vpc {
	client := ec2.NewFromConfig(cfg)
	inputVPC := &ec2.DescribeVpcsInput{}
	resultVPC, err := client.DescribeVpcs(context.TODO(), inputVPC)

	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon VPCs:")
		panic(err)
	}

	fmt.Println(len(resultVPC.Vpcs), " VPCs.")

	for _, vpc := range resultVPC.Vpcs {
		fmt.Println("VPC Id:" + *vpc.VpcId)
		fmt.Println("VPC CIDR block:" + *vpc.CidrBlock)
		fmt.Println("VPC associated CIDR blocks:")
		for _, cidr := range vpc.CidrBlockAssociationSet {
			fmt.Println("  - " + *cidr.CidrBlock)
		}
		fmt.Println("Tags:")
		for _, tag := range vpc.Tags {
			fmt.Println("  - ", *tag.Key, ": ", *tag.Value)
		}

		fmt.Println("URL: https://" + cfg.Region + ".console.aws.amazon.com/vpc/home?region=" + cfg.Region + "#VpcDetails:VpcId=" + *vpc.VpcId)
	}

	return resultVPC.Vpcs

}

func getSubnets(cfg aws.Config) []types.Subnet {
	client := ec2.NewFromConfig(cfg)
	subnetReq := &ec2.DescribeSubnetsInput{}
	resultSubnets, err := client.DescribeSubnets(context.TODO(), subnetReq)

	if err != nil {
		fmt.Println("Got an error retrieving information about your Subnets:")
		panic(err)
	}

	fmt.Println(len(resultSubnets.Subnets), " subnets.")

	for _, subnet := range resultSubnets.Subnets {
		fmt.Println("Subnet Id: " + *subnet.SubnetId)
		fmt.Println("Subnet ARN: " + *subnet.SubnetArn)
		fmt.Println("Subnet AZId: " + *subnet.AvailabilityZoneId)
		fmt.Println("Subnet AZ: " + *subnet.AvailabilityZone)
		fmt.Println("Subnet CIDR block: " + *subnet.CidrBlock)
		fmt.Println("Subnet VPCId: " + *subnet.VpcId)
		fmt.Println("Tags:")
		for _, tag := range subnet.Tags {
			fmt.Println("  - ", *tag.Key, ": ", *tag.Value)
		}
		fmt.Println("URL: https://" + cfg.Region + ".console.aws.amazon.com/vpc/home?region=" + cfg.Region + "#SubnetDetails:subnetId=" + *subnet.SubnetId)
	}

	return resultSubnets.Subnets

}
