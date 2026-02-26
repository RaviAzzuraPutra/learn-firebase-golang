package model

import "time"

type Member struct {
	ID          string    `json:"id" firestore:"id"`
	Name        *string   `json:"name" firestore:"name"`
	Email       *string   `json:"email" firestore:"email"`
	Phone       *string   `json:"phone" firestore:"phone"`
	Join_Date   time.Time `json:"join_date" firestore:"join_date"`
	Member_Code string    `json:"member_code" firestore:"member_code"`
}
