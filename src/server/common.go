package app_backend

type PromoInfo struct {
	Name      string   `json: "name"`      // name of the restaurant
	Addr      string   `json: "addr"`      // address of the restaurant
	Phone     string   `json: "phone"`     // phone number
	Thumbnail string   `json: "thumbnail"` // url for thumbnail
	Images    []string `json: "images"`    // array of urls for images
}
