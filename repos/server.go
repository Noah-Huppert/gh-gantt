package repos

import (
	"context"
)

// DefaultRepositoriesServer implements the Repositories service
type DefaultRepositoriesServer struct{}

// NewDefaultRepositoriesServer creates a new DefaultRepositoriesServer instance
func NewDefaultRepositoriesServer() *DefaultRepositoriesServer {
	return &DefaultRepositoriesServer{}
}

func (s *DefaultRepositoriesServer) List(ctx context.Context, req *ListReq) (
	*ListResp, error) {

	return &ListResp{
		Repositories: []*Repository{},
	}, nil
}
