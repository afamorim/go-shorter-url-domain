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

func (r UrlRepositoryMock) Save(url *model.Url) error {
	url.Id = 1
	return nil
}

func (r UrlRepositoryMock) FindByShorter(shorterUrl string) (interface{}, error) {
	url := model.Url{
		Id:          1,
		OriginalUrl: "http://www.teste.com.br",
		CompressUrl: "http://yougotit.com",
	}

	return url, nil
}

func TestSave(t *testing.T) {
	urlRepositoryMock := UrlRepositoryMock{}
	urlService := NewUrlService(urlRepositoryMock)
	url := model.Url{}
	urlService.Save(&url)
	fmt.Println(url.Id)
	assert.Equal(t, 1, url.Id)
	//t.Error("ERROR")
}
