package hub

import "context"

type IHubService interface {
	GetHub(ctx context.Context, req GetHubRequest) (GetHubResponse, error)
}
