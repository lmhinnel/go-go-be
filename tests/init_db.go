package tests

import (
	"time"

	"github.com/lmhuong711/go-go-be/entities"
	"github.com/lmhuong711/go-go-be/services"
)

func InitTestDB() {
	bd, _ := time.Parse("2006-01-02", "2000-01-01")

	fixedTime, _ := time.Parse("2006-01-02T15:04:05.000000Z", "2024-05-22T19:00:00.000000+07:00")
	students := []entities.Students{
		{
			ID:         1,
			FirstName:  "Perry",
			LastName:   "Platypus",
			Phone:      "84987654320",
			Email:      "perry_platypus@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Tech University",
			Address:    "1 Elm Street",
			CreatedAt:  fixedTime,
			UpdatedAt:  fixedTime,
		},
		{
			FirstName:  "Phineas",
			LastName:   "Flynn",
			Phone:      "84987654321",
			Email:      "phineas_flynn@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Abc University",
			Address:    "2 Elm Street",
		},
		{
			FirstName:  "Ferb",
			LastName:   "Fletcher",
			Phone:      "84987654322",
			Email:      "ferb_fletcher@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Xyz University",
			Address:    "3 Elm Street",
		},
		{
			FirstName:  "Candace",
			LastName:   "Flynn",
			Phone:      "84987654323",
			Email:      "candace_flynn@gmail.com",
			Birthday:   bd,
			Gender:     0,
			University: "Tech University",
			Address:    "4 Elm Street",
		},
		{
			FirstName:  "Isabella",
			LastName:   "Garcia-Shapiro",
			Phone:      "84987654324",
			Email:      "isabella_gracia_shapiro@gmail.com",
			Birthday:   bd,
			Gender:     0,
			University: "Abc University",
			Address:    "5 Elm Street",
		},
		{
			FirstName:  "Buford",
			LastName:   "Van Stomm",
			Phone:      "84987654325",
			Email:      "buford_van_stomm@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Xyz University",
			Address:    "6 Elm Street",
		},
		{
			FirstName:  "Baljeet",
			LastName:   "Tjinder",
			Phone:      "84987654326",
			Email:      "baljeet_tjinder@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Tech University",
			Address:    "7 Elm Street",
		},
		{
			FirstName:  "Jeremy",
			LastName:   "Johnson",
			Phone:      "84987654327",
			Email:      "jeremy_johnson@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Abc University",
			Address:    "8 Elm Street",
		},
		{
			FirstName:  "Stacy",
			LastName:   "Hirano",
			Phone:      "84987654328",
			Email:      "stacy_hirano@gmail.com",
			Birthday:   bd,
			Gender:     0,
			University: "Xyz University",
			Address:    "9 Elm Street",
		},
		{
			FirstName:  "Vanessa",
			LastName:   "Doofenshmirtz",
			Phone:      "84987654329",
			Email:      "vanessa_doffenshmirtz@gmail.com",
			Birthday:   bd,
			Gender:     0,
			University: "Tech University",
			Address:    "10 Elm Street",
		},
		{
			FirstName:  "Dr. Heinz",
			LastName:   "Doofenshmirtz",
			Phone:      "84987654330",
			Email:      "heinz_doffenshmirtz@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Abc University",
			Address:    "11 Elm Street",
		},
		{
			FirstName:  "Major",
			LastName:   "Monogram",
			Phone:      "84987654331",
			Email:      "major_monogram@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Xyz University",
			Address:    "12 Elm Street",
		},
		{
			FirstName:  "Carl",
			LastName:   "Karl",
			Phone:      "84987654332",
			Email:      "carl_karl@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Tech University",
			Address:    "13 Elm Street",
		},
		{
			FirstName:  "Norm",
			LastName:   "Norm",
			Phone:      "84987654333",
			Email:      "norm_norm@gmail.com",
			Birthday:   bd,
			Gender:     1,
			University: "Abc University",
			Address:    "14 Elm Street",
		},
	}

	for _, student := range students {
		services.CreateStudent(&student)
	}

}
