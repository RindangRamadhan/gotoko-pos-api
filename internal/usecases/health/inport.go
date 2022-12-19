package health

import (
	"context"
)

type Inport interface {
	Execute(context.Context) InportResponse
}

type InportResponse struct {
	Name    string `json:"name"`
	MySQL   bool   `json:"mysql"`
	Version string `json:"version"`
}
