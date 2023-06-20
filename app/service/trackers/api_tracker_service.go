package trackers

import (
	entities "delos/app/db/entities"
	repo "delos/app/db/repositories/api_histories"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func TrackApi(method *string, path *string, user_agent *string) {
	paths := []string{"farm", "farms", "pond", "ponds"}
	realPath := strings.Split(*path, "/")[1]
	if slices.Contains(paths, realPath) {
		fmt.Println("{}, {}, {}", *method, *path, *user_agent)
		repo.CreateApiHistory(&entities.ApiHistoryEntity{Method: *method, Url: realPath, UserAgent: *user_agent})
	}
}
