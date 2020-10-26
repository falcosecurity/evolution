resource "aws_s3_bucket" "distribution_bucket" {
  bucket = var.bucket_name
  acl    = "public-read"

  tags = {
    Name = var.bucket_name
  }
}

resource "aws_s3_bucket" "logging_bucket" {
  bucket = var.logging_bucket_name
  acl    = "private"

  tags = {
    Name = var.logging_bucket_name
  }
}

resource "aws_acm_certificate" "cert" {
  domain_name       = var.distribution_name_alias
  validation_method = "DNS"
  provider          = aws.us

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_cloudfront_distribution" "distribution" {
  origin {
    domain_name = aws_s3_bucket.distribution_bucket.bucket_regional_domain_name
    origin_id   = var.distribution_origin_id
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Falco distribution channel"
  default_root_object = "index.html"

  logging_config {
    include_cookies = false
    bucket          = "${var.logging_bucket_name}.s3.amazonaws.com"
    prefix          = "falco-distribution"
  }

  aliases = [var.distribution_name_alias]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = var.distribution_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  price_class = "PriceClass_All"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn = aws_acm_certificate.cert.arn
    ssl_support_method  = "sni-only"
  }
}
