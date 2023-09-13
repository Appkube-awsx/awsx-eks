package command

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-eks/command/clustercmd"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var AwsxEksClusterCmd = &cobra.Command{
	Use:   "getClusterList",
	Short: "getClusterList command gets cluster list",
	Long:  `getClusterList command gets cluster list of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetClusterList(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

// json.Unmarshal
func GetClusterList(auth client.Auth) ([]*eks.DescribeClusterOutput, error) {
	log.Println("getting cluster list")

	eksClient := client.GetClient(auth, client.EKS_CLIENT).(*eks.EKS)
	request := &eks.ListClustersInput{}
	response, err := eksClient.ListClusters(request)
	if err != nil {
		log.Fatalln("Error in getting cluster list", err)
	}
	allClusters := []*eks.DescribeClusterOutput{}

	for _, clusterName := range response.Clusters {
		clusterDetail := clustercmd.GetCluster(eksClient, *clusterName)
		allClusters = append(allClusters, clusterDetail)
	}

	log.Println(allClusters)
	return allClusters, err
}

func Execute() {
	err := AwsxEksClusterCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxEksClusterCmd.AddCommand(clustercmd.GetConfigDataCmd)

	AwsxEksClusterCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxEksClusterCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxEksClusterCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxEksClusterCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxEksClusterCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxEksClusterCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxEksClusterCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxEksClusterCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
