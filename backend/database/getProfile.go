package database

import (
	"fmt"
	"glow/shared"
)

func GetProfile(userId int) shared.Profile {

	var profile shared.Profile
	query := "select user_name,age,address,user_email from User where user_id = ? "
	err := Db.QueryRow(query, userId).Scan(&profile.Name, &profile.Age, &profile.Address, &profile.Email)
	fmt.Println("profile - ", profile)
	if err != nil {
		fmt.Println("error while finding profile details ")
	}
	return profile

}
