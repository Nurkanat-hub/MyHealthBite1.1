# MyHealthBite
![CI](https://github.com/Nurkanat-hub/MyHealthBite1.1/actions/workflows/ci.yml/badge.svg)

MyHealthBite is a modular gRPC-based Go application that helps track user statistics and sends email notifications. The system consists of three microservices:

- **user-service**: Main service handling users and interacting with other services
- **stats-service**: Tracks and stores usage statistics
- **email-service**: Sends email notifications

---

## 🛠 Technologies Used

- Golang 1.24.1
- gRPC
- Docker & Docker Compose
- Prometheus
- MongoDB
- GitHub Actions (CI/CD)

---

## 📦 Project Structure


---

## ⚙️ CI/CD Pipeline

This project includes a fully working GitHub Actions workflow:

- Automatically builds and tests on every `push` or `pull_request` to the `main` branch
- Uses Docker Compose to bring up all services
- Runs Go unit tests for `user-service`
- Shuts down services after testing

You can find the workflow in `.github/workflows/ci.yml`.

---

## 🔧 SRE Tool

The project includes a custom SRE tool `sre_tool.go` that checks the health of all microservices. When you run:

```bash
go run sre_tool.go

🔍 Health Check Report (SRE Tool):
- user-service (localhost:50051): ✅ UP
- stats-service (localhost:50054): ✅ UP
- email-service (localhost:50058): ✅ UP
