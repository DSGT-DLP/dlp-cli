package build_frontend_env

/*
Sample Command:
$ dlp-cli build-frontend-env-file --secret "YourSecretName"
*/

import (
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
    "github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/frontend" // For frontend/
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/utils" // For utils/
    "encoding/json" // to unmarshal json
)

var secretName string // Name of the secret in AWS Secrets Manager

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
        
        // AWS secrets are parsed from key:value pairs in secrets manager for populating the .env file

        var secretsMap map[string]string // declares a variable of type Map with key:value as string:string

        if err := json.Unmarshal([]byte(*secretValue.SecretString), &secretsMap); err != nil { // stores unmarshalled json into secretsMap
            log.Fatal("error unmarshalling json: ", err)
        } 
        
        for key,value := range secretsMap {
            utils.WriteToEnvFile(key, value, path)
        }

        // Hardcoding bucket name as a constant
        utils.WriteToEnvFile("BUCKET_NAME", utils.DlpUploadBucket, path)
    },
}

func init() {
    buildFrontendEnvCmd.Flags().StringVar(&secretName, "secret", "", "Name of the secret in AWS Secrets Manager")
    frontend.FrontendCmd.AddCommand(buildFrontendEnvCmd)
}
