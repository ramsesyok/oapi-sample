package database

import (
	"landmarks/pkg/openapi"

	"gorm.io/gorm"
)

// CreateLandmark データベースに地点情報レコードを作成します.
func CreateLandmark(db *gorm.DB, landmark openapi.PostLandmarksJSONRequestBody) (int, error) {
	created := db.Create(&landmark)
	return *landmark.ID, created.Error
}

// ReadLandmark データベースから地点情報レコードの取得します.
func ReadLandmark(db *gorm.DB, ID int) (openapi.Landmark, error) {
	landmark := openapi.Landmark{}
	tx := db.Take(&landmark, ID)
	return landmark, tx.Error
}

// UpdateLandmark データベースの地点情報レコードを更新します.
func UpdateLandmark(db *gorm.DB, id int, landmarks openapi.PutLandmarksIDJSONRequestBody) error {

	current := openapi.Landmark{}
	taked := db.Take(&current, id)
	if taked.Error != nil {
		return taked.Error
	}
	landmarks.ID = &id
	updated := db.Save(&landmarks)
	return updated.Error
}

// DeleteLandmark データベースの地点情報レコードを削除します.
func DeleteLandmark(db *gorm.DB, id int) error {

	landmark := openapi.Landmark{}
	taked := db.Take(&landmark, id)
	if taked.Error != nil {
		return taked.Error
	}
	deleted := db.Delete(&landmark, id)
	return deleted.Error
}
