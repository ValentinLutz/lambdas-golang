locals {
  database_secret = {
    username = aws_db_instance.database.username
    password = aws_db_instance.database.password
    endpoint = aws_db_instance.database.endpoint
    name     = aws_db_instance.database.db_name
  }
}