package posts

import (
	"context"
	"simple-forum/internal/models/posts"
)

func (s *Service) GetALlPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := pageSize * (pageIndex - 1)

	resp, err := s.postRepo.GetAllPost(ctx, limit, offset)
	if err != nil {
		return posts.GetAllPostResponse{}, err
	}

	return resp, nil
}
