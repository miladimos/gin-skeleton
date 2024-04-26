package services

import "gin-skeleton/internal/repository"

type PostService interface {
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostRepository(postRepository repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}
