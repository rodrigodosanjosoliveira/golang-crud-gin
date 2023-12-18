package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/data/request"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/data/response"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/helper"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/model"
	"github.com/rodrigodosanjosoliveira/golang-crud-gin/repository"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagsRepository,
		Validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	return response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		t := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, t)
	}

	return tags
}
