package utils

import "regexp"

func IsDuplicateKeyError(err error) bool {
	duplicate := regexp.MustCompile(`\(SQLSTATE 23505\)$`)
	ok := duplicate.MatchString(err.Error())
	if ok {
		return ok
	}
	duplicate = regexp.MustCompile(`UNIQUE constraint failed`)
	ok = duplicate.MatchString(err.Error())
	if ok {
		return ok
	}
	return false
}
