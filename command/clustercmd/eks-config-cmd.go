package clustercmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "config data for eks cluster",
	Long:  `config data for eks cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			clusterName, _ := cmd.Flags().GetString("clusterName")
			if clusterName != "" {
				getClusterDetails(clusterName, *clientAuth)
			} else {
				log.Fatalln("cluster name not provided. program exit")
			}
		}
	},
}

func getClusterDetails(clusterName string, auth client.Auth) *eks.DescribeClusterOutput {
	log.Println("Getting aws eks cluster data")
	listClusterClient := client.GetClient(auth, client.EKS_CLIENT).(*eks.EKS)
	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}
	clusterDetailsResponse, err := listClusterClient.DescribeCluster(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func GetCluster(eksClient *eks.EKS, clusterName string) *eks.DescribeClusterOutput {
	log.Println("Getting aws cluster detail for cluster: ", clusterName)
	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}
	clusterDetailsResponse, err := eksClient.DescribeCluster(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("clusterName", "c", "", "cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("clusterName"); err != nil {
		fmt.Println(err)
	}
}
