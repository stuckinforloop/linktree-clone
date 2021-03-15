package shortener

type Redirect struct {
	Code    string `json:"code" bson:"code"`
	Handles string `json:"handles" bson:"handles" validate:"empty=false"`
}
