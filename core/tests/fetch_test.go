package tests

import (
	"errors"
	"testing"
	"time"

	"core-healtcare.com/model"
	"core-healtcare.com/usecase"
	"gorm.io/gorm"
)

var errNotFound = errors.New("record not found")
var errDB = errors.New("database error")

type mockMedicineRepo struct {
	createErr      error
	getErr         error
	getData        []model.Medicine
	findErr        error
	findData       model.Medicine
	commitErr      error
	rollbackCalled bool
}

func (m *mockMedicineRepo) Begin() *gorm.DB                           { return &gorm.DB{} }
func (m *mockMedicineRepo) Commit(_ *gorm.DB) error                   { return m.commitErr }
func (m *mockMedicineRepo) Rollback(_ *gorm.DB) error                 { m.rollbackCalled = true; return nil }
func (m *mockMedicineRepo) Create(_ *gorm.DB, _ model.Medicine) error { return m.createErr }
func (m *mockMedicineRepo) Get() ([]model.Medicine, error)            { return m.getData, m.getErr }
func (m *mockMedicineRepo) FindByID(_ int) (model.Medicine, error)    { return m.findData, m.findErr }

func TestMedicineUsecase_Create(t *testing.T) {
	expiredDate := time.Now().AddDate(1, 0, 0).Format("2006-01-02")

	tests := []struct {
		name      string
		request   model.CreateMedicine
		createErr error
		wantErr   bool
	}{
		{
			name: "create_success",
			request: model.CreateMedicine{
				NamaObat:    "Amoxicillin",
				Kategori:    "Antibiotik",
				Stok:        50,
				Harga:       15000,
				ExpiredDate: expiredDate,
			},
			createErr: nil,
			wantErr:   false,
		},
		{
			name: "create_db_error",
			request: model.CreateMedicine{
				NamaObat:    "Paracetamol",
				Kategori:    "Analgesik",
				Stok:        100,
				Harga:       2000,
				ExpiredDate: expiredDate,
			},
			createErr: errDB,
			wantErr:   true,
		},
		{
			name: "create_invalid_date_format",
			request: model.CreateMedicine{
				NamaObat:    "Ibuprofen",
				Kategori:    "NSAID",
				Stok:        30,
				Harga:       5000,
				ExpiredDate: "not-a-date",
			},
			createErr: nil,
			wantErr:   false,
		},
		{
			name: "create_zero_stock",
			request: model.CreateMedicine{
				NamaObat:    "Vitamin C",
				Kategori:    "Vitamin",
				Stok:        0,
				Harga:       3000,
				ExpiredDate: expiredDate,
			},
			createErr: nil,
			wantErr:   false,
		},
		{
			name: "create_zero_price",
			request: model.CreateMedicine{
				NamaObat:    "Sample Drug",
				Kategori:    "Sample",
				Stok:        10,
				Harga:       0,
				ExpiredDate: expiredDate,
			},
			createErr: nil,
			wantErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockMedicineRepo{createErr: tc.createErr}
			uc := usecase.MedicineUsecase(repo)

			err := uc.Create(tc.request)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
			// When create fails, rollback should have been called
			if tc.createErr != nil && !repo.rollbackCalled {
				t.Errorf("expected Rollback to be called on create failure")
			}
		})
	}
}

func TestMedicineUsecase_Get(t *testing.T) {
	sampleMedicines := []model.Medicine{
		{IDObat: 1, NamaObat: "Paracetamol", Kategori: "Analgesik", Stok: 100, Harga: 2000},
		{IDObat: 2, NamaObat: "Amoxicillin", Kategori: "Antibiotik", Stok: 50, Harga: 15000},
	}

	tests := []struct {
		name        string
		getData     []model.Medicine
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     sampleMedicines,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "succes get medicine",
		},
		{
			name:        "get_empty_list",
			getData:     []model.Medicine{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "succes get medicine",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockMedicineRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.MedicineUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d items, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected detail to be non-empty on error")
			}
		})
	}
}

