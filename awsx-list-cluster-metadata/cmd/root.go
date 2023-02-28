package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-list-clusters-metadata/client"
	"github.com/Appkube-awsx/awsx-list-clusters-metadata/vault"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/spf13/cobra"
)

var listClusterMetadataCmd = &cobra.Command{
	Use:   "getListClusterMetaDataDetails",
	Short: "getListClusterMetaDataDetails command gets resource counts",
	Long:  `getListClusterMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getElementDetails started")
		vaultUrl, _ := cmd.Flags().GetString("vaultUrl")
		accountNo, _ := cmd.Flags().GetString("accountId")
		region, _ := cmd.Flags().GetString("zone")
		acKey, _ := cmd.Flags().GetString("accessKey")
		secKey, _ := cmd.Flags().GetString("secretKey")
		env, _ := cmd.Flags().GetString("env")

		if vaultUrl != "" && accountNo != "" && env != "" {
			if region == "" {
				log.Fatalln("Zone not provided. Program exit")
				return
			}
			log.Println("Getting account details")
			data, err := vault.GetAccountDetails(vaultUrl, accountNo)
			if err != nil {
				log.Println("Error in calling the account details api. \n", err)
				return
			}
			if data.AccessKey == "" || data.SecretKey == "" {
				log.Println("Account details not found.")
				return
			}
			getListClusterMetadata(region, data.AccessKey, data.SecretKey, env)
		} else if region != "" && acKey != "" && secKey != "" && env != "" {
			getListClusterMetadata(region, acKey, secKey, env)
		} else {
			log.Fatal("region", secKey)
			log.Fatal("AWS credentials like accesskey/secretkey/region/crossAccountRoleArn not provided. Program exit")
			return
		}

	},
}

type Tags struct {
	Environment string `json:"Environment"`
}

type Cluster struct {
	Tags Tags
}

type Response struct {
	Cluster Cluster
}

// json.Unmarshal
func getListClusterMetadata(region string, accessKey string, secretKey string, env string) *eks.ListClustersOutput {
	listClusterClient := client.GetClient(region, accessKey, secretKey)
	listClusterRequest := &eks.ListClustersInput{}
	listClusterResponse, err := listClusterClient.ListClusters(listClusterRequest)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	for _, clusterName := range listClusterResponse.Clusters {
		clusterDetails := getClusterDetails(region, accessKey, secretKey, *clusterName)

		// clusterDetails.Cluster.Tags.Environment
		var responseObject Response
		jsonedResponse, _ := json.Marshal(clusterDetails)
		json.Unmarshal([]byte(string(jsonedResponse)), &responseObject)
		if env == responseObject.Cluster.Tags.Environment {
			log.Println(responseObject.Cluster.Tags)
		} else {
			log.Println("No cluster present for Env:", env)

		}
	}
	return listClusterResponse
}

func getClusterDetails(region string, accessKey string, secretKey string, clusterName string) *eks.DescribeClusterOutput {
	log.Println("Getting aws List Cluster Count summary")
	listClusterClient := client.GetClient(region, accessKey, secretKey)
	input := &eks.DescribeClusterInput{
		Name: aws.String(clusterName),
	}
	clusterDetailsResponse, err := listClusterClient.DescribeCluster(input)
	// log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}

	return clusterDetailsResponse
}

func Execute() {
	err := listClusterMetadataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	listClusterMetadataCmd.Flags().String("vaultUrl", "", "vault end point")
	listClusterMetadataCmd.Flags().String("accountId", "", "aws account number")
	listClusterMetadataCmd.Flags().String("zone", "", "aws region")
	listClusterMetadataCmd.Flags().String("accessKey", "", "aws access key")
	listClusterMetadataCmd.Flags().String("secretKey", "", "aws secret key")
	listClusterMetadataCmd.Flags().String("env", "", "aws env Required")

}
