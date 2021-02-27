package meta

const (
	MsgSearchOptions = `
Type 'quit' to exit at anytime. Press 'Enter' to continue
	
	Select search options:
	 * Press 1 to search
	 * Press 2 to view a list of searchable fields
	 * Type 'quit' to exit
`

	MsgSelectItem  = `Select 1) Users or 2) Tickets or 3) Organizations`
	MsgSearchField = `Enter search term`
	MsgSearchValue = `Enter search value`
)

const (
	OptionSearch     = 1
	OptionViewFields = 2
	OptionQuit       = "quit"
)

const (
	ItemUsers         = 1
	ItemTickets       = 2
	ItemOrganizations = 3
)

const SearchNoResult = "No results found"

const DisplaySpaceLength = 50
