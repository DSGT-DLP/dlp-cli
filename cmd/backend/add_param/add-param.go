package add_param

/*
$ go run main.go backend add-param --name "YourParameterName" --value "YourParameterValue" --type "String"
*/

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ssm"
    "github.com/spf13/cobra"
    "github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
)

var (
    paramName  string
    paramValue string
    paramType  string
)

var (
    awsRegion = "us-west-2"
)

var addParamCmd = &cobra.Command{
    Use:   "add-param",
    Short: "Add a new param to AWS Parameter Store",
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
            Name:  aws.String(paramName),
            Value: aws.String(paramValue),
            Type:  aws.String(paramType),
        }

        _, err := svc.PutParameter(input)
        if err != nil {
            fmt.Println("Error adding parameter:", err)
            return
        }

        fmt.Println("Parameter added successfully.")
    },
}

func init() {
    addParamCmd.Flags().StringVar(&paramName, "name", "", "Name of the parameter")
    addParamCmd.Flags().StringVar(&paramValue, "value", "", "Value of the parameter")
    addParamCmd.Flags().StringVar(&paramType, "type", "", "Type of the parameter (String, StringList, SecureString)")
    backend.BackendCmd.AddCommand(addParamCmd)
}
