package database

import (
	"fmt"
	"landmarks/pkg/openapi"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetLandmarks(t *testing.T) {
	type args struct {
		name    string
		page    int
		perPage int
		offset  int
		end     int
	}
	testNames := map[string]args{}
	for r := 0; r < 100; r++ {
		itemSize := len(expectedItems)
		perPage := rand.Intn(itemSize) + 1
		maxPage := itemSize / perPage
		page := rand.Intn(maxPage) + 1
		offset := (page - 1) * perPage
		tname := fmt.Sprintf("Search %d:%d", page, perPage)
		if _, ok := testNames[tname]; !ok {
			testNames[tname] = args{
				name:    tname,
				page:    page,
				perPage: perPage,
				offset:  offset,
				end:     offset + perPage,
			}
		}
	}

	// for tname, args := range testNames {
	// 	t.Run(tname, func(t *testing.T) {
	// 		landmarks, err := GetLandmarks(expectedDB, args.page, args.perPage)
	// 		assert.NoError(t, err)
	// 		expected := expectedItems[args.offset:args.end]
	// 		for index, actual := range landmarks.Items {
	// 			diff := cmp.Diff(expected[index], actual, cmpopts.IgnoreFields(actual, "ID"))
	// 			assert.Empty(t, diff)
	// 		}
	// 	})
	// }
}

func TestSearchLandmarks(t *testing.T) {
	type args struct {
		db        *gorm.DB
		condition openapi.PostLandmarksSearchJSONRequestBody
	}
	tests := []struct {
		name    string
		args    args
		want    openapi.Landmarks
		wantErr bool
	}{
		{
			name: "Sort:Latitude:Asc",
			args: args{
				db: expectedDB,
				condition: openapi.LandmarkSearchQuery{
					Page:    1,
					PerPage: 5,
					Sort: &openapi.SortField{
						Field: "latitude",
						Type:  openapi.SortAscend,
					},
				},
			},
			want: openapi.Landmarks{
				Total: 47,
				Count: 5,
				Items: []openapi.Landmark{
					{Name: "那覇市", Description: "沖縄県の県庁所在地", Latitude: 26.212445, Longitude: 127.680922, Altitude: 0.0},
					{Name: "鹿児島市", Description: "鹿児島県の県庁所在地", Latitude: 31.560171, Longitude: 130.558025, Altitude: 0.0},
					{Name: "宮崎市", Description: "宮崎県の県庁所在地", Latitude: 31.911034, Longitude: 131.423887, Altitude: 0.0},
					{Name: "長崎市", Description: "長崎県の県庁所在地", Latitude: 32.75004, Longitude: 129.867251, Altitude: 0.0},
					{Name: "熊本市", Description: "熊本県の県庁所在地", Latitude: 32.7898, Longitude: 130.741584, Altitude: 0.0},
				},
			},
			wantErr: false,
		},
		{
			name: "name・後方一致・山市+Sort:Latitude:Desc",
			args: args{
				db: expectedDB,
				condition: openapi.LandmarkSearchQuery{
					Page:    1,
					PerPage: 10,
					Filter: &openapi.FilterField{
						Field: "name",
						Value: "山市",
						Type:  openapi.SuffixMatch,
					},
					Sort: &openapi.SortField{
						Field: "latitude",
						Type:  openapi.SortDescend,
					},
				},
			},
			want: openapi.Landmarks{
				Total: 47,
				Count: 4,
				Items: []openapi.Landmark{
					{Name: "富山市", Description: "富山県の県庁所在地", Latitude: 36.695265, Longitude: 137.211305, Altitude: 0.0},
					{Name: "岡山市", Description: "岡山県の県庁所在地", Latitude: 34.661739, Longitude: 133.935032, Altitude: 0.0},
					{Name: "和歌山市", Description: "和歌山県の県庁所在地", Latitude: 34.226111, Longitude: 135.1675, Altitude: 0.0},
					{Name: "松山市", Description: "愛媛県の県庁所在地", Latitude: 33.841642, Longitude: 132.765682, Altitude: 0.0},
				},
			},
			wantErr: false,
		},

		{
			name: "name・前方一致・福",
			args: args{
				db: expectedDB,
				condition: openapi.LandmarkSearchQuery{
					Page:    1,
					PerPage: 10,
					Filter: &openapi.FilterField{
						Field: "name",
						Value: "福",
						Type:  openapi.PrefixMach,
					},
				},
			},
			want: openapi.Landmarks{
				Total: 47,
				Count: 3,
				Items: []openapi.Landmark{
					{Name: "福島市", Description: "福島県の県庁所在地", Latitude: 37.750029, Longitude: 140.467771, Altitude: 0.0},
					{Name: "福井市", Description: "福井県の県庁所在地", Latitude: 36.065209, Longitude: 136.22172, Altitude: 0.0},
					{Name: "福岡市", Description: "福岡県の県庁所在地", Latitude: 33.606389, Longitude: 130.417968, Altitude: 0.0},
				},
			},
			wantErr: false,
		},
		{
			name: "name・完全一致・さいたま市",
			args: args{
				db: expectedDB,
				condition: openapi.LandmarkSearchQuery{
					Page:    1,
					PerPage: 10,
					Filter: &openapi.FilterField{
						Field: "name",
						Value: "さいたま市",
						Type:  openapi.ExtractMatch,
					},
				},
			},
			want: openapi.Landmarks{
				Total: 47,
				Count: 1,
				Items: []openapi.Landmark{
					{Name: "さいたま市", Description: "埼玉県の県庁所在地", Latitude: 35.857033, Longitude: 139.649012, Altitude: 0.0},
				},
			},
			wantErr: false,
		},
		{
			name: "name・後方一致・島市",
			args: args{
				db: expectedDB,
				condition: openapi.LandmarkSearchQuery{
					Page:    1,
					PerPage: 10,
					Filter: &openapi.FilterField{
						Field: "name",
						Value: "島市",
						Type:  openapi.SuffixMatch,
					},
				},
			},
			want: openapi.Landmarks{
				Total: 47,
				Count: 4,
				Items: []openapi.Landmark{
					{Name: "福島市", Description: "福島県の県庁所在地", Latitude: 37.750029, Longitude: 140.467771, Altitude: 0.0},
					{Name: "広島市", Description: "広島県の県庁所在地", Latitude: 34.396558, Longitude: 132.459646, Altitude: 0.0},
					{Name: "徳島市", Description: "徳島県の県庁所在地", Latitude: 34.065761, Longitude: 134.559286, Altitude: 0.0},
					{Name: "鹿児島市", Description: "鹿児島県の県庁所在地", Latitude: 31.560171, Longitude: 130.558025, Altitude: 0.0},
				},
			},
			wantErr: false,
		},
		{
			name: "description・部分一致・知",
			args: args{
				db: expectedDB,
				condition: openapi.LandmarkSearchQuery{
					Page:    1,
					PerPage: 10,
					Filter: &openapi.FilterField{
						Field: "description",
						Value: "知",
						Type:  openapi.PartialMatch,
					},
				},
			},
			want: openapi.Landmarks{
				Total: 47,
				Count: 2,
				Items: []openapi.Landmark{
					{Name: "名古屋市", Description: "愛知県の県庁所在地", Latitude: 35.180209, Longitude: 136.906582, Altitude: 0.0},
					{Name: "高知市", Description: "高知県の県庁所在地", Latitude: 33.559722, Longitude: 133.531111, Altitude: 0.0},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchLandmarks(tt.args.db, tt.args.condition)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchLandmarks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got.Count, tt.want.Count)
			assert.Len(t, got.Items, len(tt.want.Items))
			for idx, item := range got.Items {
				cmp.Equal(item, tt.want.Items[idx], cmpopts.IgnoreFields(openapi.Landmark{}, "ID"))
			}
		})
	}
}
