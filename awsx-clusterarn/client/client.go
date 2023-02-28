package client

import (
	"github.com/Appkube-awsx/awsx-list-clusters/awssession"
	"github.com/aws/aws-sdk-go/service/eks"
)

func GetClient(region string, accessKey string, secretKey string) *eks.EKS {
	awsSession := awssession.GetSessionByCreds(region, accessKey, secretKey)
	svc := eks.New(awsSession)
	return svc
}
