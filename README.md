# EKS CLi's

## To list all the EKS cluster, run the following command:

```bash
awsx-eks --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId>
```

## To retrieve the configuration details of a specific EKS clustercmd, run the following command:

```bash
awsx-eks getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId>  --clusterName <clusterName>
```
