resource "aws_s3_bucket" "distribution_bucket" {
  bucket = var.bucket_name
  acl    = "public-read"

  tags = {
    Name        = var.bucket_name
  }
}

