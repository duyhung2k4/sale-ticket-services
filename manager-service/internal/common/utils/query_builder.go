package utils

import (
	"fmt"
	"strings"
)

func AddLikeClause(field string) string {
	return fmt.Sprintf("%s LIKE ?", field)
}

func MergeConditionAND(args ...string) string {
	return strings.Join(args, " AND ")
}
