resource "aws_sfn_state_machine" "state_machine" {
  name     = module.database_cost_saver_name.name
  role_arn = aws_iam_role.state_machine_iam_role.arn

  definition = jsonencode({
    "Comment" : "Simple step function to shutdown RDS instances",
    "StartAt" : "DescribeDBInstances",
    "States" : {
      "DescribeDBInstances" : {
        "Next" : "Map",
        "Parameters" : {},
        "Resource" : "arn:aws:states:::aws-sdk:rds:describeDBInstances",
        "Type" : "Task"
      },
      "Map" : {
        "End" : true,
        "ItemProcessor" : {
          "ProcessorConfig" : {
            "Mode" : "INLINE"
          },
          "StartAt" : "StopDBInstance",
          "States" : {
            "StopDBInstance" : {
              "End" : true,
              "Parameters" : {
                "DbInstanceIdentifier.$" : "$.DbInstanceIdentifier"
              },
              "Resource" : "arn:aws:states:::aws-sdk:rds:stopDBInstance",
              "Type" : "Task"
            }
          }
        },
        "Type" : "Map",
        "InputPath" : "$.DbInstances.[?(@.DbInstanceStatus == 'available')]"
      }
    }
  })
}

resource "aws_scheduler_schedule" "scheduler" {
  name                = module.database_cost_saver_name.name
  description         = "Rule to trigger the database cost saver state machine"
  schedule_expression = "cron(0 4 * * ? *)"

  flexible_time_window {
    mode                      = "FLEXIBLE"
    maximum_window_in_minutes = 60
  }

  target {
    arn      = aws_sfn_state_machine.state_machine.arn
    role_arn = aws_iam_role.scheduler_iam_role.arn
  }
}