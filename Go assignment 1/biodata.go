package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct { // membuat struct person yang di dalamnya terdapat beberapa data yang betipe string dan int
	name, address, job, reason string
}

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	var allStudentClass = []person{
		{
			name:    "Auliya",
			address: "Madiun",
			job:     "Penyanyi",
			reason:  "Gabut",
		},
		{
			name:    "Arkhan",
			address: "Magetan",
			job:     "Apapun yang penting halal dan banyak duit",
			reason:  "Capek, pingin yang strict aja",
		},
		{
			name:    "Wardoyo",
			address: "Yogyakarta",
			job:     "Pengusaha",
			reason:  "Rahasia"},
		{
			name:    "Billie Eilish",
			address: "Nganjuk",
			job:     "Juragan Lombok",
			reason:  "Ingin menguasai dunia"},
		{
			name:    "Elon Musk",
			address: "Cepu",
			job:     "Bengkel",
			reason:  "Ingin jadi orang terkaya",
		},
	}
	for i, student := range allStudentClass {
		num := input - 1
		if num == i {
			fmt.Printf("Nama: %s\nAlamat: %s\njob: %s\nreason: %s", student.name, student.address, student.job, student.reason)
		}
	}
}
