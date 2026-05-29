package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type paymentUsecase struct {
	paymentRepository domain.PaymentRepository
}

func PaymentUsecase(paymentRepository domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{paymentRepository: paymentRepository}
}

func (u *paymentUsecase) Get() (data []model.Payment, message, detail string, err error) {
	data, err = u.paymentRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get payment"
	return
}

func (u *paymentUsecase) GetByID(id int) (data model.Payment, message, detail string, err error) {
	data, err = u.paymentRepository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get payment by ID"
	return
}
