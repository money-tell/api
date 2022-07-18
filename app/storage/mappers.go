package storage

import "github.com/katalabut/money-tell-api/app/entities"

func mapPayDateRow(results []*payDateRow) []*entities.PayDate {
	var pays []*entities.PayDate

	for _, pay := range results {
		pays = append(pays, &entities.PayDate{
			Date:  pay.Date,
			Items: mapPays(pay.Items),
		})
	}

	return pays
}

func mapPays(results []*payRow) []*entities.Pay {
	var pays []*entities.Pay

	for _, pay := range results {
		pays = append(pays, mapPay(pay))
	}

	return pays
}

func mapPay(pay *payRow) *entities.Pay {
	return &entities.Pay{
		Title:       pay.Title,
		Price:       pay.Price,
		PaymentDate: pay.PaymentDate,
	}
}

/* USER */

func mapUserRow(user *userRow) *entities.User {
	return &entities.User{
		ID:       user.ID.Hex(),
		Email:    user.Email,
		Password: user.Password,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}
