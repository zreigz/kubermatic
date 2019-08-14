package main

import (
	"fmt"
	"os"

	"github.com/kubermatic/kubermatic/api/pkg/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func main() {
	app := cli.NewApp()
	app.Name = "subnet tester"
	app.Usage = ""
	app.Version = "v1.0.0"
	app.Description = "Helper tool to test subnet-describe"

	endpointFlag := cli.StringFlag{
		Name:  "endpoint, e",
		Value: "",
		Usage: "S3 endpoint",
	}
	accessKeyIDFlag := cli.StringFlag{
		Name:   "access-key-id",
		Value:  "",
		EnvVar: "ACCESS_KEY_ID",
		Usage:  "S3 AccessKeyID",
	}
	secretAccessKeyFlag := cli.StringFlag{
		Name:   "secret-access-key",
		Value:  "",
		EnvVar: "SECRET_ACCESS_KEY",
		Usage:  "S3 SecretAccessKey",
	}
	logDebugFlag := cli.BoolFlag{
		Name:  "log-debug",
		Usage: "Enables more verbose logging",
	}

	defaultLogFormat := log.FormatJSON
	logFormatFlag := cli.GenericFlag{
		Name:  "log-format",
		Value: &defaultLogFormat,
		Usage: fmt.Sprintf("Use one of [%v] to change the log output", log.AvailableFormats),
	}

	app.Flags = []cli.Flag{
		logDebugFlag,
		logFormatFlag,
	}

	app.Commands = []cli.Command{
		{
			Name:   "describe",
			Usage:  "describes a subnet",
			Action: describeSubnets,
			Flags: []cli.Flag{
				endpointFlag,
				accessKeyIDFlag,
				secretAccessKeyFlag,
			},
		},
	}

	// setup logging
	app.Before = func(c *cli.Context) error {
		format := c.GlobalGeneric("log-format").(*log.Format)
		rawLog := log.New(c.GlobalBool("log-debug"), *format)
		logger = rawLog.Sugar()

		return nil
	}

	defer func() {
		if logger != nil {
			if err := logger.Sync(); err != nil {
				fmt.Println(err)
			}
		}
	}()

	err := app.Run(os.Args)
	// Only log failures when the logger has been setup, otherwise
	// we know it's been a CLI parsing failure and the cli package
	// has already output the error and printed the usage hints.
	if err != nil && logger != nil {
		logger.Fatalw("Failed to run command", zap.Error(err))
	}
}

func describeSubnets(c *cli.Context) error {
	svc := ec2.New(session.New())
	input := &ec2.DescribeSubnetsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("vpc-id"),
				Values: []*string{
					aws.String("vpc-05f7c3d610c646741"),
				},
			},
		},
	}

	result, err := svc.DescribeSubnets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return fmt.Errorf("%s", err.Error())
		}
	}

	fmt.Println(result)
	return nil
}
