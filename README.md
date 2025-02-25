# Church Leadership Training Platform

A comprehensive web and mobile platform for training pastors and church leaders, featuring media streaming, assessments, and real-time communication capabilities.

## 🚀 Features

- **Authentication & Authorization**
  - Secure user registration and login
  - Role-based access control (Admin, Trainer, Student)
  - JWT-based authentication

- **Media Streaming**
  - Audio/video content delivery
  - Adaptive bitrate streaming
  - Content management system
  - Offline access support

- **Assessment System**
  - Quiz creation and management
  - Automated scoring
  - Progress tracking
  - Performance analytics

- **Real-time Communication**
  - Live chat functionality
  - Discussion forums
  - Peer-to-peer messaging
  - Group discussions

## 🛠 Technology Stack

### Backend (Golang)

- **Framework**: Gin v1.9.x
- **Databases**:
  - PostgreSQL 15.x (User data, Assessments)
  - MongoDB 6.x (Chat messages)
  - Redis 7.x (Caching)
- **Message Queue**: RabbitMQ
- **Media Server**: Nginx-RTMP
- **API Gateway**: Kong

### Frontend (React)

- React 18.x
- TypeScript 5.x
- Redux Toolkit
- Material-UI v5.x
- React Query
- Socket.io

### Mobile (Flutter)

- Flutter SDK 3.x
- BLoC Pattern
- Dio for HTTP
- Hive for storage
- WebSocket for real-time features

### DevOps

- Docker & Docker Compose
- Kubernetes
- GitHub Actions
- Prometheus & Grafana
- ELK Stack

## 📋 Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- Flutter 3.16.0 or higher
- Docker & Docker Compose
- Kubernetes (optional for production)

## 🚦 Getting Started

### Local Development Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/collageman/training_app.git
    cd training_app
    ```env

2. Start the infrastructure services:

    ```bash
    cd docker
    docker-compose up -d
    ```

3. Start the backend services:

    ```bash
    # For each service (auth, media, assessment, chat)
    cd backend/[service]-service
    go mod tidy
    go run cmd/main.go
    ```

4. Start the frontend web application:

    ```bash
    cd frontend/web
    npm install
    npm run dev
    ```

5. Start the Flutter mobile app:

    ```bash
    cd mobile/flutter_app
    flutter pub get
    flutter run
    ```

### Environment Variables

Create `.env` files in each service directory. Example for auth-service:

```env
# Auth Service
DB_HOST=localhost
DB_PORT=5432
DB_NAME=church_training
DB_USER=admin
DB_PASSWORD=adminpass
JWT_SECRET=your-secret-key
```

## 📁 Project Structure

```plaintext
church-training-platform/
├── backend/
│   ├── auth-service/
│   ├── media-service/
│   ├── assessment-service/
│   └── chat-service/
├── frontend/
│   └── web/
├── mobile/
│   └── flutter_app/
├── k8s/
└── docker/

## 🔄 API Endpoints

### Authentication Service

- POST /api/auth/register
- POST /api/auth/login
- POST /api/auth/refresh
- GET /api/auth/profile

### Media Service

- GET /api/media/content
- GET /api/media/stream/{id}
- POST /api/media/upload
- PUT /api/media/{id}

### Assessment Service

- GET /api/assessments
- POST /api/assessments/create
- POST /api/assessments/submit
- GET /api/assessments/results

### Chat Service

- WebSocket /ws/chat
- GET /api/chat/messages
- POST /api/chat/send

## 🧪 Testing

### Backend Testing

```bash
cd backend/[service]-service
go test ./...
```

### Frontend Testing

```bash
cd frontend/web
npm test
```

### Mobile Testing

```bash
cd mobile/flutter_app
flutter test
```

## 🚀 Deployment

### Docker Deployment

```bash
docker-compose -f docker/docker-compose.prod.yml up -d
```

### Kubernetes Deployment

```bash
kubectl apply -k k8s/overlays/prod
```

## 📈 Monitoring

- Prometheus: <http://localhost:9090>
- Grafana: <http://localhost:3000>
- RabbitMQ Management: <http://localhost:15672>

## 🔐 Security Considerations

- All endpoints are SSL/TLS encrypted
- JWT tokens for authentication
- Rate limiting implemented
- Input validation on all endpoints
- Media content access control
- End-to-end encryption for chat

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 👥 Team

- Backend Team: [Names]
- Frontend Team: [Names]
- Mobile Team: [Names]
- DevOps Team: [Names]

## 📞 Support

For support, email <jefferyasamani7@gmail.com>