func TestMedicineUsecase_GetByID(t *testing.T) {
	sample := model.Medicine{IDObat: 1, NamaObat: "Paracetamol", Kategori: "Analgesik", Stok: 100, Harga: 2000}

	tests := []struct {
		name        string
		id          int
		findData    model.Medicine
		findErr     error
		wantErr     bool
		wantMessage string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get medicine",
		},
		{
			name:        "get_by_id_not_found",
			id:          999,
			findData:    model.Medicine{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
		{
			name:        "get_by_id_db_error",
			id:          1,
			findData:    model.Medicine{},
			findErr:     errDB,
			wantErr:     true,
			wantMessage: "data not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockMedicineRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.MedicineUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.IDObat != tc.findData.IDObat {
				t.Errorf("expected IDObat %d, got %d", tc.findData.IDObat, data.IDObat)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockPatientRepo struct {
	getData  []model.Patient
	getErr   error
	findData model.Patient
	findErr  error
}

func (m *mockPatientRepo) Get() ([]model.Patient, error)         { return m.getData, m.getErr }
func (m *mockPatientRepo) FindByID(_ int) (model.Patient, error) { return m.findData, m.findErr }

func TestPatientUsecase_Get(t *testing.T) {
	samplePatients := []model.Patient{
		{IDPatient: 1, Nama: "Budi Santoso", NIK: "3201010101010001", Gender: "Laki-laki"},
		{IDPatient: 2, Nama: "Sari Dewi", NIK: "3201010101010002", Gender: "Perempuan"},
	}

	tests := []struct {
		name        string
		getData     []model.Patient
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     samplePatients,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "success get patient",
		},
		{
			name:        "get_empty",
			getData:     []model.Patient{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get patient",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockPatientRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.PatientUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d patients, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestPatientUsecase_GetByID(t *testing.T) {
	sample := model.Patient{IDPatient: 1, Nama: "Budi Santoso", NIK: "3201010101010001"}

	tests := []struct {
		name        string
		id          int
		findData    model.Patient
		findErr     error
		wantErr     bool
		wantMessage string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get patient by ID",
		},
		{
			name:        "get_by_id_not_found",
			id:          999,
			findData:    model.Patient{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockPatientRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.PatientUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.IDPatient != tc.findData.IDPatient {
				t.Errorf("expected IDPatient %d, got %d", tc.findData.IDPatient, data.IDPatient)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockAppointmentRepo struct {
	getData  []model.Appointment
	getErr   error
	findData model.Appointment
	findErr  error
}

func (m *mockAppointmentRepo) Get() ([]model.Appointment, error) { return m.getData, m.getErr }
func (m *mockAppointmentRepo) FindByID(_ int) (model.Appointment, error) {
	return m.findData, m.findErr
}

func TestAppointmentUsecase_Get(t *testing.T) {
	sampleAppointments := []model.Appointment{
		{
			IDAppointment: 1,
			IDPatient:     1,
			IDDoctor:      2,
			Keluhan:       "Demam tinggi",
			TekananDarah:  "120/80",
			SuhuTubuh:     38.5,
			BeratBadan:    65.0,
			Status:        "Menunggu",
		},
	}

	tests := []struct {
		name        string
		getData     []model.Appointment
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     sampleAppointments,
			getErr:      nil,
			wantErr:     false,
			wantCount:   1,
			wantMessage: "success get appointment",
		},
		{
			name:        "get_empty",
			getData:     []model.Appointment{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get appointment",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockAppointmentRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.AppointmentUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d appointments, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestAppointmentUsecase_GetByID(t *testing.T) {
	sample := model.Appointment{IDAppointment: 1, Keluhan: "Demam tinggi", Status: "Menunggu"}

	tests := []struct {
		name        string
		id          int
		findData    model.Appointment
		findErr     error
		wantErr     bool
		wantMessage string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get appointment by ID",
		},
		{
			name:        "get_by_id_not_found",
			id:          0,
			findData:    model.Appointment{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
		{
			name:        "get_by_id_negative_id",
			id:          -1,
			findData:    model.Appointment{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockAppointmentRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.AppointmentUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.IDAppointment != tc.findData.IDAppointment {
				t.Errorf("expected IDAppointment %d, got %d", tc.findData.IDAppointment, data.IDAppointment)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockMedicalRecordRepo struct {
	getData  []model.MedicalRecord
	getErr   error
	findData model.MedicalRecord
	findErr  error
}

func (m *mockMedicalRecordRepo) Get() ([]model.MedicalRecord, error) { return m.getData, m.getErr }
func (m *mockMedicalRecordRepo) FindByID(_ int) (model.MedicalRecord, error) {
	return m.findData, m.findErr
}

func TestMedicalRecordUsecase_Get(t *testing.T) {
	sampleRecords := []model.MedicalRecord{
		{IDRecord: 1, IDAppointment: 1, IDICD: 76, HasilLab: "Normal", Tindakan: "Rawat Jalan"},
		{IDRecord: 2, IDAppointment: 2, IDICD: 309, HasilLab: "Gula darah tinggi", Tindakan: "Insulin"},
	}

	tests := []struct {
		name        string
		getData     []model.MedicalRecord
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success_multiple",
			getData:     sampleRecords,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "success get medical record",
		},
		{
			name:        "get_success_single",
			getData:     sampleRecords[:1],
			getErr:      nil,
			wantErr:     false,
			wantCount:   1,
			wantMessage: "success get medical record",
		},
		{
			name:        "get_empty",
			getData:     []model.MedicalRecord{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get medical record",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockMedicalRecordRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.MedicalRecordUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d records, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestMedicalRecordUsecase_GetByID(t *testing.T) {
	sample := model.MedicalRecord{
		IDRecord:      1,
		IDAppointment: 1,
		IDICD:         76,
		HasilLab:      "Leukosit normal",
		Tindakan:      "Pemberian antibiotik",
	}

	tests := []struct {
		name        string
		id          int
		findData    model.MedicalRecord
		findErr     error
		wantErr     bool
		wantMessage string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get medical record by ID",
		},
		{
			name:        "get_by_id_not_found",
			id:          99,
			findData:    model.MedicalRecord{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockMedicalRecordRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.MedicalRecordUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.IDRecord != tc.findData.IDRecord {
				t.Errorf("expected IDRecord %d, got %d", tc.findData.IDRecord, data.IDRecord)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockPrescriptionRepo struct {
	getData  []model.Prescription
	getErr   error
	findData model.Prescription
	findErr  error
}

func (m *mockPrescriptionRepo) Get() ([]model.Prescription, error) { return m.getData, m.getErr }
func (m *mockPrescriptionRepo) FindByID(_ int) (model.Prescription, error) {
	return m.findData, m.findErr
}

func TestPrescriptionUsecase_Get(t *testing.T) {
	samplePrescriptions := []model.Prescription{
		{IDResep: 1, IDRecord: 1, IDObat: 1, Jumlah: 10, AturanPakai: "3x sehari sesudah makan"},
		{IDResep: 2, IDRecord: 1, IDObat: 2, Jumlah: 5, AturanPakai: "1x sehari sebelum tidur"},
	}

	tests := []struct {
		name        string
		getData     []model.Prescription
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     samplePrescriptions,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "success get prescription",
		},
		{
			name:        "get_empty",
			getData:     []model.Prescription{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get prescription",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockPrescriptionRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.PrescriptionUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d prescriptions, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestPrescriptionUsecase_GetByID(t *testing.T) {
	sample := model.Prescription{IDResep: 1, IDRecord: 1, IDObat: 1, Jumlah: 10, AturanPakai: "3x sehari"}

	tests := []struct {
		name        string
		id          int
		findData    model.Prescription
		findErr     error
		wantErr     bool
		wantMessage string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get prescription by ID",
		},
		{
			name:        "get_by_id_not_found",
			id:          999,
			findData:    model.Prescription{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockPrescriptionRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.PrescriptionUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.IDResep != tc.findData.IDResep {
				t.Errorf("expected IDResep %d, got %d", tc.findData.IDResep, data.IDResep)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockPaymentRepo struct {
	getData  []model.Payment
	getErr   error
	findData model.Payment
	findErr  error
}

func (m *mockPaymentRepo) Get() ([]model.Payment, error)         { return m.getData, m.getErr }
func (m *mockPaymentRepo) FindByID(_ int) (model.Payment, error) { return m.findData, m.findErr }

func TestPaymentUsecase_Get(t *testing.T) {
	samplePayments := []model.Payment{
		{IDPayment: 1, IDAppointment: 1, Total: 150000, MetodePembayaran: "BPJS", StatusPembayaran: "Lunas"},
		{IDPayment: 2, IDAppointment: 2, Total: 250000, MetodePembayaran: "Tunai", StatusPembayaran: "Pending"},
	}

	tests := []struct {
		name        string
		getData     []model.Payment
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     samplePayments,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "success get payment",
		},
		{
			name:        "get_empty",
			getData:     []model.Payment{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get payment",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockPaymentRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.PaymentUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d payments, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestPaymentUsecase_GetByID(t *testing.T) {
	sample := model.Payment{IDPayment: 1, Total: 150000, MetodePembayaran: "BPJS", StatusPembayaran: "Lunas"}

	tests := []struct {
		name        string
		id          int
		findData    model.Payment
		findErr     error
		wantErr     bool
		wantMessage string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get payment by ID",
		},
		{
			name:        "get_by_id_not_found",
			id:          999,
			findData:    model.Payment{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockPaymentRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.PaymentUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.IDPayment != tc.findData.IDPayment {
				t.Errorf("expected IDPayment %d, got %d", tc.findData.IDPayment, data.IDPayment)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockICD10Repo struct {
	getData  []model.ICD10
	getErr   error
	findData model.ICD10
	findErr  error
}

func (m *mockICD10Repo) Get() ([]model.ICD10, error)         { return m.getData, m.getErr }
func (m *mockICD10Repo) FindByID(_ int) (model.ICD10, error) { return m.findData, m.findErr }

func TestICD10Usecase_Get(t *testing.T) {
	sampleICD10 := []model.ICD10{
		{IDICD: 1, KodeICD: "A01", NamaPenyakit: "Typhoid and paratyphoid fevers"},
		{IDICD: 76, KodeICD: "A90", NamaPenyakit: "Dengue fever [classical dengue]"},
		{IDICD: 309, KodeICD: "E11", NamaPenyakit: "Type 2 diabetes mellitus"},
	}

	tests := []struct {
		name        string
		getData     []model.ICD10
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     sampleICD10,
			getErr:      nil,
			wantErr:     false,
			wantCount:   3,
			wantMessage: "success get icd10",
		},
		{
			name:        "get_empty",
			getData:     []model.ICD10{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get icd10",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockICD10Repo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.ICD10Usecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d ICD10 entries, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestICD10Usecase_GetByID(t *testing.T) {
	sample := model.ICD10{IDICD: 309, KodeICD: "E11", NamaPenyakit: "Type 2 diabetes mellitus"}

	tests := []struct {
		name        string
		id          int
		findData    model.ICD10
		findErr     error
		wantErr     bool
		wantMessage string
		wantKode    string
	}{
		{
			name:        "get_by_id_success",
			id:          309,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get icd10 by ID",
			wantKode:    "E11",
		},
		{
			name:        "get_by_id_not_found",
			id:          9999,
			findData:    model.ICD10{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
			wantKode:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockICD10Repo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.ICD10Usecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.KodeICD != tc.wantKode {
				t.Errorf("expected KodeICD %q, got %q", tc.wantKode, data.KodeICD)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockDiseaseMonitoringRepo struct {
	getData  []model.DiseaseMonitoring
	getErr   error
	findData model.DiseaseMonitoring
	findErr  error
}

func (m *mockDiseaseMonitoringRepo) Get() ([]model.DiseaseMonitoring, error) {
	return m.getData, m.getErr
}
func (m *mockDiseaseMonitoringRepo) FindByID(_ int) (model.DiseaseMonitoring, error) {
	return m.findData, m.findErr
}

func TestDiseaseMonitoringUsecase_Get(t *testing.T) {
	now := time.Now()
	sampleMonitoring := []model.DiseaseMonitoring{
		{IDMonitoring: 1, IDICD: 76, Negara: "Indonesia", TotalKasus: 5000, TotalKematian: 50, TotalSembuh: 4800, TanggalUpdate: now},
		{IDMonitoring: 2, IDICD: 309, Negara: "Malaysia", TotalKasus: 2000, TotalKematian: 20, TotalSembuh: 1900, TanggalUpdate: now},
	}

	tests := []struct {
		name        string
		getData     []model.DiseaseMonitoring
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     sampleMonitoring,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "success get disease monitoring",
		},
		{
			name:        "get_empty",
			getData:     []model.DiseaseMonitoring{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get disease monitoring",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockDiseaseMonitoringRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.DiseaseMonitoringUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d monitoring records, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestDiseaseMonitoringUsecase_GetByID(t *testing.T) {
	sample := model.DiseaseMonitoring{
		IDMonitoring: 1, IDICD: 76, Negara: "Indonesia",
		TotalKasus: 5000, TotalKematian: 50, TotalSembuh: 4800,
	}

	tests := []struct {
		name        string
		id          int
		findData    model.DiseaseMonitoring
		findErr     error
		wantErr     bool
		wantMessage string
		wantNegara  string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get disease monitoring by ID",
			wantNegara:  "Indonesia",
		},
		{
			name:        "get_by_id_not_found",
			id:          999,
			findData:    model.DiseaseMonitoring{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
			wantNegara:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockDiseaseMonitoringRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.DiseaseMonitoringUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.Negara != tc.wantNegara {
				t.Errorf("expected Negara %q, got %q", tc.wantNegara, data.Negara)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

type mockHealthNewsRepo struct {
	getData  []model.HealthNews
	getErr   error
	findData model.HealthNews
	findErr  error
}

func (m *mockHealthNewsRepo) Get() ([]model.HealthNews, error)         { return m.getData, m.getErr }
func (m *mockHealthNewsRepo) FindByID(_ int) (model.HealthNews, error) { return m.findData, m.findErr }

func TestHealthNewsUsecase_Get(t *testing.T) {
	now := time.Now()
	sampleNews := []model.HealthNews{
		{IDNews: 1, Judul: "Waspada DBD Meningkat", Sumber: "Kemenkes", Kategori: "Penyakit Menular", TanggalPublish: now, URL: "https://example.com/dbd"},
		{IDNews: 2, Judul: "Manfaat Olahraga Pagi", Sumber: "WHO", Kategori: "Gaya Hidup Sehat", TanggalPublish: now, URL: "https://example.com/olahraga"},
	}

	tests := []struct {
		name        string
		getData     []model.HealthNews
		getErr      error
		wantErr     bool
		wantCount   int
		wantMessage string
	}{
		{
			name:        "get_success",
			getData:     sampleNews,
			getErr:      nil,
			wantErr:     false,
			wantCount:   2,
			wantMessage: "success get health news",
		},
		{
			name:        "get_empty",
			getData:     []model.HealthNews{},
			getErr:      nil,
			wantErr:     false,
			wantCount:   0,
			wantMessage: "success get health news",
		},
		{
			name:        "get_db_error",
			getData:     nil,
			getErr:      errDB,
			wantErr:     true,
			wantCount:   0,
			wantMessage: "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockHealthNewsRepo{getData: tc.getData, getErr: tc.getErr}
			uc := usecase.HealthNewsUsecase(repo)

			data, message, detail, err := uc.Get()

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(data) != tc.wantCount {
				t.Errorf("expected %d news items, got %d", tc.wantCount, len(data))
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}

func TestHealthNewsUsecase_GetByID(t *testing.T) {
	sample := model.HealthNews{
		IDNews:   1,
		Judul:    "Waspada DBD Meningkat",
		Sumber:   "Kemenkes",
		Kategori: "Penyakit Menular",
		URL:      "https://example.com/dbd",
	}

	tests := []struct {
		name        string
		id          int
		findData    model.HealthNews
		findErr     error
		wantErr     bool
		wantMessage string
		wantJudul   string
	}{
		{
			name:        "get_by_id_success",
			id:          1,
			findData:    sample,
			findErr:     nil,
			wantErr:     false,
			wantMessage: "success get health news by ID",
			wantJudul:   "Waspada DBD Meningkat",
		},
		{
			name:        "get_by_id_not_found",
			id:          500,
			findData:    model.HealthNews{},
			findErr:     errNotFound,
			wantErr:     true,
			wantMessage: "data not found",
			wantJudul:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := &mockHealthNewsRepo{findData: tc.findData, findErr: tc.findErr}
			uc := usecase.HealthNewsUsecase(repo)

			data, message, detail, err := uc.GetByID(tc.id)

			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if message != tc.wantMessage {
				t.Errorf("expected message %q, got %q", tc.wantMessage, message)
			}
			if !tc.wantErr && data.Judul != tc.wantJudul {
				t.Errorf("expected Judul %q, got %q", tc.wantJudul, data.Judul)
			}
			if tc.wantErr && detail == "" {
				t.Errorf("expected non-empty detail on error")
			}
		})
	}
}
