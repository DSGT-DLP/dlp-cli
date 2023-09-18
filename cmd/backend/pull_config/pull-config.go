
package pull_config

/*
$ go run main.go backend pull-config
*/

import (
	"os"
	"strings"
    "log"
    "fmt"
    
	"github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm" // For Parameter Store
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

			writeToEnvFile(paramName, *result.Parameter.Value, backend.BackendDir)
			writeToEnvFile(paramName, *result.Parameter.Value, "./frontend")
		}
	},
}



func writeToEnvFile(paramName string, paramValue string, path string) error {
    
    // Create the directory if it doesn't exist
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        return fmt.Errorf("error creating directory: %v", err)
    }

    content := strings.ToUpper(paramName) + "=" + paramValue

    // Open or create the ".env" file within the directory
    f, err := os.OpenFile(path+"/.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("error creating .env file: %v", err)
    }
    defer f.Close()

    fmt.Println("New code is Writing to .env file")
    fmt.Println(path)

    n, err := f.WriteString(content + "\n")
    if err != nil {
        return fmt.Errorf("error writing to .env file: %v", err)
    }
    fmt.Printf("Wrote %d bytes to .env file\n", n)

    err = f.Sync()
    if err != nil {
        return fmt.Errorf("error syncing .env file: %v", err)
    }

    return nil
}

func init() {
	backend.BackendCmd.AddCommand(pullConfigCmd)
}