package service

import (
	"errors"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/afamorim/go-shorter-url-domain/pkg/repository"
)

type UrlService interface {
}

type service struct {
	//urlRepository UrlRepository
}

var (
	urlRepository repository.UrlRepository
)

func NewUrlService(urlRepository *repository.UrlRepository) UrlService {
	//urlRepository = ur
	return &service{}
}

func (s *service) Save(url *model.Url) error {
	if url.OriginalUrl == "" {
		return errors.New("Original URL is mandatory")
	}
	err := urlRepository.Save(url)
	return err
}

func (s *service) FindByShorter(shoterUrl string) (interface{}, error) {

	url, err := urlRepository.FindByShorter(shoterUrl)
	return url, err
}
