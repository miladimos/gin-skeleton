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

// func (u UserServiceImpl) GetAllPost(c *gin.Context) {
// 	// defer pkg.PanicHandler(c)

// 	// logger("start to execute get all data user")

// 	data, err := u.userRepository.FindAllUser()
// 	if err != nil {
// 		log.Error("Happened Error when find all user data. Error: ", err)
// 		pkg.PanicException(constant.UnknownError)
// 	}

// 	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
// }
