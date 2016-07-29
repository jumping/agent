package g

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
)

func RegionInstanceId() (region string, instanceID string) {
	c := ec2metadata.New(session.New())
        ec2InstanceIdentifyDocument, _ := c.GetInstanceIdentityDocument()
	region     = ec2InstanceIdentifyDocument.Region
	instanceID = ec2InstanceIdentifyDocument.InstanceID
	//fmt.Println(instanceID)
	return
}

func GetNameTag(name string) (value string, err error) {

	region, instanceID := RegionInstanceId()

	svc := ec2.New(session.New(&aws.Config{Region: aws.String(region)}))

	params := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	if len(resp.Reservations) == 0 {
		return "", err
	}
	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			for _, tag := range inst.Tags {
				if (name != "") && (*tag.Key == name) {
					//fmt.Println(*tag.Key, "=", *tag.Value)
					value := strings.Replace(*tag.Value, " ", "", -1)
					return value, nil
				}
			}
		}
	}
	return
}

