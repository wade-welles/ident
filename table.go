package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"math"
	"os"
	"strconv"
)

type Key map[string]interface{}

type Table struct {
	Name string
	svc  *dynamodb.DynamoDB
}

func (t Table) Increment(c *Counter) error {
	// We lean on UpdateItem to increment the counter in an atomic way. DynamoDB
	// takes care of that for us.
	out, err := t.svc.UpdateItem(&dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"cluster": &dynamodb.AttributeValue{
				S: aws.String(c.Cluster),
			},
		},
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			"seq": &dynamodb.AttributeValueUpdate{
				Action: aws.String("ADD"),
				Value: &dynamodb.AttributeValue{
					N: aws.String("1"),
				},
			},
		},
		TableName:    aws.String(t.Name),
		ReturnValues: aws.String("UPDATED_NEW"),
	})

	if err != nil {
		return err
	}

	// Otherwise, let's get our value and put it in to the counter!
	attr, ok := out.Attributes["seq"]

	if !ok {
		return ErrMissingSeq
	}

	i, err := strconv.ParseInt(aws.StringValue(attr.N), 10, 64)

	if err != nil {
		return err
	}

	if i > math.MaxUint16 {
		return ErrRanOutOfIDs
	}

	c.Seq = uint16(i)

	return nil
}

func NewTable(name string) *Table {
	// NOTE: This should be taken care of by the enviro, but we're explicitly
	// doing the override because why the the Lambda role fucks with this thing.
	credentials := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"",
	)

	config := aws.NewConfig().
		WithRegion(os.Getenv("AWS_REGION")).
		WithCredentials(credentials)

	sess := session.New(config)
	svc := dynamodb.New(sess)
	return &Table{name, svc}
}
