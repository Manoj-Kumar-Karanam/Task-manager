# 📝 Task Manager API

A simple and efficient Task Manager RESTful API built with Go (Golang), Gin framework, GORM ORM, PostgreSQL, and JWT authentication. This project helps manage users and their tasks with basic authentication and CRUD operations.

## 🚀 Features

* User registration and login with JWT authentication
* Task CRUD (Create, Read, Update, Delete) operations
* Middleware for route protection
* PostgreSQL database with GORM models
* RESTful API structure
* Unit testing (WIP)

## 🛠️ Tech Stack

* [Go (Golang)](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin) – Web framework
* [GORM](https://gorm.io/) – ORM for Golang
* [PostgreSQL](https://www.postgresql.org/) – Relational database
* [JWT](https://github.com/golang-jwt/jwt) – Authentication
* [Docker](https://www.docker.com/) *(optional for containerization)*

---

## 📁 Project Structure

```
Task-manager/
├── controllers/         # Route handlers
├── middleware/          # JWT middleware
├── models/              # GORM models for User and Task
├── routes/              # API route definitions
├── utils/               # Utility functions (e.g., JWT generation)
├── main.go              # Entry point
├── go.mod / go.sum      # Dependency management
└── .env                 # Environment variables
```

---

## 🔐 Authentication

Authentication is handled using JWTs (JSON Web Tokens). Upon successful login, a token is returned and should be used in the `Authorization` header for protected routes.

```
Authorization: Bearer <your_token>
```

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Manoj-Kumar-Karanam/Task-manager.git
cd Task-manager
```

### 2. Set Up Environment Variables

Create a `.env` file in the root directory with the following:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=task_db
JWT_SECRET=your_jwt_secret
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run main.go
```

---

## 📬 API Endpoints

### Auth

* `POST /register` – Register a new user
* `POST /login` – Authenticate and get JWT token

### Tasks (Protected)

* `GET /tasks` – Get all tasks for the logged-in user
* `POST /tasks` – Create a new task
* `PUT /tasks/:id` – Update a task
* `DELETE /tasks/:id` – Delete a task

---


## 📌 Future Improvements

* Role-based access control (Admin/User)
* Task categories/labels
* Task due dates and reminders
* Frontend integration with React
* CI/CD setup with GitHub Actions

---

## 🙌 Contributing

Contributions are welcome! Feel free to fork the repo and submit a pull request.

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).

---

## 👨‍💻 Author

**Manoj Kumar Karanam**
[GitHub](https://github.com/Manoj-Kumar-Karanam)
