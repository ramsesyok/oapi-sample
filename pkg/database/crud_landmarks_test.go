package database

import (
	"errors"
	"landmarks/pkg/openapi"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/gorm"
)

// TestCreateLandmark 地点情報レコード作成のテスト
// 1. 指定データの登録（正常）
func TestCreateLandmark(t *testing.T) {
	type args struct {
		db       *gorm.DB
		landmark openapi.PostLandmarksJSONRequestBody
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "東京タワー登録",
			args: args{
				db: expectedDB,
				landmark: openapi.Landmark{
					Name:        "東京タワー",
					Latitude:    35.6586193045004,
					Longitude:   139.7454050822132,
					Altitude:    333.0,
					Description: "東京都港区芝公園にある総合電波塔で、正式名称は日本電波塔である。",
				},
			},
			want:    48,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateLandmark(tt.args.db, tt.args.landmark)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostLandmarks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PostLandmarks() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestReadLandmark 地点情報レコード読み込みのテスト
// 1. 存在しないデータ読み込み(NotFound)
// 2. テストデータ上の全点の読み込み及び内容チェック（正常）
func TestReadLandmark(t *testing.T) {
	type args struct {
		db *gorm.DB
		ID int
	}
	tests := []struct {
		name    string
		args    args
		want    openapi.Landmark
		wantErr bool
	}{
		{
			name: "存在しないデータ",
			args: args{
				ID: 1000,
				db: expectedDB,
			},
			wantErr: true,
		},
	}
	for id, expected := range expectedItems {
		tests = append(tests, struct {
			name    string
			args    args
			want    openapi.Landmark
			wantErr bool
		}{
			name: expected.Name, args: args{db: expectedDB, ID: id + 1}, want: expectedItems[id], wantErr: false,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLandmark(tt.args.db, tt.args.ID)

			if (err != nil) != tt.wantErr {
				//エラーが発生しないケースでエラーが出た場合
				t.Errorf("ReadLandmark() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					// エラーが発生するケースで、想定通りのエラーでない場合
					t.Errorf("ReadLandmark() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !cmp.Equal(got, tt.want, cmpopts.IgnoreFields(got, "ID")) {
				// 取得できたデータが一致しない場合
				t.Errorf("ReadLandmark() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestUpdateLandmark 地点レコード更新のテスト
func TestUpdateLandmark(t *testing.T) {
	type args struct {
		db        *gorm.DB
		id        int
		landmarks openapi.PutLandmarksIDJSONRequestBody
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "存在しないデータ",
			args: args{
				id: 1000,
				db: expectedDB,
			},
			wantErr: true,
		},
	}
	for id, expected := range expectedItems {
		item := expectedItems[id]
		item.Description = "テスト"
		tests = append(tests, struct {
			name    string
			args    args
			wantErr bool
		}{
			name: expected.Name, args: args{db: expectedDB, id: id + 1, landmarks: item}, wantErr: false,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateLandmark(tt.args.db, tt.args.id, tt.args.landmarks); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLandmark() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				actual := openapi.Landmark{}
				if taked := expectedDB.Take(&actual, tt.args.id); taked.Error != nil {
					if (err != nil) != tt.wantErr {
						t.Errorf("UpdateLandmark() error = %v, wantErr %v", err, tt.wantErr)
					}
				}
				if actual.Description != tt.args.landmarks.Description {
					// 取得できたデータが一致しない場合
					t.Errorf("UpdateLandmark() = %v, want %v", actual, tt.args.landmarks)
				}
			}
		})
	}
}

// TestDeleteLandmarksId 地点レコードの削除のテスト
// 1. 存在しないデータの削除(NotFound)
// 2. 2件目のデータの削除(正常系,青森)
func TestDeleteLandmarksId(t *testing.T) {
	type args struct {
		db *gorm.DB
		id int
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
	}{
		{
			name: "削除失敗",
			args: args{
				db: expectedDB,
				id: len(testData) + 100,
			},
			wantErr:     true,
			expectedErr: gorm.ErrRecordNotFound,
		},
	}
	for id, expected := range expectedItems {
		tests = append(tests, struct {
			name        string
			args        args
			wantErr     bool
			expectedErr error
		}{
			name: expected.Name, args: args{db: expectedDB, id: id + 1}, wantErr: false, expectedErr: nil,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteLandmark(tt.args.db, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteLandmarksId() error = %v, wantErr %v", err, tt.wantErr)
			} else if !errors.Is(err, tt.expectedErr) {
				t.Errorf("DeleteLandmarksId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
