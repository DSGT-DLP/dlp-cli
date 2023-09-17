package update_param

/*
$ go run main.go backend update-param --name "ExistingParameterName" --value "NewParameterValue" --type "String" --overwrite true
*/

import (
	"fmt"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/spf13/cobra"
)

var (
	paramName  string
	paramValue string
	paramType  string
	overwrite  bool
)

var (
	awsRegion = "us-west-2"
)

var updateParamCmd = &cobra.Command{
	Use:   "update-param",
	Short: "Update an existing parameter value in AWS Parameter Store",
	Run: func(cmd *cobra.Command, args []string) {
		if paramName == "" || paramValue == "" || paramType == "" {
			fmt.Println("Error: You must provide a name, value, and type for the parameter.")
			return
		}

		// Create an AWS session with the specified region
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		}))

		svc := ssm.New(sess)
		input := &ssm.PutParameterInput{
			Name:      aws.String(paramName),
			Value:     aws.String(paramValue),
			Type:      aws.String(paramType),
			Overwrite: aws.Bool(overwrite),
		}

		_, err := svc.PutParameter(input)
		if err != nil {
			fmt.Println("Error updating parameter:", err)
			return
		}

		fmt.Println("Parameter updated successfully.")
	},
}

func init() {
	updateParamCmd.Flags().StringVar(&paramName, "name", "", "Name of the parameter to update")
	updateParamCmd.Flags().StringVar(&paramValue, "value", "", "New value for the parameter")
	updateParamCmd.Flags().StringVar(&paramType, "type", "", "Type of the parameter (String, StringList, SecureString)")
	updateParamCmd.Flags().BoolVar(&overwrite, "overwrite", true, "Whether to overwrite the parameter if it exists")
	backend.BackendCmd.AddCommand(updateParamCmd)
}
