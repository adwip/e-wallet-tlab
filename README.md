# E-Wallet T-Lab

A RESTful API service for digital wallet management built with Go, Echo framework, and MySQL.

##  Prerequisites

- Go 1.25.0 or higher
- MySQL 8.0 or higher
- Make (optional, for using Makefile commands)

## 🛠️ Installation & Setup

### 1. Clone the repository
```bash
git clone <repository-url>
cd e-wallet-tlab
```

### 2. Install dependencies
```bash
go mod download
go mod vendor
```

### 3. Database Setup

Create a MySQL database:
```sql
CREATE DATABASE tlab_wallet_db;
```

### 4. Configuration

**Option A: Using config.json (Recommended)**

Create `internal/shared/config/config.json`:
```json
{
    "service": {
        "name": "T-Lab Wallet",
        "port": ":8082",
        "log_file": "log/tlab-wallet.log"
    },
    "db": {
        "host": "root:password@tcp(localhost:3306)/tlab_wallet_db?parseTime=true"
    }
}
```

**Option B: Using .env file**

Copy `.env.example` to `.env` and configure:
```env
SERVICE_NAME=T-Lab Wallet
SERVICE_PORT=:8082
SERVICE_LOG_FILE=log/tlab-wallet.log
DB_HOST=root:password@tcp(localhost:3306)/tlab_wallet_db?parseTime=true
```

### 5. Run Database Migration
```bash
make migrate
```

## 🏃 Running the Service

### Linux/Mac
```bash
make run
```

### Windows
```bash
make windows-run
```

The service will start on `http://localhost:8082`

## � Available Commands

```bash
make migrate       # Run database migrations
make run          # Run service (Linux/Mac)
make windows-run  # Run service (Windows)
```

## 📚 API Endpoints

The service exposes the following endpoints at `http://localhost:8082/api/v1`:

**Public Endpoints:**
- `POST /auth/register` - User registration
- `POST /auth/login` - User login

**Protected Endpoints (require JWT token):**
- `GET /users/profile` - Get user profile
- `GET /wallets/balance` - Get wallet balance
- `POST /wallets/topup` - Top up wallet
- `POST /transactions/transfer` - Transfer to another wallet
- `GET /transactions/history` - Get transaction history

For detailed API documentation, see `api_endpoints.md`

## 🐛 Troubleshooting

**Database Connection Error:**
- Verify MySQL is running
- Check database credentials in config
- Ensure database exists

**Port Already in Use:**
- Change port in config.json
- Kill process using the port

**Migration Errors:**
- Ensure database is accessible
- Check database user permissions
