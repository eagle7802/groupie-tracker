package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Unmarshal() ([]Artist, error) {
	art_parse, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	defer art_parse.Body.Close()

	body, err := io.ReadAll(art_parse.Body)
	if err != nil {
		return nil, err
	}
	var Art []Artist

	err = json.Unmarshal(body, &Art)
	if err != nil {
		return nil, err
	}

	rel_parse, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer rel_parse.Body.Close()

	body_rel, err := io.ReadAll(rel_parse.Body)

	var Rel Relation
	err = json.Unmarshal(body_rel, &Rel)
	if err != nil {
		return nil, err
	}

	for i := range Rel.Index {
		newMap := make(map[string][]string)
		for key, val := range Rel.Index[i].DatesLocations {
			newKey := deleteSymbols(key)
			newMap[newKey] = val
		}
		Rel.Index[i].DatesLocations = newMap
	}

	for i := range Art {
		Art[i].DatesLocations = Rel.Index[i].DatesLocations
	}
	return Art, nil
}
func deleteSymbols(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", ", ")
	s = strings.ToTitle(s)
	return s
}
