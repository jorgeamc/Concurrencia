provider "aws" {
  region = "us-east-1"  # Cambia la región según tus necesidades
}

resource "aws_dynamodb_table" "example" {
  name           = "users_dynamo_db_test"
  billing_mode   = "PAY_PER_REQUEST"  # La capa gratuita usa Pay-Per-Request
  hash_key       = "user-id"
  attribute {
    name = "user-id"
    type = "S"
  }
}