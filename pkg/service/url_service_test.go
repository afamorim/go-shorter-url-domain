package service

import (
	"testing"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UrlRepositoryMock struct {
	mock.Mock
}

func (r *UrlRepositoryMock) Save(url model.Url) (model.Url, error) {
	args := r.Called(url)

	return args.Get(0).(model.Url), args.Error(1)
}

func (r *UrlRepositoryMock) FindByShorter(shorterUrl string) (model.Url, error) {
	url := model.Url{
		Id:          1,
		OriginalUrl: "http://www.teste.com.br",
		CompressUrl: "http://yougotit.com",
	}

	return url, nil
}

func TestSaveSucess(t *testing.T) {
	url := model.Url{
		OriginalUrl: "http://www.teste.com",
	}

	urlRepositoryMock := UrlRepositoryMock{}
	urlRepositoryMock.On("Save", url).Return(
		model.Url{
			Id: 2,
		},
		nil)

	urlService := NewUrlService(&urlRepositoryMock)

	newUrl, _ := urlService.Save(url)

	assert.Equal(t, 2, newUrl.Id)

	urlRepositoryMock.AssertExpectations(t)
}

func TestSaveEmptyUrlError(t *testing.T) {
	url := model.Url{}

	urlRepositoryMock := UrlRepositoryMock{}
	urlRepositoryMock.On("Save", url).Return(
		model.Url{
			Id: 2,
		},
		nil).Times(4)

	urlService := NewUrlService(&urlRepositoryMock)

	_, err := urlService.Save(url)
	assert.True(t, err != nil, err.Error())

	urlRepositoryMock.AssertExpectations(t)
}
