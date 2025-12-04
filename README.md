# KinRelay

![Go Version](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![PostgreSQL](https://img.shields.io/badge/database-PostgreSQL-blue)
![Status](https://img.shields.io/badge/status-In_Development-orange)

A lightweight, self-hosted emergency contact system designed to help you reach the right people quickly when it matters.

## Screenshots

![Sign-In](./static/screenshots/signin.png)

![Contacts](./static/screenshots/contacts.png)

## Overview

KinRelay lets users access a predefined list of emergency contacts using a simple 6-digit PIN. It is built with a Go REST API, PostgreSQL for storage, and a minimalist frontend served through NGINX.

## Features

* Fast PIN-based access without accounts or passwords
* Emergency contacts stored in PostgreSQL
* Simple responsive UI for quick use on desktop or mobile
* Go REST API with middleware-based request authentication
* Frontend and static assets served via NGINX, with optional API reverse proxy

## Project Structure

```text
kinrelay/
├── backend/         # Go API, middleware, handlers, models
├── frontend/        # HTML, CSS, and JS for the UI
├── static/          # Logo, favicon, and other public assets
└── .env             # Environment configuration
```

## How It Works

1. **Sign In (index.html)**
   The user enters a 6-digit PIN. The PIN is stored in the browser using `localStorage` and sent to the backend using a custom header.

2. **Contact List (contacts.html)**
   If the PIN is valid, the backend returns the contact list and the frontend renders it in a clean, readable layout.

## Stack

* Go (Gorilla Mux for routing)
* PostgreSQL
* HTML, CSS, JavaScript
* NGINX (static serving and optional reverse proxy for the API)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourname/kinrelay.git
   ```

2. Create a `.env` file with your database URL:

   ```bash
   KINRELAY_DB_URL=postgres://user:pass@localhost:5432/kinrelay_db
   ```

3. Run the Go API:

   ```bash
   go run backend/main.go
   ```

4. Configure NGINX to serve the frontend and reverse proxy `/api/` to the Go server.

## Security Notes

* All API requests require a valid `X-PIN` header
* No user accounts, sessions, or tokens
* Designed for short, focused interactions with minimal moving parts

## Roadmap

* Create, edit, and delete contacts
* Optional PIN expiry or rotation
* Admin-only contact management
* Accessibility and UI refinements

## License

MIT License

Created by Tom Hornbuckle
