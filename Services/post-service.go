package Services

import (
	"gg-cms/DTOs"
	"gg-cms/DataRepos"
	"gg-cms/Models"
)

type PostService interface {
	Save(Models.Post) (Models.Post, error)
	Update(Models.Post) (Models.Post, error)
	Delete(ID string) error
	Find(permaLink string) (Models.Post, error)
	FindAll(limit int64, skip int64, areActive bool) (DTOs.Posts, error)
}

type postService struct {
	postRepo DataRepos.PostRepo
}

func NewPostService(repo DataRepos.PostRepo) PostService {
	return &postService{
		postRepo: repo,
	}
}

func (service *postService) Save(post Models.Post) (Models.Post, error) {
	newPost, err := service.postRepo.Insert(post)
	return newPost, err
}

func (service *postService) Update(post Models.Post) (Models.Post, error) {
	newPost, err := service.postRepo.Update(post)
	return newPost, err
}

func (service *postService) Delete(id string) error {
	err := service.postRepo.Delete(id)
	return err
}

func (service *postService) Find(permaLink string) (Models.Post, error) {
	newPost, err := service.postRepo.Get(permaLink)
	return newPost, err
}

func (service *postService) FindAll(limit int64, skip int64, areActive bool) (DTOs.Posts, error) {
	posts, err := service.postRepo.GetAllActive(limit, skip, areActive)
	return posts, err
}