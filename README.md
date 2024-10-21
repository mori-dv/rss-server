# RSS Server

This project is an RSS server implemented in Go, designed to manage and distribute RSS feeds to users. It features a RESTful API using `chi` for routing and provides user authentication, feed management, and content scraping.

## Features

- User registration and management
- RSS feed creation, following, and retrieval
- Authentication for protected routes
- Automated feed scraping

## Requirements

- Go 1.16+
- PostgreSQL

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/mori-dv/rss-server.git
   ```
2. Set up environment variables by creating a `.env` file based on `.env.example`:
   ```
   BASE_DIR=<path_to_project>
   PORT=<port_number>
   DB_URL=<database_url>
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run
   ```

## Endpoints

- **GET** `/v1/healthz` - Check server health
- **POST** `/v1/user/new` - Create a new user
- **GET** `/v1/user/get` - Get user details (requires authentication)
- **POST** `/v1/feed/new` - Create a new feed (requires authentication)
- **GET** `/v1/feeds` - Get all available feeds
- **GET** `/v1/posts` - Get posts for a user (requires authentication)

## Middleware

- **CORS**: Configured to allow requests from all origins.
- **Authentication**: Protected endpoints require user authentication.

## License

This project is licensed under the MIT License.

