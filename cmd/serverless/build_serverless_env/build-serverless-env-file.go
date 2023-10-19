package build_serverless_env

/*
Sample Command:
$ dlp-cli build-serverless-env-file --sst "YourSSTVariables" --dev-endpoints "YourDevEndpoints"
*/

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/serverless" // For serverless/
	"github.com/spf13/cobra"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/utils" // For utils/
)

var sstVariables string // Variables related to sst
var devEndpoints string // Development endpoints

var buildServerlessEnvCmd = &cobra.Command{
	Use:   "build-serverless-env-file",
	Short: "Build .env file for serverless/",
	Run: func(cmd *cobra.Command, args []string) {
		path := "./serverless"

		// Adding sst variables to the .env file
		utils.WriteToEnvFile("SST_VARIABLES", sstVariables, path)

		// Adding dev endpoints to the .env file
		utils.WriteToEnvFile("DEV_ENDPOINTS", devEndpoints, path)

		// Hardcoding bucket name as a constant
		utils.WriteToEnvFile("BUCKET_NAME", utils.DlpUploadBucket, path)
	},
}

func init() {
	buildServerlessEnvCmd.Flags().StringVar(&sstVariables, "sst", "", "Variables related to serverless stack (sst)")
	buildServerlessEnvCmd.Flags().StringVar(&devEndpoints, "dev-endpoints", "", "Development endpoints for serverless")
	serverless.ServerlessCmd.AddCommand(buildServerlessEnvCmd)
}
