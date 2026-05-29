# Healthcare Management System - API Documentation

**Version:** 1.0  
**Last Updated:** May 29, 2026  
**Base URL:** `http://localhost:12100`  
**Framework:** Echo v4 (Go)  
**Database:** PostgreSQL (Railway)  
**Cache:** Redis (Railway)  

---

## Table of Contents

1. [Overview](#overview)
2. [Authentication](#authentication)
3. [Environment Variables](#environment-variables)
4. [API Endpoints](#api-endpoints)
5. [Response Format](#response-format)
6. [Error Handling](#error-handling)
7. [Rate Limiting & Performance](#rate-limiting--performance)
8. [Data Models](#data-models)
9. [Testing](#testing)
10. [Deployment](#deployment)

---

## Overview

The Healthcare Management System API is a comprehensive RESTful API built with Go and Echo framework. It provides endpoints for managing:

- **Patients** - Patient information and records
- **Appointments** - Doctor-patient appointment scheduling
- **Medical Records** - Clinical notes and patient medical history
- **Medicines** - Pharmacy inventory management
- **Prescriptions** - Medicine prescriptions from medical records
- **Payments** - Appointment payment tracking
- **ICD-10 Codes** - Disease classification codes
- **Disease Monitoring** - Global disease statistics and tracking
- **Health News** - Medical/health news feed

### Architecture

The system follows a **4-Layer Domain-Driven Design** pattern:

```
HTTP Request
    ↓
Handler Layer (HTTP Controllers)
    ↓
Usecase Layer (Business Logic)
    ↓
Repository Layer (Data Access)
    ↓
Database / Redis
```

### Technology Stack

| Component | Technology | Version |
|-----------|-----------|---------|
| Language | Go | 1.25.0 |
| Web Framework | Echo | v4.15.2 |
| Database | PostgreSQL | Latest |
| ORM | GORM | v1.31.1 |
| Cache/Streaming | Redis | v8 |
| Logging | Logrus | v1.9.4 |
| API Documentation | Swagger/OpenAPI | v2 |

---

## Authentication

All API endpoints require **API Key authentication** via the `api-key` header.

### Header Format

```
api-key: 760af993-42f4-452b-9242-20f05cf41871
```

### Authentication Flow

1. **Request:** Include `api-key` header in every request
2. **Validation:** Middleware validates the API key
3. **Response:** 
   - ✅ **200 OK** - Valid key, request processed
   - ❌ **403 Forbidden** - Missing or invalid key

### Without Authentication

```bash
curl -X GET http://localhost:12100/api/medicine
```

**Response (403 Forbidden):**
```json
{
  "message": "invalid api key"
}
```

### With Authentication

```bash
curl -X GET http://localhost:12100/api/medicine \
  -H "api-key: 760af993-42f4-452b-9242-20f05cf41871"
```

---

## Environment Variables

### Core Service (.env)

Located in: `core/.env`

| Variable | Example | Description |
|----------|---------|-------------|
| **PORT** | 12100 | Server listening port |
| **DB_NAME** | railway | PostgreSQL database name |
| **DB_USER** | postgres | Database username |
| **DB_HOST** | zephyr.proxy.rlwy.net | Database host (Railway) |
| **DB_PORT** | 31011 | Database port |
| **DB_PASSWORD** | EAXhvK... | Database password |
| **REDIS_HOST** | ballast.proxy.rlwy.net | Redis host (Railway) |
| **REDIS_PORT** | 30050 | Redis port |
| **REDIS_PASSWORD** | wSTwAv... | Redis password |
| **API_KEY** | 760af993-... | API authentication key |
| **SECRET** | 0143ee2... | JWT/encryption secret |
| **DB_SLAVE_NAME** | railway | Read replica database name |
| **DB_SLAVE_USER** | postgres | Read replica username |
| **DB_SLAVE_HOST** | viaduct.proxy.rlwy.net | Read replica host |
| **DB_SLAVE_PORT** | 12136 | Read replica port |
| **DB_SLAVE_PASSWORD** | BDAQvV... | Read replica password |
| **DB_SLAVE_SCHEMA** | db_slave_healthcare | Read replica schema |

### Gateway Service (.env)

Located in: `gateway/.env`

| Variable | Example | Description |
|----------|---------|-------------|
| **PORT** | 3000 | Gateway server port |
| **REDIS_HOST** | ballast.proxy.rlwy.net | Redis host |
| **REDIS_PORT** | 30050 | Redis port |
| **REDIS_PASSWORD** | wSTwAv... | Redis password |

---

## API Endpoints

### 1. Medicines Management

#### Get All Medicines

```http
GET /api/medicine
```

**Headers:**
```
api-key: 760af993-42f4-452b-9242-20f05cf41871
Content-Type: application/json
```

**Response (200 OK):**
```json
{
  "message": "succes get medicine",
  "data": [
    {
      "id_obat": 1,
      "nama_obat": "paracetamol",
      "kategori": "pil",
      "stok": 100,
      "harga": 2000,
      "expired_date": "2030-12-12T00:00:00Z"
    }
  ]
}
```

#### Get Medicine by ID

```http
GET /api/medicine/:id
```

**Example:**
```http
GET /api/medicine/1
```

**Path Parameters:**
| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Medicine ID |

**Response (200 OK):**
```json
{
  "message": "succes get medicine",
  "data": {
    "id_obat": 1,
    "nama_obat": "paracetamol",
    "kategori": "pil",
    "stok": 100,
    "harga": 2000,
    "expired_date": "2030-12-12T00:00:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "message": "Medicine not found",
  "detail": "record not found"
}
```

---

### 2. Appointments Management

#### Get All Appointments

```http
GET /api/appointment
```

**Response (200 OK):**
```json
{
  "message": "success get appointment",
  "data": [
    {
      "id_appointment": 2,
      "id_patient": 1,
      "id_doctor": 4,
      "tanggal": "2026-05-29T00:00:00+07:00",
      "keluhan": "Demam dan batuk",
      "tekanan_darah": "120/80",
      "suhu_tubuh": 37.5,
      "berat_badan": 65.5,
      "status": "PENDING",
      "patient": {
        "id_patient": 1,
        "id_status": 0,
        "nama": "TEST",
        "nik": "",
        "tanggal_lahir": "0001-01-01T00:00:00Z",
        "gender": "",
        "alamat": "",
        "no_hp": "",
        "golongan_darah": "",
        "status": {
          "id_status": 0,
          "nama_status": ""
        }
      },
      "doctor": {
        "id_user": 4,
        "id_role": 1,
        "id_cabang": 1,
        "nama": "Dr. Budi",
        "email": "budi@test.com",
        "password": "123456",
        "no_hp": "08123456789",
        "role": {
          "id_role": 0,
          "nama_role": ""
        },
        "branch": {
          "id_cabang": 0,
          "nama_cabang": "",
          "alamat": ""
        }
      }
    }
  ]
}
```

#### Get Appointment by ID

```http
GET /api/appointment/:id
```

**Example:**
```http
GET /api/appointment/2
```

---

### 3. Medical Records Management

#### Get All Medical Records

```http
GET /api/medical-record
```

**Response (200 OK):**
```json
{
  "message": "success get medical record",
  "data": [
    {
      "id_record": 1,
      "id_appointment": 2,
      "id_icd": 1,
      "hasil_lab": "Normal",
      "hasil_radiologi": "Tidak ada kelainan",
      "tindakan": "Pemberian obat",
      "catatan": "Istirahat cukup",
      "appointment": {
        "id_appointment": 2,
        "id_patient": 1,
        "id_doctor": 4,
        "tanggal": "2026-05-29T00:00:00+07:00",
        "keluhan": "Demam dan batuk",
        "tekanan_darah": "120/80",
        "suhu_tubuh": 37.5,
        "berat_badan": 65.5,
        "status": "PENDING"
      },
      "icd10": {
        "id_icd": 1,
        "kode_icd": "A01",
        "nama_penyakit": "Typhoid and paratyphoid fevers"
      }
    }
  ]
}
```

#### Get Medical Record by ID

```http
GET /api/medical-record/:id
```

---

### 4. Patients Management

#### Get All Patients

```http
GET /api/patient
```

**Response (200 OK):**
```json
{
  "message": "success get patient",
  "data": [
    {
      "id_patient": 1,
      "id_status": 1,
      "nama": "John Doe",
      "nik": "1234567890123456",
      "tanggal_lahir": "1990-05-15T00:00:00Z",
      "gender": "M",
      "alamat": "Jl. Merdeka No. 123",
      "no_hp": "08123456789",
      "golongan_darah": "O",
      "status": {
        "id_status": 1,
        "nama_status": "Active"
      }
    }
  ]
}
```

#### Get Patient by ID

```http
GET /api/patient/:id
```

---

### 5. Prescriptions Management

#### Get All Prescriptions

```http
GET /api/prescription
```

**Response (200 OK):**
```json
{
  "message": "success get prescription",
  "data": [
    {
      "id_resep": 1,
      "id_record": 1,
      "id_obat": 1,
      "jumlah": 10,
      "aturan_pakai": "3 x 1 tablet setelah makan",
      "medical_record": { ... },
      "medicine": { ... }
    }
  ]
}
```

#### Get Prescription by ID

```http
GET /api/prescription/:id
```

---

### 6. Payments Management

#### Get All Payments

```http
GET /api/payment
```

**Response (200 OK):**
```json
{
  "message": "success get payment",
  "data": [
    {
      "id_payment": 1,
      "id_appointment": 2,
      "total": 150000,
      "metode_pembayaran": "credit_card",
      "status_pembayaran": "PAID",
      "appointment": { ... }
    }
  ]
}
```

#### Get Payment by ID

```http
GET /api/payment/:id
```

---

### 7. ICD-10 Codes Management

#### Get All ICD-10 Codes

```http
GET /api/icd10
```

**Response (200 OK):**
```json
{
  "message": "success get icd10",
  "data": [
    {
      "id_icd": 1,
      "kode_icd": "A01",
      "nama_penyakit": "Typhoid and paratyphoid fevers"
    },
    {
      "id_icd": 2,
      "kode_icd": "A01.0",
      "nama_penyakit": "Typhoid fever"
    }
  ]
}
```

#### Get ICD-10 Code by ID

```http
GET /api/icd10/:id
```

---

### 8. Disease Monitoring

#### Get All Disease Monitoring Data

```http
GET /api/disease-monitoring
```

**Response (200 OK):**
```json
{
  "message": "success get disease monitoring",
  "data": [
    {
      "id_monitoring": 1,
      "id_icd": 1,
      "negara": "Indonesia",
      "total_kasus": 1500,
      "total_kematian": 75,
      "total_sembuh": 1200,
      "tanggal_update": "2026-05-29T10:00:00Z",
      "icd10": {
        "id_icd": 1,
        "kode_icd": "A01",
        "nama_penyakit": "Typhoid and paratyphoid fevers"
      }
    }
  ]
}
```

#### Get Disease Monitoring by ID

```http
GET /api/disease-monitoring/:id
```

---

### 9. Health News

#### Get All Health News

```http
GET /api/health-news
```

**Response (200 OK):**
```json
{
  "message": "success get health news",
  "data": [
    {
      "id_news": 1,
      "judul": "Tips Pencegahan COVID-19",
      "sumber": "Ministry of Health",
      "kategori": "Prevention",
      "tanggal_publish": "2026-05-28T15:30:00Z",
      "url": "https://example.com/news/1"
    }
  ]
}
```

#### Get Health News by ID

```http
GET /api/health-news/:id
```

---

## Response Format

### Success Response (2xx)

All successful responses follow this format:

```json
{
  "message": "success get [entity]",
  "data": [ ... ]  // or single object for GetByID
}
```

### Error Response (4xx, 5xx)

All error responses follow this format:

```json
{
  "message": "error message",
  "detail": "detailed error information"
}
```

---

## Error Handling

### HTTP Status Codes

| Status Code | Meaning | Example Scenario |
|------------|---------|-----------------|
| 200 | OK | Successful GET request |
| 201 | Created | Successful POST request (if implemented) |
| 400 | Bad Request | Invalid ID parameter (non-integer) |
| 403 | Forbidden | Missing or invalid API key |
| 404 | Not Found | Resource ID does not exist |
| 500 | Internal Server Error | Database query failure |

### Error Examples

#### 403 - Missing API Key

```bash
curl -X GET http://localhost:12100/api/medicine
```

**Response:**
```json
{
  "message": "nil token"
}
```

#### 403 - Invalid API Key

```bash
curl -X GET http://localhost:12100/api/medicine \
  -H "api-key: invalid-key-12345"
```

**Response:**
```json
{
  "message": "invalid api key"
}
```

#### 404 - Resource Not Found

```bash
curl -X GET http://localhost:12100/api/medicine/999 \
  -H "api-key: 760af993-42f4-452b-9242-20f05cf41871"
```

**Response:**
```json
{
  "message": "medicine not found",
  "detail": "record not found"
}
```

#### 400 - Invalid Parameter

```bash
curl -X GET http://localhost:12100/api/medicine/abc \
  -H "api-key: 760af993-42f4-452b-9242-20f05cf41871"
```

**Response:**
```json
{
  "message": "invalid id",
  "detail": "strconv.Atoi: parsing \"abc\": invalid syntax"
}
```

#### 500 - Database Error

```json
{
  "message": "error message",
  "detail": "database connection failed"
}
```

---

## Rate Limiting & Performance

### Current Implementation

- **No explicit rate limiting** implemented
- **Recommendations:**
  1. Implement token bucket algorithm for rate limiting
  2. Cache frequently accessed data in Redis
  3. Implement pagination for list endpoints
  4. Add response timeout configurations

### Performance Tips

1. **Use GetByID endpoints** when only one record is needed
2. **Filter at database level** when possible
3. **Monitor Redis consumer** for event processing delays
4. **Check PostgreSQL logs** for slow queries

---

## Data Models

### Medicine Model
```
{
  id_obat:       integer (PK)
  nama_obat:     string(150)
  kategori:      string(100)
  stok:          integer
  harga:         decimal(12,2)
  expired_date:  date
}
```

### Appointment Model
```
{
  id_appointment:  integer (PK)
  id_patient:      integer (FK)
  id_doctor:       integer (FK)
  tanggal:         timestamp
  keluhan:         text
  tekanan_darah:   string(20)
  suhu_tubuh:      decimal(4,1)
  berat_badan:     decimal(5,2)
  status:          string(50)
}
```

### Medical Record Model
```
{
  id_record:        integer (PK)
  id_appointment:   integer (FK)
  id_icd:           integer (FK)
  hasil_lab:        text
  hasil_radiologi:  text
  tindakan:         text
  catatan:          text
}
```

### Patient Model
```
{
  id_patient:      integer (PK)
  id_status:       integer (FK)
  nama:            string(150)
  nik:             string(30) UNIQUE
  tanggal_lahir:   date
  gender:          string(20)
  alamat:          text
  no_hp:           string(20)
  golongan_darah:  string(5)
}
```

### Prescription Model
```
{
  id_resep:       integer (PK)
  id_record:      integer (FK)
  id_obat:        integer (FK)
  jumlah:         integer
  aturan_pakai:   text
}
```

### Payment Model
```
{
  id_payment:           integer (PK)
  id_appointment:       integer (FK)
  total:                decimal(12,2)
  metode_pembayaran:    string(50)
  status_pembayaran:    string(50)
}
```

### ICD-10 Model
```
{
  id_icd:        integer (PK)
  kode_icd:      string(20) UNIQUE
  nama_penyakit: string(200)
}
```

### Disease Monitoring Model
```
{
  id_monitoring:   integer (PK)
  id_icd:          integer (FK)
  negara:          string(100)
  total_kasus:     integer
  total_kematian:  integer
  total_sembuh:    integer
  tanggal_update:  timestamp
}
```

### Health News Model
```
{
  id_news:          integer (PK)
  judul:            string(255)
  sumber:           string(150)
  kategori:         string(100)
  tanggal_publish:  timestamp
  url:              text
}
```

---

## Testing

### Using cURL

```bash
# Test with API key
curl -X GET http://localhost:12100/api/medicine \
  -H "api-key: 760af993-42f4-452b-9242-20f05cf41871"

# Test by ID
curl -X GET http://localhost:12100/api/medicine/1 \
  -H "api-key: 760af993-42f4-452b-9242-20f05cf41871"
```

### Using PowerShell

```powershell
$apiKey = "760af993-42f4-452b-9242-20f05cf41871"
$headers = @{ "api-key" = $apiKey }

# Get all medicines
$response = Invoke-WebRequest -Uri "http://localhost:12100/api/medicine" `
  -Headers $headers -Method Get
$response.Content | ConvertFrom-Json

# Get specific medicine
$response = Invoke-WebRequest -Uri "http://localhost:12100/api/medicine/1" `
  -Headers $headers -Method Get
$response.Content | ConvertFrom-Json
```

### Using Postman

1. Create new collection: "Healthcare API"
2. Add Authorization header to collection:
   - Type: API Key
   - Key: `api-key`
   - Value: `760af993-42f4-452b-9242-20f05cf41871`
3. Create requests for each endpoint
4. Test with environment variables for base URL

### Using Python

```python
import requests
import json

api_key = "760af993-42f4-452b-9242-20f05cf41871"
base_url = "http://localhost:12100"
headers = {"api-key": api_key}

# Get all medicines
response = requests.get(f"{base_url}/api/medicine", headers=headers)
print(json.dumps(response.json(), indent=2))

# Get medicine by ID
response = requests.get(f"{base_url}/api/medicine/1", headers=headers)
print(json.dumps(response.json(), indent=2))
```

---

## Deployment

### Prerequisites

- Go 1.25.0
- PostgreSQL 18.4+ (Railway)
- Redis 8+ (Railway)
- Environment variables configured in `.env`

### Building

```bash
cd core
go mod tidy
go build -o healthcare-core main.go
```

### Running

```bash
# Development
cd core
go run .

# Production
./healthcare-core
```

### Docker Deployment (Future)

```dockerfile
FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o healthcare main.go
EXPOSE 12100
CMD ["./healthcare"]
```

### Environment Setup for Deployment

Create `.env` file with:
```env
PORT=12100
DB_NAME=railway
DB_USER=postgres
DB_PASSWORD=xxxxx
DB_HOST=zephyr.proxy.rlwy.net
DB_PORT=31011
REDIS_HOST=ballast.proxy.rlwy.net
REDIS_PORT=30050
REDIS_PASSWORD=xxxxx
API_KEY=760af993-42f4-452b-9242-20f05cf41871
SECRET=xxxxx
```

---

## API Endpoint Summary

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | /api/medicine | Get all medicines | ✅ |
| GET | /api/medicine/:id | Get medicine by ID | ✅ |
| GET | /api/appointment | Get all appointments | ✅ |
| GET | /api/appointment/:id | Get appointment by ID | ✅ |
| GET | /api/patient | Get all patients | ✅ |
| GET | /api/patient/:id | Get patient by ID | ✅ |
| GET | /api/medical-record | Get all medical records | ✅ |
| GET | /api/medical-record/:id | Get medical record by ID | ✅ |
| GET | /api/prescription | Get all prescriptions | ✅ |
| GET | /api/prescription/:id | Get prescription by ID | ✅ |
| GET | /api/payment | Get all payments | ✅ |
| GET | /api/payment/:id | Get payment by ID | ✅ |
| GET | /api/icd10 | Get all ICD-10 codes | ✅ |
| GET | /api/icd10/:id | Get ICD-10 code by ID | ✅ |
| GET | /api/disease-monitoring | Get all disease monitoring data | ✅ |
| GET | /api/disease-monitoring/:id | Get disease monitoring by ID | ✅ |
| GET | /api/health-news | Get all health news | ✅ |
| GET | /api/health-news/:id | Get health news by ID | ✅ |

---

## Swagger/OpenAPI Documentation

Access the Swagger UI at:
```
http://localhost:12100/swagger/
```

Swagger JSON schema available at:
```
http://localhost:12100/swagger/doc.json
```

---

## Support & Troubleshooting

### Common Issues

1. **"invalid api key" error**
   - Verify API key in `.env` file
   - Check header name: must be `api-key` (lowercase)

2. **"Connection refused" error**
   - Ensure core service is running: `go run .`
   - Verify port 12100 is available: `netstat -ano | findstr :12100`

3. **"Database connection failed"**
   - Verify PostgreSQL credentials in `.env`
   - Check Railway database connection status
   - Ensure network access to Railway proxy

4. **"record not found"**
   - Verify the ID exists in database
   - Check database has been seeded with data

### Logging

Check application logs:
```bash
# View recent logs
tail -f core-health-2026-05-29.log

# Search for errors
grep -i error core-health-2026-05-29.log
```

### Redis Streaming

The system uses Redis streams for event-driven architecture:
- **Stream:** `HEALTHCARE:STREAM`
- **Consumer Group:** `HEALTHCARE:GROUP`
- **Consumer:** `HEALTHCARE:CONSUMER`

Monitor Redis consumer:
```bash
redis-cli
> XINFO STREAM HEALTHCARE:STREAM
> XINFO GROUPS HEALTHCARE:STREAM
```

---

**Document Generated:** May 29, 2026  
**API Server Status:** ✅ Running on port 12100  
**Database:** ✅ Connected to Railway PostgreSQL  
**Redis Cache:** ✅ Connected to Railway Redis
