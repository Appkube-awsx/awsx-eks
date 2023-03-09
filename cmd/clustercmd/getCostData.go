package clustercmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-eks/authenticator"
	"github.com/Appkube-awsx/awsx-eks/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetCostDataCmd = &cobra.Command{
	Use:   "getCostData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		env := cmd.Parent().PersistentFlags().Lookup("env").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()
		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, env, externalId)
		

		if authFlag {
			getClusterCostDetail(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}

func getClusterCostDetail(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*costexplorer.GetCostAndUsageOutput, error) {
	log.Println("Getting cost data")
	costClient := client.GetCostClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2023-01-01"),
			End:   aws.String("2023-02-01"),
		},
		Metrics: []*string{
			aws.String("UnblendedCost"),
		},
		Granularity: aws.String("MONTHLY"),
		Filter: &costexplorer.Expression{
			Dimensions: &costexplorer.DimensionValues{
				Key: aws.String("SERVICE"),
				Values: []*string{
					aws.String("Clusters"),
				},
			},
		},
	}

	costData, err := costClient.GetCostAndUsage(input)
	if err != nil {
		log.Fatalln("Error: in getting cost data", err)
	}
	log.Println(costData)
	return costData, err
}

func init() {
	GetCostDataCmd.Flags().StringP("clusterName", "t", "", "Cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("clusterName"); err != nil {
		fmt.Println(err)
	}
}