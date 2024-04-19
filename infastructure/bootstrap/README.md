# Bootstrap

This module is responsible for creating the necessary resources to save CDKTF state in AWS.

## Usage

1. Create a stack in the desired region and account.
    ```go
    bootstrap.NewStack(app, "eu-central-1", "test")
    ```
2. Create a config for the deployment stage.
    ```go
    var stageConfigs = map[string]*StageConfig{
        "eu-central-1-test": {
            Region:      "eu-central-1",
            Environment: "test",
            Profile:     "admin",
            // Bucket:      jsii.String("terraform-state-5449924400404832213"),
        },
    }
    ```
3. Run the following command to create the resources.
    ```shell
    REGION=eu-central-1 ENVIRONMENT=test RESOURCE=bootstrap mage cdk:deploy;
    ```
4. Add the previous created bucket to the stage config.
    ```go
    var stageConfigs = map[string]*StageConfig{
        "eu-central-1-test": {
            Region:      "eu-central-1",
            Environment: "test",
            Profile:     "admin",
            Bucket:      jsii.String("terraform-state-5449924400404832213"),
        },
    }
    ```
5. Run terrafrom init to migrate the local state to the remote state.
   ```shell
   terraform -chdir=cdktf.out/stacks/boostrap-eu-central-1-test init -backend-config=profile=admin -migrate-state
   ```