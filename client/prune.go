package client

import (
	"context"
	"fmt"

	controlapi "github.com/moby/buildkit/api/services/control"
	"github.com/pkg/errors"
)

//Prune the system
func (c *Client) Prune(ctx context.Context) (map[string]int64, error) {
	fmt.Println(">> Invoking Prune() from Client..")
	_, err := c.controlClient().Prune(ctx, &controlapi.PruneRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to Prune system")
	}

	return nil, nil
}
