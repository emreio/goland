# Goland Repository Instructions

## Overview
**Repository**: `emreio/goland`  
**Description**: My GoLang Learnings  
**Language Composition**: Go (98.4%), Dockerfile (1.6%)  
**Created**: January 3, 2023  
**Last Updated**: September 13, 2023

---

## Project Structure

### Root Directory
- **main.go** - Main HTTP server application with API handlers
- **go.mod** - Go module definition (module: `mylearnings.go/main`, Go 1.19)
- **go.sum** - Go dependencies lock file
- **Dockerfile** - Docker containerization (Go 1.19 base image)
- **data.dat** - Data storage file
- **.gitignore** - Git ignore rules

### Subdirectories
- **InModuleLib/** - Custom HTTP server module library
- **business/** - Business logic package (contains `GetUsers()` and `GetRandomNumber()` functions)
- **play/** - Experimental/learning code for Go fundamentals:
  - User struct manipulation
  - Pointer and value receiver methods
  - Environment variables
  - String encoding and byte handling
- **sandbox/** - Session management learning:
  - Session manager implementation
  - Authentication pattern
  - Concurrency with sync.Mutex
  - HTTP cookie handling

---

## Core Features

### HTTP Server (main.go)
The application runs a custom HTTP server with the following endpoints:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/get` | GET | Basic GET handler returning "hello go" |
| `/post` | POST | POST handler returning confirmation message |
| `/api/swagger` | GET | API endpoint returning JSON response |
| `/api/query` | GET | Query parameter handler |
| `/getUsers` | GET | Calls business package to retrieve users |
| `/random` | GET | Calls business package to get random number |
| `/heartbeat` | GET | Server health check with connection count |

**Server Configuration**:
- Port: 2222
- Version: 1
- Connection counter tracking total requests

### Key Components
1. **Custom HTTP Server** (`InModuleLib`): Handles routing and request processing
2. **Authorization Middleware**: Header-based authorization checking
3. **JSON Serialization**: Structured response objects
4. **Error Handling**: Basic error responses and status codes

---

## Docker Support

The project includes a Dockerfile for containerization:
- **Base Image**: golang:1.19
- **Build Process**: 
  - Downloads and verifies Go modules
  - Copies source code
  - Builds binary to `/usr/local/bin/app`
  - Exposes default entrypoint: `app` command

**Build & Run**:
```bash
docker build -t goland:latest .
docker run -p 2222:2222 goland:latest
```

---

## Getting Started

### Prerequisites
- Go 1.19 or higher
- Docker (optional, for containerized deployment)

### Installation & Running

1. **Clone the repository**:
   ```bash
   git clone https://github.com/emreio/goland.git
   cd goland
   ```

2. **Download dependencies**:
   ```bash
   go mod download
   go mod verify
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```

4. **Test endpoints**:
   ```bash
   # Heartbeat
   curl http://localhost:2222/heartbeat
   
   # GET API
   curl http://localhost:2222/api/swagger
   
   # Basic handler
   curl http://localhost:2222/get
   ```

---

## Development Guidelines

### Code Organization
- **Main Application**: Root level `main.go` - HTTP server and routing
- **Libraries**: `InModuleLib/` - Reusable HTTP server components
- **Business Logic**: `business/` - Core application functionality
- **Learning/Experiments**: `play/` and `sandbox/` - Educational code

### Dependencies
- **golang.org/x/text v0.10.0** - Text handling and encoding

### Testing & Validation
- Connection count tracking for monitoring
- Authorization header checking for security patterns
- JSON marshaling for API responses
- Error handling with appropriate HTTP status codes

---

## Key Learning Areas Covered

This repository demonstrates learning in the following Go concepts:

1. **HTTP Server Development**
   - Custom HTTP handlers
   - Routing and middleware
   - JSON serialization

2. **Object-Oriented Patterns in Go**
   - Structs and methods
   - Value vs. pointer receivers
   - Interface implementations

3. **Concurrency & Threading**
   - Sync primitives (Mutex)
   - Session management

4. **Data Handling**
   - Environment variables
   - String and byte manipulation
   - File operations

5. **Containerization**
   - Docker multi-stage builds
   - Go application deployment

---

## Contributing Guidelines

As a personal learning repository, contributions follow these practices:

1. **Branch Naming**: Use descriptive names for experimental features
2. **Commit Messages**: Clear messages explaining learning objectives
3. **Code Comments**: Document non-obvious implementations
4. **Testing**: Manually test HTTP endpoints before commits
5. **Documentation**: Update this file when adding new features or modules

---

## Current Status

- **Open Issues**: 1
- **Stars**: 0
- **Forks**: 0
- **License**: None specified
- **Public**: Yes

---

## Related Resources

- [Go Documentation](https://golang.org/doc/)
- [Go HTTP Package](https://golang.org/pkg/net/http/)
- [Go Module Documentation](https://golang.org/doc/modules)
- [Docker Go Best Practices](https://docs.docker.com/language/golang/)

---

## Notes for Contributors

- This is a learning project focused on Go fundamentals
- The repository demonstrates practical implementations of HTTP servers, session management, and Go concepts
- Code may not follow production standards; focus is on learning
- Feel free to explore the `play/` and `sandbox/` directories for experimental implementations

---

*Last Generated: 2026-05-30*
