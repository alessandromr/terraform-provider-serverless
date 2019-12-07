# Terraform Provider AWS Serverless Provider

This is a Work in Progress provider to speed up serverless workflows. 
This provider is compatible with official terraform aws provider and they can easily work togheter.

## Why
Creating serverless workflow in terraform can be a little bit tedious, specially if we are talking about Api Gateway + Lambda.
This provider is inspirated by Serverless Framework and abstract all the tedious logic, leaving the developer only the simple and fast part of the job.
Using this provider **will not limit your terraform code**, you can still use official provider along side and interact with serverless resources in the standard manner.


## Examples

### Example (WiP Not Stable) with Api Gateway
In this example is shown how simple is creating a function triggered by Api Gateway.
If we were not using AWS Serverless Provider we would have to create all Api Gateway logics (Rest Api, Resources, Methods, Integrations, ...).


```hcl
#created with AWS Serverless
resource "aws_serverless_function" "test_function" {
    filename      = "lambda_function_payload.zip"
    function_name = "lambda_function_name"
    handler       = "exports.test"
    source_code_hash = "${filebase64sha256("lambda_function_payload.zip")}"
    runtime = "nodejs8.10"

    environment {
        variables = {
            foo = "bar"
        }
    }

    event = {
        http{
            path = "/test"
            method = "GET"
            api_exist = false
        }
    }

}
```



### Example (WiP Not Stable) with S3
```hcl

#created with official provider
resource "aws_s3_bucket" "test_bucket" {
    bucket = "test_bucket"
    acl    = "private"
}

#created with AWS Serverless
resource "aws_serverless_function" "test_function" {
    filename      = "lambda_function_payload.zip"
    function_name = "lambda_function_name"
    handler       = "exports.test"
    source_code_hash = "${filebase64sha256("lambda_function_payload.zip")}"
    runtime = "nodejs8.10"

    environment {
        variables = {
            foo = "bar"
        }
    }

    event = {
        s3{
            bucket = aws_s3_bucket.test_bucket.id
            suffix = ".json"
            event_type = ["s3:ObjectCreated"]
        }
    }

}

```

