# KinRelay 🔗  
![Go Version](https://img.shields.io/badge/Go-1.21-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![PostgreSQL](https://img.shields.io/badge/database-PostgreSQL-blue)
![Status](https://img.shields.io/badge/status-In_Development-orange)
A system to get you connected with your emergency contacts when all else fails

# Screenshots 🖼️

### Sign-In Page
![Sign-In](./static/screenshots/signin.png)

### Contact List Page
![Contacts](./static/screenshots/contacts.png)

# KinRelay

KinRelay is a lightweight, self-hosted emergency contact system. It allows users to securely access a list of predefined contacts using a simple 6-digit PIN. The project is built with Go, PostgreSQL, and a minimalist frontend served via NGINX.

---

## Features ✨

- 🔐 PIN-based access with no accounts or passwords
- 🗃️ Contact list is stored securely in a PostgreSQL database
- ⚡ Simple, responsive frontend designed for quick access
- 🔧 REST API built in Go with middleware-based authentication
- 🌐 Static assets and frontend served via NGINX

---

## Project Structure 🗂️

```
kinrelay/
├── backend/         # Go API, middleware, handlers, models
├── frontend/        # HTML, CSS, and JS for the UI
├── static/          # Logo, favicon, and other public assets
└── .env             # Environment configuration
```

---

## How It Works ⚙️

1. **Sign In (index.html)**  
   The user enters a 6-digit PIN. This is stored in the browser using `localStorage` and passed to the backend as a custom header.

2. **Contact List (contacts.html)**  
   If the PIN is valid, the backend returns a list of emergency contacts, which are displayed in a clean list format.

---

## Stack 🧰

- Go (Gorilla Mux for routing)
- PostgreSQL
- HTML/CSS/JavaScript
- NGINX (for static file serving and API proxy)

---

## Getting Started 🚀

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

4. Make sure NGINX is set up to serve the frontend and reverse proxy `/api/` to the Go server.

---

## Security 🔒

- All API requests require a valid `X-PIN` header
- No user accounts or session tokens
- Designed for short, secure interactions

---

## Roadmap 🛣️

- Add/edit/delete contact support
- Optional PIN expiry or rotation
- Admin-only contact management
- UI improvements for accessibility

---

## License 📄

MIT License

---

Created by Tom Hornbuckle 👤