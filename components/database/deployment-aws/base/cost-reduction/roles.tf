data "aws_iam_policy_document" "state_machine_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["states.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "state_machine_iam_role" {
  name               = "${module.database_cost_saver_name.name}-state-machine"
  assume_role_policy = data.aws_iam_policy_document.state_machine_assume_role.json
}

data "aws_iam_policy_document" "state_machine_document" {
  statement {
    actions = [
      "rds:DescribeDBInstances",
      "rds:StopDBInstance",
    ]

    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "state_machine_policy" {
  role   = aws_iam_role.state_machine_iam_role.id
  policy = data.aws_iam_policy_document.state_machine_document.json
}


data "aws_iam_policy_document" "scheduler_assume_role" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["scheduler.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "scheduler_iam_role" {
  name               = "${module.database_cost_saver_name.name}-scheduler"
  assume_role_policy = data.aws_iam_policy_document.scheduler_assume_role.json
}

data "aws_iam_policy_document" "scheduler_document" {
  statement {
    actions = [
      "states:StartExecution",
    ]

    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "scheduler_policy" {
  role   = aws_iam_role.scheduler_iam_role.id
  policy = data.aws_iam_policy_document.scheduler_document.json
}