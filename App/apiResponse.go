package App


type ApiResponse struct {
	Code int  `bson:"code"`
	Type string  `bson:"type"`
	Message interface{}  `bson:"message"`
}
