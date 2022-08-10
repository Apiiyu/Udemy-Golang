package main

import "fmt"

func main() {
	student := map[string]string{
		"name":       "Rafi Khoirulloh",
		"class":      "Sistem Informasi",
		"university": "Universitas Widyatama",
	}

	privacyData := map[string]float64{
		"IP":  3.70,
		"IPK": 3.80,
		"GPA": 3.90,
	}

	if student["university"] == "Universitas Widyatama" && privacyData["IP"] < 3.60 {
		fmt.Println("Student is from Widyatama")
	} else if student["university"] == "Universitas Widyatama" || privacyData["IP"] == 3.70 {
		fmt.Println("Cumlaude Student from Universitas Widyatama")
	} else {
		fmt.Println("Student is from other university")
	}

	//	--> If short hand statement
	if totalPrivacyData := len(privacyData); totalPrivacyData > 3 {
		fmt.Println("Total privacy data is more than 3")
	} else {
		fmt.Println("Total privacy data is less than 3")
	}
}
