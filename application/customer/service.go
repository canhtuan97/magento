package customer

type Login struct {
	username string `json:"username"`
	password string `json:"password"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}
