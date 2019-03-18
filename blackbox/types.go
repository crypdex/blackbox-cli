package blackbox

type InitRequest struct {
	Mnemonic string `json:"mnemonic"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Force    bool   `json:"force"`
}

type InitResponse struct {
	Mnemonic string `json:"mnemonic"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateAddressRequest struct {
	Rescan    bool `json:"rescan"`
	Watchonly bool `json:"watchonly"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}
