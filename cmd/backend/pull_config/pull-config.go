
package pull_config

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
			Region: aws.String("us-west-2")},
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
        return fmt.Errorf("Error creating directory: %v", err)
    }

    content := strings.ToUpper(paramName) + "=" + paramValue

    // Open or create the ".env" file within the directory
    f, err := os.OpenFile(path+"/.env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("Error creating .env file: %v", err)
    }
    defer f.Close()

    fmt.Println("New code is Writing to .env file")
    fmt.Println(path)

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
	backend.BackendCmd.AddCommand(pullConfigCmd)
}



// old code 

/*

func writeToEnvFile(paramName string, paramValue string, path string) error {
    envFile, err := os.OpenFile(path+"/hellosuryadlpcli.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    fmt.Println("Writing to old code .env file")
    fmt.Println(path)

    if err != nil {
        return fmt.Errorf("Error opening .env file: %v", err)
    }

    defer envFile.Close()

    _, err = envFile.WriteString(strings.ToUpper(paramName) + "=" + paramValue + "\n")

    if err != nil {
        return fmt.Errorf("Error writing to .env file: %v", err)
    }

    return nil
}
*/





/*
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/frontend"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/utils"
	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm" // For Parameter Store
    "github.com/aws/aws-sdk-go/service/secretsmanager" // For Secrets Manager

)
// Define the directories where I want to write the .env files
var directories = []string{"/training", "/frontend"}

// Define the AWS region and Parameter Store path
const (
    awsRegion = "us-east-1"
    ssmPath   = "/cdk-bootstrap/hnb659fds/version"
)

func main() {
    // Create an AWS session
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(awsRegion),
    }))

    // Create an SSM client
    ssmClient := ssm.New(sess)

    // Retrieve the values from Parameter Store
    for _, directory := range directories {
        envFilePath := directory + "/.env"
        envExamplePath := directory + "/.env.example"

        // Retrieve the values from Parameter Store
		params := &ssm.GetParametersInput{
    Names: []*string{
        aws.String("/cdk-bootstrap/hnb659fds/version"), // Actual parameter path
        // Add more parameter names as needed
    },
    WithDecryption: aws.Bool(true),


        resp, err := ssmClient.GetParameters(params)
        if err != nil {
            fmt.Printf("Error retrieving parameters: %v\n", err)
            os.Exit(1)
        }

        // Create or update .env files
        envContent := ""
        envExampleContent := ""

        for _, param := range resp.Parameters {
            envContent += fmt.Sprintf("%s=%s\n", *param.Name, *param.Value)
            envExampleContent += fmt.Sprintf("%s=\n", *param.Name)
        }

        // Write .env and .env.example files
        if err := ioutil.WriteFile(envFilePath, []byte(envContent), 0644); err != nil {
            fmt.Printf("Error writing %s: %v\n", envFilePath, err)
            os.Exit(1)
        }

        if err := ioutil.WriteFile(envExamplePath, []byte(envExampleContent), 0644); err != nil {
            fmt.Printf("Error writing %s: %v\n", envExamplePath, err)
            os.Exit(1)
        }

        fmt.Printf("Values written to %s and %s\n", envFilePath, envExamplePath)
        }
    }
}



func main() {
    // Create a session with AWS credentials.
    sess := session.Must(session.NewSession())

    // Get the secret from Secrets Manager.
    svc := secretsmanager.New(sess)
    params := &secretsmanager.GetSecretValueInput{
         SecretId: aws.String("MY_SECRET_ID"), //will configure actual secret id  later
    }
    resp, err := svc.GetSecretValue(params)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Write the secret to the .env file in the current directory.
    secret := resp.SecretString
    os.Setenv("SECRET_KEY", strings.TrimSpace(secret))  //will configure actual secret key later
}

*/

