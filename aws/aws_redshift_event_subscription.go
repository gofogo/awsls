// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

func ListRedshiftEventSubscription(client *Client) ([]Resource, error) {
	req := client.redshiftconn.DescribeEventSubscriptionsRequest(&redshift.DescribeEventSubscriptionsInput{})

	var result []Resource

	p := redshift.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_redshift_event_subscription",
				ID:     *r.CustSubscriptionId,
				Region: client.redshiftconn.Config.Region,
				Tags:   tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
