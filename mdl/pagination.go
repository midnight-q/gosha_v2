package mdl

type Pagination struct {
	CurrentPage int
	PerPage     int
}

func (pagination Pagination) GetOffset() int {
	return (pagination.CurrentPage - 1) * pagination.PerPage
}
