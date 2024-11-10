# go-aws-lambda

## Summary

Includes Go based Lambda functions for demonstration purposes

## Deploy from Local

```bash
# Build Lambda function binary
./build-deploy.sh
Building functions...
Found Lambda Function File at cmd/functions/helloperson/main.go...
Compiling Lambda function to binary...
Found Lambda Function File at cmd/functions/hellopersonapigateway/main.go...
Compiling Lambda function to binary...
Finished!

# Manually upload in the UI or use AWS CLI to publish

# Invoke API
curl -H 'Content-Type: application/json' https://7k<redacted>mc.execute-api.eu-central-1.amazonaws.com/lambda-go-hello -d '{"first_name": "Gareth", "last_name": "B"}'
...
```