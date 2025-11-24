# 📦 Go E-Commerce Backend API

A scalable and production-ready **Golang backend API** built with a clean architecture approach.  
This project provides full **Product CRUD**, **User Authentication**, **JWT token generation**, **repository pattern**, **advanced routing**, and **PostgreSQL** support.

The entire project is structured for maintainability and real-world backend development practices.

---

## 🚀 Features

### 🛒 Product Module
- Full CRUD operations
- Complete Product CRUD API (Create, Read, Update, Delete) including dynamic routes like `GET /products/{productID}`
- Server-side pagination (LIMIT/OFFSET)
- Clean domain-driven structure

### 👤 Authentication Module
- User registration & login
- JWT authentication using **pure-Go HS256** (no external dependencies)
- Secure JWT middleware for protected routes

### 🛠 Advanced Routing
- Feature-based modular routing
- Global middlewares (CORS, logging, auth)
- Clean route definitions and controller separation

### 🗄 Database Layer
- PostgreSQL with **sqlx**
- Repository pattern with interface-based dependency injection
- Efficient, readable SQL queries
- Database migrations (up/down)

### 🔧 Configuration
- `.env` loader using singleton pattern (loaded once)
- Configurable database, JWT secret, and server settings

### 🧱 Architecture
This project follows a DDD-inspired architecture:

