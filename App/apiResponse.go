package App


type ApiResponse struct {
	Code int  `bson:"code"`
	Type string  `bson:"type"`
	Message string  `bson:"message"`
}
