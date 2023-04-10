package appl_row

type UserSendMessage struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type UserSendData struct {
	CardNumber string `json:"card_number"`
	Validity   string `json:"validity"`
	OwnerName  string `json:"owner_name"`
	CvvCode    string `json:"cvv_code"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}
