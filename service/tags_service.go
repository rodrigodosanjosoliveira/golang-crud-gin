package service

import (
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/data/request"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
