package tfutil

func NewStateFilePath(region string, env string, resource string) string {
	return region + "-" + env + "/" + resource + "/terraform.tfstate"
}
