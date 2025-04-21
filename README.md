# Course Platform Microservice

## Overview

The Course Platform Microservice is a robust, scalable backend service designed to manage educational courses. Built with Go and following clean architecture principles, this microservice provides a comprehensive set of RESTful APIs for course management operations.

## Table of Contents

1. [Features](#features)
2. [Technical Specifications](#technical-specifications)
3. [API Documentation](#api-documentation)
4. [Architecture](#architecture)
5. [Getting Started](#getting-started)
6. [Development Guide](#development-guide)
7. [Configuration](#configuration)
8. [Future Roadmap](#future-roadmap)
9. [License](#license)

## Features

### Core Functionality
- Course creation and management
- Course listing and retrieval
- Course deletion
- Course search functionality
- In-memory data persistence

### Technical Features
- RESTful API design
- Clean architecture implementation
- Middleware support
- Error handling
- Logging system

## Technical Specifications

### Requirements
- Go 1.21 or later
- Docker 20.10 or later
- Docker Compose 2.0 or later

### Dependencies
- Echo (Web Framework)
- Go Modules (Dependency Management)
- Standard Library Components

## API Documentation

### Endpoints

#### 1. Course Creation
- **Endpoint**: `POST /v1/courses`
- **Description**: Creates a new course entry
- **Request Body**:
  ```json
  {
      "title": "Introduction to Go",
      "description": "Learn the basics of Go programming",
      "author": "John Doe",
      "price": 99.99
  }
  ```
- **Response**: 
  ```json
  {
      "id": "course-123"
  }
  ```
- **Status Codes**: 
  - 201: Created
  - 400: Bad Request
  - 500: Internal Server Error

#### 2. Course Listing
- **Endpoint**: `GET /v1/courses`
- **Description**: Retrieves all available courses
- **Response**:
  ```json
  {
      "courses": [
          {
              "id": "course-123",
              "title": "Introduction to Go",
              "description": "Learn the basics of Go programming",
              "author": "John Doe",
              "price": 99.99
          }
      ]
  }
  ```
- **Status Codes**:
  - 200: OK
  - 500: Internal Server Error

#### 3. Course Retrieval
- **Endpoint**: `GET /v1/courses/:id`
- **Description**: Retrieves a specific course by ID
- **Response**:
  ```json
  {
      "id": "course-123",
      "title": "Introduction to Go",
      "description": "Learn the basics of Go programming",
      "author": "John Doe",
      "price": 99.99
  }
  ```
- **Status Codes**:
  - 200: OK
  - 404: Not Found
  - 500: Internal Server Error

#### 4. Course Deletion
- **Endpoint**: `DELETE /v1/courses/:id`
- **Description**: Removes a course from the system
- **Response**: 204 No Content
- **Status Codes**:
  - 204: No Content
  - 404: Not Found
  - 500: Internal Server Error

#### 5. Course Search
- **Endpoint**: `POST /v1/courses/search`
- **Description**: Searches courses by keyword
- **Request Body**:
  ```json
  {
      "query": "Go"
  }
  ```
- **Response**:
  ```json
  {
      "courses": [
          {
              "id": "course-123",
              "title": "Introduction to Go",
              "description": "Learn the basics of Go programming",
              "author": "John Doe",
              "price": 99.99
          }
      ]
  }
  ```
- **Status Codes**:
  - 200: OK
  - 400: Bad Request
  - 500: Internal Server Error

## Architecture

### System Design
```
MS/
├── cmd/                 # Application entry points
├── internal/           # Private application code
│   ├── app/           # Application setup and configuration
│   ├── deliveries/    # Delivery mechanisms (HTTP, gRPC, etc.)
│   ├── repositories/  # Data access layer
│   ├── services/      # Business logic layer
│   └── usecases/      # Application use cases
├── pkg/               # Public library code
│   ├── domain/        # Domain models
│   └── reqresp/       # Request/response models
└── test/              # Test files
```

### Architecture Principles
1. **Domain Layer**
   - Business entities
   - Business rules
   - Domain logic

2. **Use Case Layer**
   - Application-specific business rules
   - Use case implementations
   - Business logic orchestration

3. **Interface Adapters Layer**
   - API controllers
   - Data converters
   - External interface implementations

4. **Frameworks & Drivers Layer**
   - Web framework
   - Database connections
   - External service integrations

## System Diagrams

### Entity Relationship Diagram (ERD)

The following diagram represents the data model for the Course Platform:

```
+------------------+       +------------------+
|     Course       |       |     Author       |
+------------------+       +------------------+
| PK: id           |       | PK: id           |
| title            |       | name             |
| description      |       | email            |
| price            |       | bio              |
| FK: author_id    |------>| created_at       |
| created_at       |       | updated_at       |
| updated_at       |       +------------------+
| status           |
+------------------+
```

#### Entity Descriptions

1. **Course**
   - `id`: Unique identifier (Primary Key)
   - `title`: Course title
   - `description`: Course description
   - `price`: Course price
   - `author_id`: Reference to author (Foreign Key)
   - `created_at`: Creation timestamp
   - `updated_at`: Last update timestamp
   - `status`: Course status (active/inactive)

2. **Author**
   - `id`: Unique identifier (Primary Key)
   - `name`: Author's full name
   - `email`: Author's email address
   - `bio`: Author's biography
   - `created_at`: Creation timestamp
   - `updated_at`: Last update timestamp

### Sequence Diagrams

The system's main business flows are documented in PlantUML sequence diagrams located in `docs/sequence.puml`. These diagrams illustrate:

1. **Course Creation Flow**
   - Client request handling
   - Service layer processing
   - Use case execution
   - Repository interaction
   - Response generation

2. **Course Retrieval Flow**
   - Single course retrieval
   - Course listing
   - Error handling
   - Response formatting

3. **Course Deletion Flow**
   - Deletion request processing
   - Repository cleanup
   - Success confirmation

4. **Course Search Flow**
   - Search query processing
   - Result filtering
   - Response generation

