# go-aws-lambda

## Summary

Includes Go based Lambda functions for demonstration purposes

## Deploy from Local

```bash
# Build Lambda function binary
GOOS=linux go build -o bootstrap main.go

# Zip lambda function
zip aws-helloworld-lamba.zip bootstrap
  adding: bootstrap (deflated 46%)
  
# Manually upload in the UI or use AWS CLI to publish

# Invoke API
curl -H 'Content-Type: application/json' https://7k<redacted>mc.execute-api.eu-central-1.amazonaws.com/lambda-go-hello -d '{"first_name": "Gareth", "last_name": "B"}'
```