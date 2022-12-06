package dal

import (
	"backend/database"
	"backend/models"
)

func CreatePost(post *models.Post) error {
	
	if result := database.DB.Create(post); result.Error != nil {
		return result.Error
	}

	return nil
}

func ReadPost(postId string) (*models.Post, error) {
	post:= new(models.Post)

	if result := database.DB.First(&post, postId); result.Error != nil {
		return  nil, result.Error
	}

	return post, nil 
}
