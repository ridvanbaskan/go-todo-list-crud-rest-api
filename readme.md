# Todo List REST API

This is a Todo List application with a RESTful API built using Go, GORM for PostgreSQL database integration, and Gorilla for HTTP routing.

## Features

- User Authentication: The API supports user authentication using JSON Web Tokens (JWT). Users can register, log in, and access protected routes by providing a valid JWT.

- Task Management: The API allows users to create, read, update, and delete tasks. Each task has properties such as a title, description, due date, priority, and completion status.

- Task Filtering and Sorting: Users can retrieve tasks based on specific criteria such as due date, priority, completion status, or any other custom attributes. The API provides query parameters to filter and sort tasks accordingly.

- Pagination: The API supports pagination of task results, enabling the retrieval of tasks in smaller, manageable chunks. Users can specify the page number and the number of tasks per page.

## Endpoints

- `POST /register`: Register a new user with a unique username and password. Returns a JWT upon successful registration.

- `POST /login`: Log in with a registered username and password. Returns a JWT upon successful login.

- `GET /tasks`: Retrieve all tasks. Supports filtering, sorting, and pagination. Query parameters can be used to filter tasks based on due date, priority, completion status, etc. Pagination parameters (`page` and `limit`) can be used to control the number of tasks returned per page and the current page number.

- `POST /tasks`: Create a new task by providing the necessary details in the request body.

- `GET /tasks/{id}`: Retrieve a specific task by its ID.

- `PUT /tasks/{id}`: Update an existing task by providing the updated details in the request body.

- `DELETE /tasks/{id}`: Delete a specific task by its ID.

## Installation

##### Clone the repository:

```bash
git clone https://github.com/your-username/todo-list-api.git
```