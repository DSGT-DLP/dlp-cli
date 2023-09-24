package pull_config

/*
$ go run main.go backend pull-config
*/

import (
	"log"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/utils" // For utils/
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm" // For Parameter Store
	"github.com/spf13/cobra"
)

var pullConfigCmd = &cobra.Command{
	Use:   "pull-config",
	Short: "Pulls values from parameter store/secrets manager and writes them to .env files",
	Long: `This command pulls values from parameter store/secrets manager and writes them to .env files in /training and /frontend.`,
    Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(backend.AwsRegion)},
		)

		if err != nil {
			log.Fatal("Error creating AWS session: ", err)
		}

		ssmClient := ssm.New(sess)

		parameterNames := []string{"parameter-name"}

		for _, paramName := range parameterNames {
			result, err := ssmClient.GetParameter(&ssm.GetParameterInput{
				Name:           aws.String(paramName),
				WithDecryption: aws.Bool(true),
			})

			if err != nil {
				log.Fatal("Error getting parameter: ", err)
			}

			utils.WriteToEnvFile(paramName, *result.Parameter.Value, backend.BackendDir)
			utils.WriteToEnvFile(paramName, *result.Parameter.Value, "./frontend")
		}
	},
}

func init() {
	backend.BackendCmd.AddCommand(pullConfigCmd)
}