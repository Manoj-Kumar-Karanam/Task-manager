# 📝 Task Manager API

A simple and efficient **Task Manager application** built with **Go (Golang)**, **Gin**, **GORM**, **PostgreSQL**, **JWT Authentication**, and a **React.js frontend**.  
It provides user authentication and full CRUD operations for managing tasks.

## 🚀 Features


- 🔑 User registration and login with JWT authentication  
- 📌 Task CRUD (Create, Read, Update, Delete) operations  
- 🛡️ Middleware for route protection  
- 🗄️ PostgreSQL database with GORM models  
- 🌐 RESTful API structure  
- 🎨 React.js frontend with protected routes  
- ✅ Unit testing (WIP)  

---

## 🛠️ Tech Stack


**Backend**
- Go (Golang) – Core backend  
- Gin – Web framework  
- GORM – ORM for Golang  
- PostgreSQL – Relational database  
- JWT – Authentication  
- Docker *(optional for containerization)*  

**Frontend**
- React.js – User interface  
- Axios – API requests  
- React Router – Client-side routing  
- TailwindCSS   

---

## 📁 Project Structure

```
Task-manager/
- backend
├── controllers/         # Route handlers
├── middleware/          # JWT middleware
├── models/              # GORM models for User and Task
├── routes/              # API route definitions
├── utils/               # Utility functions (e.g., JWT generation)
├── main.go              # Entry point
├── go.mod / go.sum      # Dependency management
└── .env                 # Environment variables


- frontend/
│── src/
│ ├── assets/ # Static assets (images, logos, etc.)
│ ├── components/ # Reusable components
│ │ ├── ProtectedRoute.jsx
│ │ ├── TaskForm.jsx
│ │ └── TaskList.jsx
│ ├── pages/ # Page-level components
│ │ ├── Dashboard.jsx
│ │ ├── Login.jsx
│ │ └── Register.jsx
│ ├── api.js # API helper functions
│ ├── App.css
│ ├── App.jsx # Root component
│ ├── index.css
│ └── main.jsx # Entry point
```

## 🔐 Authentication

1. User registers or logs in via the frontend.  
2. Backend validates credentials and issues a JWT token.  
3. Frontend stores the token (`localStorage`).  
4. All protected API requests include the token in the `Authorization` header:
5. Backend validates token → grants/denies access.  

## ⚙️ Setup Instructions

## 🎨 Frontend Flow

- **Login/Register Page** → Users can create an account or log in.  
- **Dashboard** → Displays user’s tasks.  
- **Task Management**:  
  - ➕ Add new task  
  - ✏️ Edit/update task  
  - 🗑️ Delete task  
- **Logout** → Clears JWT token and redirects to login.  

---

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

### 3.Backend setup:
cd backend
cp .env.example .env   # configure environment variables
go mod tidy
go run main.go

### 4. Frontend setup:
cd frontend
npm install
npm start


### 5. Install Dependencies

```bash
go mod tidy
npm install
```

### 6. Run the Application

```bash
go run main.go
npm run dev
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


* Task categories/labels
* Task due dates and reminders
* Notifications (email/push)
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
