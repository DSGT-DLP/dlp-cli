package build_training_env

/*
$ your_cli_tool build-training-env --secret "YourTrainingSecretName" --bucket "YourTrainingBucketName"
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

var secretNameTraining string // Name of the secret in AWS Secrets Manager
var bucketNameTraining string // Name of the bucket for training

var buildTrainingEnvCmd = &cobra.Command{
    Use:   "build-training-env",
    Short: "Build .env file for training/",
    Run: func(cmd *cobra.Command, args []string) {
        sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
        if err != nil {
            log.Fatal("Error creating AWS session: ", err)
        }

        smClient := secretsmanager.New(sess)
        secretValue, err := smClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(secretNameTraining)})
        if err != nil {
            log.Fatal("Error retrieving secret: ", err)
        }

        path := "./training"

        // Adding secrets to the .env file
        writeToEnvFile(secretNameTraining, *secretValue.SecretString, path)

        // Adding bucket name to the .env file
        writeToEnvFile("BUCKET_NAME", bucketNameTraining, path)
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
    buildTrainingEnvCmd.Flags().StringVar(&secretNameTraining, "secret", "", "Name of the secret in AWS Secrets Manager for training")
    buildTrainingEnvCmd.Flags().StringVar(&bucketNameTraining, "bucket", "", "Name of the training bucket")
    backend.BackendCmd.AddCommand(buildTrainingEnvCmd)
}
