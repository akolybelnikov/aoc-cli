# Advent of Code CLI

A CLI tool to automate downloading inputs for Advent of Code, written in Go.

## Installation

You can install the CLI directly using Go without cloning the repository.

### Requirements
- Go 1.19+ installed

### Install via Go
```bash
go install github.com/akolybelnikov/aoc-cli@latest
```

## Auth - Session Management

### Why Session Management?
Advent of Code (AoC) puzzles and inputs require authentication. AoC uses a session cookie (`session` token) to verify users. This CLI stores the session locally, allowing automated puzzle downloads without repeated logins.

### How It Works:
1. **First Run (No Session Found):**
    - CLI checks for `.aoc-session` in the user's home directory.
    - If not found, the CLI prompts for the session token.
    - The token is saved securely in `.aoc-session` (with `0600` permissions).

2. **Subsequent Runs (Session Exists):**
    - CLI reads the token from `.aoc-session`.
    - The token is used to authenticate and fetch puzzles/inputs.

3. **Security:**
    - Token stored with restricted permissions (user-only read/write).

---

### How to Get Your Session Token:
1. Log in to [Advent of Code](https://adventofcode.com).
2. Open Developer Tools (`F12` or `Ctrl+Shift+I`).
3. Navigate to **Application** -> **Cookies**.
4. Copy the value of the `session` cookie.
5. Paste it into the CLI when prompted.

---

### Notes:
- **Session Expiry:** If the session expires, the CLI will prompt for a new token.
- **No Auto-Login:** The CLI cannot extract the session directly from your browser for security reasons.
- **No OAuth:** AoC does not offer an API, so session cookies are the only way to authenticate.

---

This method keeps authentication simple and mirrors the manual download process while enabling automation.

