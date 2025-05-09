```
██████╗  ██╗███████╗██████╗ ██╗  ██╗ ██████╗ ██╗     ██╗██████╗ ██╗   ██╗███████╗
██╔══██╗███║██╔════╝██╔══██╗██║  ██║██╔═████╗██║    ███║██╔══██╗██║   ██║██╔════╝
██║  ██║╚██║███████╗██████╔╝███████║██║██╔██║██║    ╚██║██║  ██║██║   ██║███████╗
██║  ██║ ██║╚════██║██╔═══╝ ██╔══██║████╔╝██║██║     ██║██║  ██║██║   ██║╚════██║
██████╔╝ ██║███████║██║     ██║  ██║╚██████╔╝███████╗██║██████╔╝╚██████╔╝███████║
╚═════╝  ╚═╝╚══════╝╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝╚═════╝  ╚═════╝ ╚══════╝
```


#  API-GAME-CATALOG-CL

Api for games catalog of CityLog app.

## 🚀 Objective
Create an API for the games available in CityLog app.

---

## 🧱 Tech Stack

| Domain              | Tools used                                       |
|---------------------|--------------------------------------------------|
| Infrastructure      | Terraform, OpenStack                             |
| Orchestration       | Kubernetes (K3s)                                 |
| Containerization    | Docker                                           |
| Main language       | Golang                                           |
| CI/CD               | GitHub Actions                                   |
| Backup              | Bash scripts, PVCs, CronJobs                     |

---

## 📦 Features

- Infrastructure as Code (IaC)
- CI/CD
- Automated backups

---

## 📌 Roadmap

- [x] Grounds
- [ ] Clean-code ++
- [ ] K3s cluster with TLS ingress
- [ ] GitHub Actions integration (build/test/deploy)
- [ ] Backup and restoration scripts
- [ ] Final documentation

---


## Project Structure

```
api-game-catalog/
├── api/
│   ├── proto/
│   │   ├── catalog.proto
│   │   ├── review.proto
│   │   └── user.proto
│   └── swagger/
├── cmd/
│   ├── catalog-service/
│   │   └── main.go
│   ├── review-service/
│   │   └── main.go
│   └── user-service/
│       └── main.go
├── internal/
│   ├── catalog/
│   │   ├── models/
│   │   ├── repository/
│   │   ├── service/
│   │   └── delivery/
│   ├── review/
│   │   ├── models/
│   │   ├── repository/
│   │   ├── service/
│   │   └── delivery/
│   ├── user/
│   │   ├── models/
│   │   ├── repository/
│   │   ├── service/
│   │   └── delivery/
│   └── pkg/
│       ├── middleware/
│       ├── auth/
│       ├── logger/
│       └── errors/
├── pkg/
│   ├── grpc/
│   │   └── client.go
│   ├── database/
│   │   ├── postgres.go
│   │   └── mongodb.go
│   └── cache/
│       └── redis.go
├── deployments/
│   ├── docker-compose.yml
│   └── kubernetes/
├── scripts/
├── go.mod
└── go.sum
```
