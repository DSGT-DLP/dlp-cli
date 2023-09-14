package add_param

/*
$ dlp-cli add-param --name "YourParameterName" --value "YourParameterValue" --type "String"
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

var addParamCmd = &cobra.Command{
    Use:   "add-param",
    Short: "Add a new param to AWS Parameter Store",
    Run: func(cmd *cobra.Command, args []string) {
        if paramName == "" || paramValue == "" || paramType == "" {
            fmt.Println("Error: You must provide a name, value, and type for the parameter.")
            return
        }

        svc := ssm.New(session.New())
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
