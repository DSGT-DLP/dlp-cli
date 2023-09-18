package remove_param

/*
$ dlp-cli backend remove-param --name "YourParameterName"
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
    paramName string
)

var removeParamCmd = &cobra.Command{
    Use:   "remove-param",
    Short: "Remove a parameter from AWS Parameter Store",
    Run: func(cmd *cobra.Command, args []string) {
        if paramName == "" {
            fmt.Println("Error: You must provide the name of the parameter to remove.")
            return
        }

        // Create an AWS session with the specified region
        sess := session.Must(session.NewSession(&aws.Config{
            Region: aws.String(backend.AwsRegion),
        }))

        svc := ssm.New(sess)
        input := &ssm.DeleteParameterInput{
            Name: aws.String(paramName),
        }

        _, err := svc.DeleteParameter(input)
        if err != nil {
            fmt.Println("Error removing parameter:", err)
            return
        }

        fmt.Println("Parameter removed successfully.")
    },
}

func init() {
    removeParamCmd.Flags().StringVar(&paramName, "name", "", "Name of the parameter to remove")
    backend.BackendCmd.AddCommand(removeParamCmd)
}
