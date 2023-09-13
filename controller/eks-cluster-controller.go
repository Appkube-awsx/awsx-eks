package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-eks/command"
	"github.com/aws/aws-sdk-go/service/eks"
	"log"
)

func GetEksClusterByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) ([]*eks.DescribeClusterOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetEksClustersByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetEksClusterByUserCreds(region string, accesskey string, secretKey string, crossAccountRoleArn string, externalId string) ([]*eks.DescribeClusterOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accesskey, secretKey, crossAccountRoleArn, externalId)
	return GetEksClustersByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetEksClustersByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) ([]*eks.DescribeClusterOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.GetClusterList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetEksClusters(clientAuth *client.Auth) ([]*eks.DescribeClusterOutput, error) {
	response, err := command.GetClusterList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
