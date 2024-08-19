package internal

import (
	"strconv"
	"strings"
)

func Suggestions(dataset []Artist) map[string]string {
	opt := make(map[string]string)

	for _, val := range dataset {

		opt[val.Name] = " -> name"
		// search by members

		for _, m := range val.Members {
			opt[m] = " -> member"
		}
		opt[strconv.Itoa(val.CreationDate)] = " -> creation date"

		opt[val.FirstAlbum] = " -> first album"
		// dates and locations search

		for key, elem := range val.DatesLocations {
			opt[key] = " -> concert location"

			for _, d := range elem {
				opt[d] = " -> concert date"
			}
		}
	}

	return opt
}

func SearchBar(list []Artist, targetStr string) []Artist {
	var newList []Artist
	target := strings.ToLower(targetStr)
	target_arr := strings.Split(target, " -> ")
	flag := false
	for _, val := range list {
		flag = false
		// search by name
		if strings.Contains(strings.ToLower(val.Name), target_arr[0]) {
			flag = true
		}
		// search by members
		for _, members := range val.Members {
			if strings.Contains(strings.ToLower(members), target_arr[0]) {
				flag = true
			}
		}
		// search by creation date
		if strings.Contains(strings.ToLower(strconv.Itoa(val.CreationDate)), target_arr[0]) {
			flag = true
		}
		// first albume date
		if strings.Contains(strings.ToLower(val.FirstAlbum), target_arr[0]) {
			flag = true
		}
		// dates and locations search
		for key, element := range val.DatesLocations {
			if strings.Contains(strings.ToLower(key), target_arr[0]) {
				flag = true
			}
			for _, vals := range element {
				if strings.Contains(strings.ToLower(vals), target_arr[0]) {
					flag = true
				}
			}
		}
		if flag {
			newList = append(newList, val)
		}
	}
	return newList
}
