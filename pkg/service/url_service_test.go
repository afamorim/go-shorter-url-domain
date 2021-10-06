package service

import (
	"fmt"
	"testing"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UrlRepositoryMock struct {
	mock.Mock
}

func (r *UrlRepositoryMock) Save(url model.Url) error {
	args := r.Called(url)

	return args.Error(0)
}

func (r *UrlRepositoryMock) FindByShorter(shorterUrl string) (interface{}, error) {
	url := model.Url{
		Id:          1,
		OriginalUrl: "http://www.teste.com.br",
		CompressUrl: "http://yougotit.com",
	}

	return url, nil
}

func TestSave(t *testing.T) {
	url := model.Url{}

	urlRepositoryMock := UrlRepositoryMock{}
	urlRepositoryMock.On("Save", url).Return(nil)

	urlService := NewUrlService(&urlRepositoryMock)

	urlService.Save(url)
	fmt.Println(url.Id)
	assert.Equal(t, 1, url.Id)

	urlRepositoryMock.AssertExpectations(t)
}
