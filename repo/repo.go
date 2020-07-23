package repo

import (
	"encoding/json"
	"io/ioutil"
	"fmt"

	"go-graph-starter/model"
)

type Repo struct {
}

//GetData ... gets the data from repository
func (r *Repo) GetData() []model.User {
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Print("Error:", err)
	}

	var result map[string]model.User
	if err = json.Unmarshal(content, &result); err != nil {
		fmt.Print("Error:", err)
	}
	users := []model.User{}
	for _, v := range result { 
		users = append(users, v)
	 }
	return users
}
