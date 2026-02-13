# 💝 Valentine Proposal Site - Backend

A Go-based backend API for creating and managing personalized Valentine's Day proposal links. This service allows users to generate unique proposal URLs with custom messages that can be shared with their loved ones.

## 🚀 Features

- **Generate Unique Proposal Links**: Create personalized proposal URLs with custom sender, receiver, and message details
- **Retrieve Proposal Data**: Fetch proposal information using unique short codes
- **In-Memory Storage**: Fast, lightweight data storage using an in-memory store
- **CORS Enabled**: Configured to work seamlessly with frontend applications
- **RESTful API**: Clean and simple API endpoints

## 🛠️ Tech Stack

- **Language**: Go 1.25.5
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) v1.11.0
- **CORS Middleware**: gin-contrib/cors v1.7.6
- **Architecture**: Clean architecture with separation of concerns (handlers, services, repositories, models)

## 📁 Project Structure

```
VALENTINE-PROPOSAL-SITE-BACKEND/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internals/
│   ├── config/
│   │   └── database.go          # Database configuration
│   ├── handler/
│   │   └── url_handler.go       # HTTP request handlers
│   ├── models/
│   │   ├── proposal.go          # Proposal data model
│   │   └── proposal_dto.go      # Data transfer object
│   ├── service/
│   │   └── proposal_service.go  # Business logic layer
│   ├── store/
│   │   ├── memory_store.go      # In-memory storage implementation
│   │   └── proposal_repo.go     # Repository interface
│   └── util/
│       └── utils.go             # Utility functions
├── go.mod                        # Go module dependencies
└── go.sum                        # Dependency checksums
```

## 📋 Prerequisites

- [Go](https://golang.org/dl/) 1.25.5 or higher
- Git (for cloning the repository)

## ⚙️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Dhruvdangi03/VALENTINE-PROPOSAL-SITE-BACKEND.git
   cd VALENTINE-PROPOSAL-SITE-BACKEND
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the server**
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8080`

## 🔌 API Endpoints

### 1. Create Proposal Link

**Endpoint**: `POST /`

**Request Body**:
```json
{
  "Email": "sender@example.com",
  "Sender": "John",
  "Receiver": "Jane",
  "Message": "Will you be my Valentine?"
}
```

**Response**:
```json
{
  "link": "http://localhost:8080/abc123"
}
```

### 2. Get Proposal Data

**Endpoint**: `GET /:code`

**Example**: `GET /abc123`

**Response**:
```json
{
  "data": {
    "Email": "sender@example.com",
    "Sender": "John",
    "Receiver": "Jane",
    "Message": "Will you be my Valentine?",
    "CreatedTime": "2026-02-13T10:30:00Z",
    "ExpiryTime": "2026-02-14T23:59:59Z",
    "Count": 1
  }
}
```

**Error Response** (404):
```json
{
  "error": "short URL not found"
}
```

## 🌐 CORS Configuration

The API is configured to accept requests from:
- `http://localhost:5173` (default Vite development server)

To modify CORS settings, edit the configuration in `cmd/server/main.go`:

```go
r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "OPTIONS"},
    AllowHeaders:     []string{"Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
}))
```

## 🏗️ Architecture

This project follows **Clean Architecture** principles:

- **Handler Layer**: Handles HTTP requests and responses
- **Service Layer**: Contains business logic
- **Repository Layer**: Manages data persistence
- **Models**: Defines data structures

## 🧪 Development

### Build the application
```bash
go build -o valentine-backend cmd/server/main.go
```

### Run the built binary
```bash
./valentine-backend
```

## 📝 Data Model

### Proposal
```go
type Proposal struct {
    ID          uint      // Unique identifier
    Email       string    // Sender's email
    Sender      string    // Sender's name
    Receiver    string    // Receiver's name
    Message     string    // Proposal message
    CreatedTime time.Time // Creation timestamp
    ExpiryTime  time.Time // Expiration timestamp
    Count       int64     // View count
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 👨‍💻 Author

**Dhruv Dangi** - [GitHub](https://github.com/Dhruvdangi03)

## 💡 Related Projects

This backend is designed to work with the Valentine Proposal Site frontend. Make sure both services are running for the complete experience.

---

Made with ❤️ for Valentine's Day
