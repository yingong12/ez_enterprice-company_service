package utils

// map
var filterMap = []string{
	"district",
	"register_capital",
	"estimate_value",
	"industry",
}
var sortMap = []string{
	"register_capital",
	"estimate_value",
	"name",
}

func ParseFilter(k int) string {
	return filterMap[k]
}

func ParseSortColumn(k int) string {
	return sortMap[k]
}
