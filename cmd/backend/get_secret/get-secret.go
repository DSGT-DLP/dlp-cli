package get_secret

/*
dlp-cli get-secret --name "YourSecretName"
*/


import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var secretName string

var getSecretCmd = &cobra.Command{
	Use:   "get-secret",
	Short: "Retrieve a secret from AWS Secrets Manager",
	Run: func(cmd *cobra.Command, args []string) {
		if secretName == "" {
			fmt.Println("Error: You must provide the name of the secret.")
			return
		}

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)
		if err != nil {
			fmt.Println("Error creating AWS session:", err)
			return
		}

		smClient := secretsmanager.New(sess)

		input := &secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretName),
		}

		result, err := smClient.GetSecretValue(input)
		if err != nil {
			fmt.Println("Error retrieving secret:", err)
			return
		}

		fmt.Printf("Secret [%s]: %s\n", secretName, *result.SecretString)
	},
}

func init() {
	getSecretCmd.Flags().StringVar(&secretName, "name", "", "Name of the secret to retrieve")
	backend.BackendCmd.AddCommand(getSecretCmd)
}
