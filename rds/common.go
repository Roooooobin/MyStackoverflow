package rds

import "strconv"

func FormParentsKey(id int) string {
	return strconv.Itoa(id) + ":parents"
}
