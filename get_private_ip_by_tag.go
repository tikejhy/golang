package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    sess := session.Must(session.NewSession())

    productPtr := flag.String("product", "meotic", "Product, Default: meotic")
    rolePtr := flag.String("role", "web", "Role e.g: web, api, mysql; Default: web")
    envPtr := flag.String("env", "deploy", "Environment e.g: prod, qa, or deploy; Default: deploy")
    countPtr := flag.Int("count", 1, "a Integer")
    awsPtr := flag.String("region", "eu-west-1", "AWS Region, Default: eu-west-1" )

    flag.Parse()

    awsRegion := *awsPtr
    svc := ec2.New(sess, &aws.Config{Region: aws.String(awsRegion)})
    params := &ec2.DescribeInstancesInput{
        Filters: []*ec2.Filter{
            {
                Name: aws.String("tag:Name"),
                Values: []*string{
                    aws.String(strings.Join([]string{*productPtr, *envPtr, *rolePtr}, ".")),
                },
            },
        },
    }
    resp, err := svc.DescribeInstances(params)
    if err != nil {
        fmt.Println("there was an error listing instances in", awsRegion, err.Error())
        log.Fatal(err.Error())
    }

    numberofInstance := 0
    for idx := range resp.Reservations {

        for range resp.Reservations[idx].Instances {
            numberofInstance = idx

        }
    }
    i := 1
    sum := *countPtr - i
    if sum < 0 {
        println("Count doesnot satisfy number of instances")
        os.Exit(9)
    }
    if sum <= numberofInstance {
        println(*resp.Reservations[sum].Instances[0].PrivateIpAddress)
    }
    if sum > numberofInstance {
        println("Count doesnot satisfy number of instances")
        os.Exit(9)
    }
}
