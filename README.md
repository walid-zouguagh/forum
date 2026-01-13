# Forum Project 🗣️

A lightweight **social media/forum web application** built with **Go** for the backend and vanilla **HTML/CSS/JavaScript** for the frontend.

Users can share posts and pictures, comment, like/dislike content, and engage with the community.

**Upcoming Feature**: Private messaging/conversations between users!

## Features

- Create and share text posts
- Upload and share pictures
- Comment on posts and pictures
- Like or dislike posts, pictures, and comments
- Full user interaction system
- SQLite database for persistence
- Clean, modular Go architecture

## Live Demo

No live demo available yet. Run it locally to try it out!

## Technologies Used

- **Backend**: Go (Gin or similar routing, based on structure)
- **Database**: SQLite (`database.sqlite`)
- **Frontend**: HTML, CSS, JavaScript (served from `views/` and `static/`)
- **Containerization**: Docker support

## Project Structure
    forum/
    ├── api/              # API endpoints
    ├── cmd/              # Entry point (main.go)
    ├── controllers/      # Request handlers
    ├── middleware/       # Authentication, logging, etc.
    ├── models/           # Data structures and DB interactions
    ├── router/           # Route definitions
    ├── static/           # CSS, JS, images
    ├── utils/            # Helper functions
    ├── views/            # HTML templates
    ├── database.sqlite   # SQLite database file
    ├── Dockerfile        # Docker configuration
    ├── go.mod / go.sum   # Go modules
    ├── cleandocker.sh    # Docker cleanup script
    ├── rundocker.sh      # Docker run script
    └── README.md         # This file


## How to Run Locally

### Prerequisites

- **Go** (version 1.20 or higher recommended) → Download from [go.dev](https://go.dev/dl/)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/walid-zouguagh/forum.git

2. Navigate to the project directory:
    ```bash
    cd forum

3. Install dependencies:
    ```bash
    go mod tidy

4. Run the application:
    ```bash
    go run cmd/main.go

The server will start (usually on http://localhost:8080 or similar — check console output). Open it in your browser!

## Alternative: Run with Docker
    If you have Docker installed:
    ```bash
    ./rundocker.sh   # Builds and runs the container

    To clean up:
    ```bash
    ./cleandocker.sh

## Contributing
    Contributions are welcome! Feel free to:
    . Fix bugs
    . Add new features (e.g., private messaging, user profiles, notifications)
    . Improve security or performance
    . Enhance the UI/UX

    Fork the repo, create a branch, and submit a Pull Request.

## Enjoy building your community! 🚀
    Made with 💻 by 
    walid zouguagh
    Zakaria bessadou
    yassine rahhaoui
    yassine bahbib
    achraf margoum