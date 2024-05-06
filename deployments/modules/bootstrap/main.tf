resource "aws_s3_bucket" "state" {
  bucket = module.s3_bucket_name.name

  lifecycle {
    prevent_destroy = false
  }
}

resource "aws_s3_bucket_public_access_block" "state" {
  bucket = aws_s3_bucket.state.bucket

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_versioning" "state" {
  bucket = aws_s3_bucket.state.bucket

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "state" {
  bucket = aws_s3_bucket.state.bucket

  rule {
    bucket_key_enabled = true

    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

resource "aws_s3_bucket_lifecycle_configuration" "state" {
  bucket = aws_s3_bucket.state.bucket

  rule {
    id     = "${module.s3_bucket_name.name}-1"
    status = "Enabled"

    noncurrent_version_expiration {
      newer_noncurrent_versions = 10
      noncurrent_days           = 30
    }
  }
}

resource "aws_dynamodb_table" "state_lock" {
  name         = module.dynamodb_name.name
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}