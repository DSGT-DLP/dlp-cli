package build_frontend_env

/*
$ dlp-cli build-frontend-env-file --secret "YourSecretName" --bucket "YourBucketName"
*/

import (
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
    "github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/frontend" // For frontend/
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/utils" // For utils/
)

var secretName string // Name of the secret in AWS Secrets Manager.
var bucketName string // Name of the bucket (Assuming one bucket here). 

var buildFrontendEnvCmd = &cobra.Command{
    Use:   "build-frontend-env-file",
    Short: "Build .env file for frontend/",
    Run: func(cmd *cobra.Command, args []string) {
        sess, err := session.NewSession(&aws.Config{Region: aws.String(frontend.AwsRegion)})
        if err != nil {
            log.Fatal("error creating AWS session: ", err)
        }

        smClient := secretsmanager.New(sess)
        secretValue, err := smClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(secretName)})
        if err != nil {
            log.Fatal("error retrieving secret: ", err)
        }

        path := "./frontend"

        // Adding secrets to the .env file
        utils.WriteToEnvFile(secretName, *secretValue.SecretString, path)

        // ** need to add bucket name to .env file **
        utils.WriteToEnvFile("BUCKET_NAME", bucketName, path)
    },
}

func init() {
    buildFrontendEnvCmd.Flags().StringVar(&secretName, "secret", "", "Name of the secret in AWS Secrets Manager")
    buildFrontendEnvCmd.Flags().StringVar(&bucketName, "bucket", "", "Name of the bucket")
    frontend.FrontendCmd.AddCommand(buildFrontendEnvCmd)
}
