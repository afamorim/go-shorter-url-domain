package repository

import "github.com/afamorim/go-shorter-url-domain/pkg/model"

type UrlRepository interface {
	Save(url model.Url) error
	FindByShorter(shorterUrl string) (model.Url, error)
}
