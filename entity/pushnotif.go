package entity

// PushNotif is use to hold information
// status 0 is not sent
// status 1 is sent
type PushNotif struct {
	Status      int8   `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
