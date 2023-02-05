package payloads

import (
	"net/http"

	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
)

type Image struct {
	owner   string    `json:"owner"`
	path    string    `json:"path"`
}

func (s Image) GetName() string {
	return "images"
}

func (s Image) GetAlias() string {
	return "u"
}

func (s Image) GetColumns() []string {
	return database.GetColumns[Image](true)
}

func (s Image) GetMandatoryColumns() []string {
	return database.GetColumns[Image](false)
}

func GetImageByOwner(w http.ResponseWriter, r *http.Request, owner string) *Image {
	cond := database.NewCondition("owner", database.EqualTo, owner)

	arr, err := database.Select[Image](r.Context(), cond)
	if err != nil {
		log.Logger.Errorf("couldn't get image: %+v", err)

		if w != nil {
			http.Error(w, "internal error: couldn't get image", http.StatusInternalServerError)
		}

		return nil
	}

	if len(arr) == 0 {
		log.Logger.Errorf("couldn't find image for owner id `%s`", owner)
		if w != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		return nil
	}

	return &arr[0]
}
