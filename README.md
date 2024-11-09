# go-aws-lambda

## Summary

Includes Go based Lambda functions for demonstration purposes

## Deploy from Local

```bash
# Build Lambda function binary
GOOS=linux go build -o bootstrap main.go

# Zip lambda function
zip aws-helloworld-lamba.zip bootstrap                                                                                                                                              ✔  1803  21:37:47
  adding: bootstrap (deflated 46%)
  
# Manually upload in the UI or use AWS CLI to publish

```