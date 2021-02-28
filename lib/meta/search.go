package meta

type SearchEntry struct {
	Field string
	Value string
}

const SearchNoResult = "No results found"

const (
	DisplaySpaceLength = 50
	DisplayDashLength  = 50
)

const MsgRetrySearching = "\nPress 1 to retry searching, the other number to back to menu"

const RetrySearchingInput = 1
