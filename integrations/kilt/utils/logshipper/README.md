Simple log shipper to cloudwatch

Forwards stdin to CW Log group configured with `__CW_LOG_GROUP`

Requires credentials that can create log groups. See [here](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html)
for the default credentials chain in the aws go sdk.
