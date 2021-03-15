package shortener

type url string

type Redirect struct {
	Code    string `json:"code" bson:"code"`
	Handles []url  `json:"handles" bson:"handles" validate:"empty=false"`
}
