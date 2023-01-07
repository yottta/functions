# AWS EC2 Shutdown

This function is shutting down all instances in the region it is deployed.

## Configuration
This function does not require any special configuration as the only env variable that it is using is one that is provided by the AWS Lambda environment.
* AWS_REGION - the region where the function is deployed and also the region from which the function is shutting down the EC2 instances.

## Deployment
At the moment it only can be deployed as a zip file containing the binary.

Just run `make zip` and it generates a zip archive with the binary built inside, called `app`.