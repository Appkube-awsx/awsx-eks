package clustercmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-eks/authenticator"
	"github.com/Appkube-awsx/awsx-eks/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			clusterName, _ := cmd.Flags().GetString("clusterName")
			if clusterName != "" {
				getClusterDetails(region, crossAccountRoleArn, acKey, secKey, clusterName, externalId)
			} else {
				log.Fatalln("clusterName not provided. Program exit")
			}
		}
	},
}

func getClusterDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, clusterName string, externalId string) *eks.DescribeClusterOutput {
	log.Println("Getting aws cluster data")
	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
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

func init() {
	GetConfigDataCmd.Flags().StringP("clusterName", "t", "", "Cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("clusterName"); err != nil {
		fmt.Println(err)
	}
}
