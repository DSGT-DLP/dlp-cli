package build_serverless_env

/*
$ your_cli_tool build-serverless-env --sst "YourSSTVariables" --dev-endpoints "YourDevEndpoints" --bucket "YourServerlessBucketName"
*/

import (
    "os"
    "strings"
    "fmt"
    "github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/serverless" // For frontend/
)

var sstVariables string // Variables related to sst
var devEndpoints string // Development endpoints
var bucketNameServerless string // Name of the bucket for serverless

var buildServerlessEnvCmd = &cobra.Command{
    Use:   "build-serverless-env",
    Short: "Build .env file for serverless/",
    Run: func(cmd *cobra.Command, args []string) {
        path := "./serverless"

        // Adding sst variables to the .env file
        writeToEnvFile("SST_VARIABLES", sstVariables, path)

        // Adding dev endpoints to the .env file
        writeToEnvFile("DEV_ENDPOINTS", devEndpoints, path)

        // Adding bucket name to the .env file
        writeToEnvFile("BUCKET_NAME", bucketNameServerless, path)
    },
}

func writeToEnvFile(paramName string, paramValue string, path string) error {
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return fmt.Errorf("Error creating directory: %v", err)
    }

    content := strings.ToUpper(paramName) + "=" + paramValue
    f, err := os.OpenFile(path+"/.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("Error creating .env file: %v", err)
    }
    defer f.Close()

    fmt.Println("Writing to .env file", path)
    n, err := f.WriteString(content + "\n")
    if err != nil {
        return fmt.Errorf("Error writing to .env file: %v", err)
    }
    fmt.Printf("Wrote %d bytes to .env file\n", n)

    err = f.Sync()
    if err != nil {
        return fmt.Errorf("Error syncing .env file: %v", err)
    }

    return nil
}

func init() {
    buildServerlessEnvCmd.Flags().StringVar(&sstVariables, "sst", "", "Variables related to serverless stack (sst)")
    buildServerlessEnvCmd.Flags().StringVar(&devEndpoints, "dev-endpoints", "", "Development endpoints for serverless")
    buildServerlessEnvCmd.Flags().StringVar(&bucketNameServerless, "bucket", "", "Name of the serverless bucket")
    serverless.ServerlessCmd.AddCommand(buildServerlessEnvCmd) 
}