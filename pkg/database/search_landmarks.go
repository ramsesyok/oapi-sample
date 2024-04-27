package database

import (
	"landmarks/pkg/openapi"

	"gorm.io/gorm"
)

func paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}
		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func GetLandmarks(db *gorm.DB, page int, perPage int) (openapi.Landmarks, error) {

	landmarks := openapi.Landmarks{
		Items: []openapi.Landmark{},
	}
	if tx := db.Scopes(paginate(page, perPage)).Find(&landmarks.Items); tx.Error != nil {
		return landmarks, tx.Error
	}
	var total int64
	_ = db.Find(&[]openapi.Landmark{}).Count(&total)
	landmarks.Total = int(total)
	landmarks.Count = len(landmarks.Items)
	return landmarks, nil
}

func SearchLandmarks(db *gorm.DB, condition openapi.PostLandmarksSearchJSONRequestBody) (openapi.Landmarks, error) {

	landmarks := openapi.Landmarks{
		Items: []openapi.Landmark{},
	}
	page := condition.Page
	perPage := condition.PerPage

	search := db.Debug().Scopes(paginate(page, perPage))
	if condition.Filter != nil {
		query := condition.Filter.Field + " LIKE ?"
		switch condition.Filter.Type {
		case openapi.PrefixMach:
			search = search.Where(query, condition.Filter.Value+"%")
		case openapi.SuffixMatch:
			search = search.Where(query, "%"+condition.Filter.Value)
		case openapi.ExtractMatch:
			search = search.Where(query, condition.Filter.Value)
		case openapi.PartialMatch:
			search = search.Where(query, "%"+condition.Filter.Value+"%")
		}
	}

	if condition.Sort != nil {
		switch condition.Sort.Type {
		case openapi.SortAscend:
			search = search.Order(condition.Sort.Field + " asc")
		case openapi.SortDescend:
			search = search.Order(condition.Sort.Field + " desc")
		}
	}
	searched := search.Find(&landmarks.Items)
	if searched.Error != nil {
		return landmarks, searched.Error
	}
	var total int64
	_ = db.Find(&[]openapi.Landmark{}).Count(&total)
	landmarks.Total = int(total)
	landmarks.Count = len(landmarks.Items)
	return landmarks, nil
}
