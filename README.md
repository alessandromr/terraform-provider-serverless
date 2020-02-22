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
resource "serverless_aws_function_http" "testhttpfunction" {
  filename = "main.zip"
  memory_size = 256
  function_name = "TestFunctionHTTP"
  handler = "main"
  runtime = "go1.x"
  role = "arn:aws:iam::12344556768:role/LambdaTestRole"
  event{
    path = "test"
    http_method = "ANY"
    api_name="TestAPI"
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
resource "serverless_aws_function_s3" "tests3" {
  filename = "main.zip"
  memory_size = 256
  function_name = "S3TestFunction"
  handler = "main"
  runtime = "go1.x"
  role = "arn:aws:iam::12345678910:role/LambdaTestRole"
  event{
    bucket = aws_s3_bucket.test_bucket.id
    event_types = ["s3:ObjectCreated:*"]
  }
}
```

