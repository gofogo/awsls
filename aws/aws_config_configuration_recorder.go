// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
)

func ListConfigConfigurationRecorder(client *Client) ([]Resource, error) {
	req := client.configserviceconn.DescribeConfigurationRecordersRequest(&configservice.DescribeConfigurationRecordersInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ConfigurationRecorders) > 0 {
		for _, r := range resp.ConfigurationRecorders {

			result = append(result, Resource{
				Type:   "aws_config_configuration_recorder",
				ID:     *r.Name,
				Region: client.configserviceconn.Config.Region,
			})
		}
	}

	return result, nil
}
