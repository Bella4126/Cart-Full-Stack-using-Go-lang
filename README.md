# Shopping Cart Application

A full-stack shopping cart web application built with Go, Gin, Gorm, PostgreSQL for the backend and React for the frontend.

## Backend Technologies
- Go
- Gin (Web Framework)
- Gorm (ORM)
- PostgreSQL (Database)
- JWT Authentication
- bcrypt (Password Hashing)

## Frontend Technologies
- React
- Vanilla CSS
- Fetch API

## Features
- User registration and authentication
- Single device login (one token per user)
- Item management
- Shopping cart functionality
- Order management
- Protected routes with JWT middleware

## Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 12+

## Setup Instructions

### Database Setup
1. Install PostgreSQL and create a database named `shopping_cart`
2. Update database credentials in `backend/database/database.go`

### Backend Setup
```bash
cd backend
go mod init shopping-cart
go mod tidy
go run main.go
```

The backend server will start on port 8080.

### Frontend Setup
```bash
npm install
npm run dev
```

The frontend will start on port 5173.

## API Endpoints

### Public Endpoints
- `POST /users` - Create new user
- `GET /users` - List all users
- `POST /users/login` - User login
- `POST /items` - Create new item
- `GET /items` - List all items

### Protected Endpoints (Require Authentication)
- `POST /carts` - Add item to cart
- `GET /carts` - List all carts
- `POST /orders` - Convert cart to order
- `GET /orders` - List all orders

## Authentication
All cart and order operations require a valid JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

## Application Flow
1. User starts at login screen
2. Can register new account or login
3. After login, sees item list
4. Click items to add to cart
5. View cart contents with Cart button
6. View order history with Order History button
7. Checkout converts cart to order
8. Success toast shown after checkout

## Database Schema
- `users` - User accounts with hashed passwords and tokens
- `items` - Available products
- `carts` - User cart items
- `orders` - Completed orders
- `order_items` - Individual items in orders

## Security Features
- Password hashing with bcrypt
- JWT token authentication
- Single device login enforcement
- Protected route middleware
- CORS configuration