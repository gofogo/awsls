// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
)

func ListMskCluster(client *Client) ([]Resource, error) {
	req := client.kafkaconn.ListClustersRequest(&kafka.ListClustersInput{})

	var result []Resource

	p := kafka.NewListClustersPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.ClusterInfoList {

			tags := map[string]string{}
			for k, v := range r.Tags {
				tags[k] = v
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_msk_cluster",
				ID:        *r.ClusterArn,
				Region:    client.kafkaconn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
