package dtos

type Item struct {
	Id       	 string
	User_id		 string
	App_name     string
	App_email 	 string
	App_password string
}

type ItemData struct {
	App_name string
}

type GetItemData struct {
	Email   string `json:"email"`
	AppName string `json:"appName"`
}

type AddItemData struct {
	UserId   		string `json:"userID"`
	AppName  		string `json:"appName"`
	AppEmail  		string `json:"appEmail"`
	AppPassword  	string `json:"appPassword"`
}

type DeleteItemData struct {
	UserId   		string `json:"userID"`
	ItemId  		string `json:"itemID"`
}

type UpdateItemData struct {
	UserId   		string `json:"userID"`
	AppName  		string `json:"appName"`
	AppEmail  		string `json:"appEmail"`
	AppPassword  	string `json:"appPassword"`
}
