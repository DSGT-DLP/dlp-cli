package build_training_env

/*
$ dlp-cli build-training-env-file --secret "YourTrainingSecretName" --bucket "YourTrainingBucketName"
*/

import (
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
    "github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/utils" // For utils/
)

var secretNameTraining string // Name of the secret in AWS Secrets Manager
var bucketNameTraining string // Name of the bucket for training

var buildTrainingEnvCmd = &cobra.Command{
    Use:   "build-training-env-file",
    Short: "Build .env file for training/",
    Run: func(cmd *cobra.Command, args []string) {
        sess, err := session.NewSession(&aws.Config{Region: aws.String(backend.AwsRegion)})
        if err != nil {
            log.Fatal("error session: ", err)
        }

        smClient := secretsmanager.New(sess)
        secretValue, err := smClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(secretNameTraining)})
        if err != nil {
            log.Fatal("error retrieving secret: ", err)
        }

        path := "./training"

        // Adding secrets to the .env file
        utils.WriteToEnvFile(secretNameTraining, *secretValue.SecretString, path)

        // Adding bucket name to the .env file
        utils.WriteToEnvFile("BUCKET_NAME", bucketNameTraining, path)
    },
}

func init() {
    buildTrainingEnvCmd.Flags().StringVar(&secretNameTraining, "secret", "", "Name of the secret in AWS Secrets Manager for training")
    buildTrainingEnvCmd.Flags().StringVar(&bucketNameTraining, "bucket", "", "Name of the training bucket")
    backend.BackendCmd.AddCommand(buildTrainingEnvCmd)
}
