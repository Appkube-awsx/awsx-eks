# EKS CLi's

## To list all the EKS cluster, run the following command:

```bash
awsx-eks --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId> --env <env>
```

## To retrieve the configuration details of a specific EKS clustercmd, run the following command:

```bash
awsx-eks getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --clusterName <clusterName>
```

## To retrieve the cost details of a specific EKS clustercmd, run the following command:

```bash
awsx-eks getCostData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env>
``` 
## To retrieve the cost Spikes details of a specific EKS clustercmd, run the following command:

```bash
awsx-eks GetCostSpike -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --granularity <granularity> --startDate <startDate> --endDate <endDate> --serviceName <"serviceName">
```
