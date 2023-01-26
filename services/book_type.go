package services

type BookDependencies struct {
	BR BookResource
}

type GetListOfBooksReq struct {
	Subject string `json:"subject"`
}

type GetListOfBooksResp struct {
	Books []Book `json:"books"`
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

type BorrowBookReq struct {
	BookKey    string `json:"key"`
	PickUpDate string `json:"pickup_date"`
	Subject    string `json:"subject"`
	UserID     int    `json:"user_id"`
}

type BorrowBookRes struct {
	Book       Book   `json:"book"`
	PickUpDate string `json:"pickup_date"`
	UserID     int    `json:"user_id"`
}

type GetBookReservationReq struct {
	UserID int `json:"user_id"`
}

type GetBookReservationRes struct {
	BookKey    string `json:"key"`
	PickUpDate string `json:"pickup_date"`
	UserID     int    `json:"user_id"`
}
