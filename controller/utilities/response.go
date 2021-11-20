package controller

// HTTPError defines the field which should be the response.
type HTTPError struct {
	Status int `json:"status" example:"error"`          
	Error  string `json:"error" example:"a message error"` 
}

// HTTPSuccess defines the field which should be the response.
type HTTPSuccess struct {
	Status int      `json:"status" example:"success"` 
	Data   interface{} `json:"data"`                     
}