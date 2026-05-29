package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type paymentUsecase struct {
	paymentRepository domain.PaymentRepository
}

func PaymentUsecase(repo domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{paymentRepository: repo}
}

func (p *paymentUsecase) Create(request model.CreatePayment) (err error) {

	req := model.Payment{
		IDAppointment:    request.IdAppointment,
		Total:            request.Total,
		MetodePembayaran: request.MetodePembayaran,
		StatusPembayaran: request.StatusPembayaran,
	}

	tx := p.paymentRepository.Begin()

	err = p.paymentRepository.Create(tx, req)
	if err != nil {
		p.paymentRepository.Rollback(tx)
		return
	}

	p.paymentRepository.Commit(tx)

	return
}