package client

import (
	"context"

	"github.com/qovery/qovery-client-go"

	"terraform-provider-qovery/client/apierrors"
)

func (c *Client) getEnvironmentStatus(ctx context.Context, environmentID string) (*qovery.Status, *apierrors.APIError) {
	status, res, err := c.api.EnvironmentMainCallsApi.
		GetEnvironmentStatus(ctx, environmentID).
		Execute()
	if err != nil || res.StatusCode >= 400 {
		return nil, apierrors.NewReadError(apierrors.APIResourceEnvironmentStatus, environmentID, res, err)
	}
	return status, nil
}
