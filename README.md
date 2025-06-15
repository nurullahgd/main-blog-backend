# Blog Backend API

This project is a REST API developed using the Go Fiber framework for a modern blog platform.

## 🚀 Features

- 🔐 JWT-based authentication
- 👥 User management
- 📝 Blog posts CRUD operations
- 🖼️ Image upload with Cloudinary integration
- 🐳 Docker support
- 🔄 CORS configuration
- 📊 Database operations with GORM
- 🛡️ Middleware support

## 🛠️ Technologies

- Go 1.x
- Fiber Framework
- GORM
- PostgreSQL
- Docker
- Cloudinary
- JWT

## 📋 Requirements

- Go 1.x or higher
- Docker and Docker Compose
- PostgreSQL
- Cloudinary account

## 🚀 Installation

1. Clone the project:
```bash
git clone https://github.com/nurullahgd/main-blog-backend.git
cd main-blog-backend
```

2. Install dependencies:
```bash
go mod download
```

3. Create `.env` file and set required variables:
```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=5432
JWT_SECRET=your_jwt_secret
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

4. Run with Docker:
```bash
docker-compose up -d
```

## 📁 Project Structure

```
.
├── controllers/     # HTTP request handlers
├── models/         # Database models
├── routes/         # API routes
├── middleware/     # Middleware functions
├── utils/          # Utility functions
├── helpers/        # Helper packages
├── database/       # Database connection and configuration
├── uploads/        # Temporary directory for uploaded files
├── main.go         # Main application file
├── Dockerfile      # Docker configuration
└── docker-compose.yml
```

## 🔒 API Endpoints

### User Operations
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - User login
- `GET /api/users/profile` - Get user profile
- `PUT /api/users/profile` - Update profile

### Blog Operations
- `GET /api/blogs` - List all blog posts
- `GET /api/blogs/:id` - Get specific blog post
- `POST /api/blogs` - Create new blog post
- `PUT /api/blogs/:id` - Update blog post
- `DELETE /api/blogs/:id` - Delete blog post


## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 👥 Contributing

1. Fork this repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Contact

Tests coming soon

Nurullah Gündoğdu - [@nurullahgd](https://github.com/nurullahgd) (https://github.com/nurullahgd)
