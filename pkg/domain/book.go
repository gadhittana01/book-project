package domain

type GetListOfBooksResp struct {
	Books []Book `json:"works"`
}

type Book struct {
	Key               string   `json:"key"`
	Title             string   `json:"title"`
	EditionCount      int      `json:"edition_count"`
	Authors           []Author `json:"authors"`
	LendingIdentifier string   `json:"lending_identifier"`
}

type Author struct {
	Name string `json:"name"`
}

type GetListOfBooksReq struct {
	Subject string `json:"subject"`
}

type BorrowBookReq struct {
	Book       Book   `json:"book"`
	PickUpDate string `json:"pickup_date"`
	UserID     int    `json:"user_id"`
}

type GeBookByKeyReq struct {
	Key     string `json:"key"`
	Subject string `json:"subject"`
}

type GetBookReservationReq struct {
	UserID int `json:"user_id"`
}
