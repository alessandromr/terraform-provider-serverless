# Terraform Provider Serverless
This is a Work in Progress provider to speed up serverless workflows. 
This terraform provider is created to work alongside the main official providers. 

## Why
Creating serverless workflow in terraform can be a little bit tedious, an example can be Api Gateway resources for Lambda.
This provider is inspirated by Serverless Framework and abstract all the tedious logic, leaving the developer only the simple and fast part of the job.
Using this provider **will not limit your terraform code**, you can still use official provider along side and interact with serverless resources in the standard manner.

## Public Clouds
Right now this terraform provider supports only AWS. With Terraform, sdk and client logic is separated from the actual provider, so in the future it will be pretty easy to extend support to the main cloud providers.  
You can find the client logic for AWS here [go-aws-serverless](https://github.com/alessandromr/go-aws-serverless). This is actually a pretty simple wrap of AWS sdk.  
Feel free to contribute.

## Examples

### Example AWS (WiP Syntax Can Change) with Api Gateway
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



### Example AWS (WiP Syntax Can Change) with S3
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

