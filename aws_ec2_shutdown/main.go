package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const defaultRegion = "us-east-1"

func main() {
	lambda.Start(HandleLambdaEvent)
}

func HandleLambdaEvent(event interface{}) error {
	fmt.Println("event", event)

	region := envOrDefault("AWS_REGION", defaultRegion)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	client := ec2.New(ec2.Options{
		Credentials: cfg.Credentials,
		Region:      region,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	instances, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
	if err != nil {
		return fmt.Errorf("describe: %w", err)
	}
	var ids []string
	for ri := range instances.Reservations {
		for i := range instances.Reservations[ri].Instances {
			instance := instances.Reservations[ri].Instances[i]
			if state := instance.State; state != nil && state.Name == types.InstanceStateNameRunning {
				ids = append(ids, *instance.InstanceId)
			}
		}
	}
	if len(ids) == 0 {
		log.Println("no instances to stop")
		return nil
	}

	if _, err := client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: ids,
	}); err != nil {
		return fmt.Errorf("stopping: %w", err)
	}
	log.Printf("instances stopped: %s", ids)
	return nil
}

func envOrDefault(key, defVal string) string {
	val := os.Getenv(key)
	val = strings.TrimSpace(val)
	if len(val) == 0 {
		return defVal
	}
	return val
}
