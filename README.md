# URL Shortener

A modular URL shortener service built with Go and MySQL. This project provides a simple yet scalable way to shorten long URLs, redirect users to the original URLs, and track click analytics. It follows a clean architecture with separated concerns for maintainability and extensibility.

## Features

- **URL Shortening**: Generate unique, 8-character short codes for long URLs.
- **Redirection**: Redirect users from short URLs to their original destinations.
- **Click Analytics**: Track the number of clicks for each short URL.
- **RESTful API**: Simple and intuitive endpoints for integration.
- **Modular Design**: Organized into configuration, database, repository, service, handler, and route layers.
- **MySQL Integration**: Persistent storage for URLs and analytics.

## Prerequisites

To run this project locally, ensure you have the following installed:

- **Go**: Version 1.22 or higher (`go version`)
- **MySQL**: Version 8.0 or higher
- **Git** (optional): For cloning the repository
- **Unix-like system** (e.g., Ubuntu)

## Project Structure

```
url-shortener/
├── cmd/
│   └── api/
│       └── main.go           # Entry point for the API server
├── internal/
│   ├── config/              # Configuration loading (e.g., DB, server settings)
│   ├── db/                  # Database connection and initialization
│   ├── handlers/            # HTTP handlers for API endpoints
│   ├── models/              # Data models (e.g., URL, analytics)
│   ├── repository/          # Database operations (CRUD)
│   ├── routes/              # API route definitions
│   └── services/            # Business logic for URL shortening and analytics
├── go.mod                   # Go module file
├── go.sum                   # Go module dependencies
├── .env                     # Environment variables
└── README.md                # Project documentation
```

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/username/url-shortener.git
cd url-shortener
```

> Replace `username` with your GitHub username or the actual repository URL.

### 2. Install MySQL

#### Install and Configure MySQL on Ubuntu:

```bash
sudo apt update
sudo apt install mysql-server
```

#### Start MySQL:

```bash
sudo systemctl start mysql
sudo systemctl enable mysql
```

#### Verify MySQL:

```bash
sudo systemctl status mysql
```

Ensure it shows `Active: active (running)`.

#### Secure MySQL (Recommended):

```bash
sudo mysql_secure_installation
```

Follow the prompts and set a root password.

#### Create the Database:

```bash
mysql -u root -p
```

```sql
CREATE DATABASE url_shortener;
EXIT;
```

### 3. Configure Environment Variables

Create a `.env` file:

```bash
nano .env
```

Add the following content:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=url_shortener
SERVER_PORT=8080
```

> Replace `yourpassword` with your MySQL root password. Adjust `DB_USER` and `DB_PASSWORD` as necessary.

### 4. Install Go Dependencies

Navigate to the project directory and install the required dependencies:

```bash
cd ~/path/to/url-shortener
go mod tidy
```

### 5. Run the Application

Start the Go server:

```bash
go run cmd/api/main.go
```

You should see:

```
Server starting on :8080
```

### 6. Test the API

#### Shorten a URL:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten
```

**Response:**
```json
{"short_url":"http://localhost:8080/abc12345"}
```

#### Access a Short URL:

```bash
curl -L http://localhost:8080/abc12345
```

Redirects to `https://example.com`

#### Check Analytics:

```bash
curl http://localhost:8080/analytics/abc12345
```

**Response:**
```json
{"original_url":"https://example.com","click_count":1,"created_at":"2025-04-25T12:00:00Z"}
```

## API Endpoints

| Method | Endpoint               | Description                     | Request Body                       | Response Body                          |
|--------|------------------------|---------------------------------|-------------------------------------|----------------------------------------|
| POST   | /shorten               | Create a short URL              | `{ "url": "https://example.com" }` | `{ "short_url": "http://..." }`       |
| GET    | /{shortCode}           | Redirect to the original URL    | None                                | Redirect                               |
| GET    | /analytics/{shortCode} | Get analytics for a short URL   | None                                | `{ "original_url": ..., ... }`       |

## Running in the Background (Optional)

To run the application in the background:

```bash
nohup go run cmd/api/main.go &
```

Logs will be written to `nohup.out`.

To stop the background process:

```bash
ps aux | grep url-shortener
kill <PID>
```

## Future Enhancements

- Add user authentication (e.g., JWT) for private URL management
- Support custom short codes
- Implement rate limiting to prevent abuse
- Add URL expiration dates
- Build a web frontend (e.g., with React or HTML/CSS)

## Contributing

Contributions are welcome!

1. Fork the repository
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit changes:
   ```bash
   git commit -m "Add feature"
   ```
4. Push the branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or feedback, open an issue on GitHub or contact the maintainer at [jimitchavdadev@gmail.com](mailto:jimitchavdadev@gmail.com).

