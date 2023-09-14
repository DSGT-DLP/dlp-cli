package build_frontend_env

/*
$ your_cli_tool build-frontend-env --secret "YourSecretName" --bucket "YourBucketName"
*/

import (
    "os"
    "strings"
    "fmt"
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
    "github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
)

var secretName string // Name of the secret in AWS Secrets Manager. 
var bucketName string // Name of the bucket (Assuming one bucket here) 

var buildFrontendEnvCmd = &cobra.Command{
    Use:   "build-frontend-env",
    Short: "Build .env file for frontend/",
    Run: func(cmd *cobra.Command, args []string) {
        sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
        if err != nil {
            log.Fatal("Error creating AWS session: ", err)
        }

        smClient := secretsmanager.New(sess)
        secretValue, err := smClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(secretName)})
        if err != nil {
            log.Fatal("Error retrieving secret: ", err)
        }

        path := "./frontend"

        // Adding secrets to the .env file
        writeToEnvFile(secretName, *secretValue.SecretString, path)

        // ** need to add project constants to .env file **
        writeToEnvFile("PROJECT_CONSTANT", "YourConstantValue", path)

        // ** need to add bucket name to .env file **
        writeToEnvFile("BUCKET_NAME", bucketName, path)
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
    buildFrontendEnvCmd.Flags().StringVar(&secretName, "secret", "", "Name of the secret in AWS Secrets Manager")
    buildFrontendEnvCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
    backend.BackendCmd.AddCommand(buildFrontendEnvCmd)
}
