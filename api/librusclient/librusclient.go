package	librusclient

import (
	"net/http"
	"os"

	golibrus "github.com/0x113/librus-api-go"
)

// LibrusClient shares the global librus API client
var LibrusClient = &golibrus.Librus{
	Client:   &http.Client{},
	Username: os.Getenv("librus_username"),
	Password: os.Getenv("librus_password"),
}
