package datatype

type Post struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

type Iterator struct {
	index int
}

func (i *Iterator) Next() int {
	i.index++
	return i.index
}

func (i Iterator) Current() int {
	return i.index
}

func NewIterator() Iterator {
	return Iterator{
		index: 0,
	}
}

type Response struct {
	Success bool   `json:"success"`
	Data    string `json:"message"`
}

type SharkAttack struct {
	ID                   int    `json:"id,omitempty"`
	Activity             string `json:"activity,omitempty"`
	Age                  string `json:"age,omitempty"`
	Area                 string `json:"area,omitempty"`
	CaseNumber           string `json:"case_number,omitempty"`
	CaseNumber0          string `json:"case_number0,omitempty"`
	Country              string `json:"country,omitempty"`
	Date                 string `json:"date,omitempty"`
	FatalYN              string `json:"fatal_y_n,omitempty"`
	Href                 string `json:"href,omitempty"`
	HrefFormula          string `json:"href_formula,omitempty"`
	Injury               string `json:"injury,omitempty"`
	InvestigatorOrSource string `json:"investigator_or_source,omitempty"`
	Location             string `json:"location,omitempty"`
	Name                 string `json:"name,omitempty"`
	OriginalOrder        string `json:"original_order,omitempty"`
	Pdf                  string `json:"pdf,omitempty"`
	Sex                  string `json:"sex,omitempty"`
	Species              string `json:"species,omitempty"`
	Time                 string `json:"time,omitempty"`
	Type                 string `json:"type,omitempty"`
	Year                 string `json:"year,omitempty"`
}
