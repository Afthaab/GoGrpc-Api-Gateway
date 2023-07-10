package domain

type Size struct {
	ID    string ` json:"id"`
	Name  string ` json:"name"`
	Price string ` json:"price"`
}

type Category struct {
	Id           string `json:"id"`
	Categoryname string `json:"categoryname"`
	Imageurl     string `json:"imageurl"`
}

type Products struct {
	Id           string   ` json:"id"`
	Productname  string   ` json:"productname"`
	Calories     string   ` json:"calories"`
	Availibility bool     ` json:"availibility"`
	Categoryid   string   ` json:"categoryid"`
	Typeid       string   ` json:"typeid"`
	Baseprice    float64  ` json:"baseprice"`
	Sizeid       []string ` json:"sizeid"`
	Imageurls    []string ` json:"imageurls"`
}
type Foodtype struct {
	Id       string `json:"id"`
	Foodtype string `json:"foodtype"`
	Imageurl string `json:"imageurl"`
}
