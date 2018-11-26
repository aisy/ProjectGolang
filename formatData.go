package main

// UsersFormat : format json pada data Users ketika dia menampilkan data
type UsersFormat struct {
	IDUSers     string `form:"idusers" json:"idusers"`
	NamaLengkap string `form:"namalengkap" json:"namalengkap"`
	NoTelp      string `form:"notlp" json:"notlp"`
}

// ResponseUsers : adalah result yang akan ditampilkan pada web render
type ResponseUsers struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []UsersFormat
}
