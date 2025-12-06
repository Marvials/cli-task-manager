# CLI Task Manager

A robust, efficient, and thoroughly architected Command Line Interface (CLI) tool for managing your daily tasks. Built with Golang and backed by PostgreSQL, it leverages modern development patterns to ensure reliability and speed.

## 🚀 Features
Task Management: Create, read, update, and delete tasks directly from your terminal.

Natural Typing: Add tasks using natural sentences without needing quotes (e.g., task add Buy coffee).

Status Tracking: Filter tasks by status (To do, Doing, Done).

Timezone Aware: Correctly handles timezones using TIMESTAMPTZ, ensuring task creation times are accurate regardless of where you are.

Persistent Storage: Uses PostgreSQL for reliable data persistence.

Clean Architecture: Built using the Factory pattern, separation of concerns (Service/Repository layers), and Dependency Injection.

## 🛠️ Tech Stack
Language: Go (Golang)

Database: PostgreSQL

CLI Framework: Cobra

Database Driver: pgx

Configuration: godotenv

## 📋 Prerequisites
Before you begin, ensure you have the following installed:

Go (version 1.22+)

PostgreSQL

## ⚙️ Installation & Setup
1. Clone the repository
```Bash
git clone https://github.com/Marvials/cli-task-manager.git
cd cli-task-manager
``` 

2. Build the application
Linux / macOS:
```Bash
go build -o task cmd/main.go
```
Optionally, you can move the binary to your path:

```Bash
mv task /usr/local/bin/
```
Windows (PowerShell):
```Powershell
go build -o task.exe cmd/main.go
```

3. Configuration
The application requires a configuration file to connect to your database. It looks for a file named .task-manager.env in your user's .config directory.

Linux / macOS
```Bash
mkdir -p ~/.config
nano ~/.config/.task-manager.env
```
Windows (PowerShell)
```PowerShell
New-Item -ItemType Directory -Force -Path "$HOME\.config"

notepad "$HOME\.config\.task-manager.env"
```

Create the file ~/.config/.task-manager.env and add your database credentials:

```
DB_HOST=your_host_or_IP
DB_PORT=your_postgres_port
DB_USER=your_postgres_user
DB_PASSWORD=your_postgres_password
DB_NAME=your_database_name
```

4. Initialize Database
Run the setup command to create the necessary tables in your database:

```Bash
./task tables
```
## 💻 Usage
Create a Task
Add a new task. You don't need quotes for multi-word descriptions.

```Bash
task add Buy groceries for the week
```
List Tasks
View your tasks in a formatted table. By default, it shows "To do" tasks.

```Bash

task list           # Lists pending tasks
task list --doing   # Lists tasks in progress
task list --done    # Lists completed tasks
task list --all     # Lists all tasks
```

Update Status
Change the status of a task using its ID. Accepted statuses: To do, Doing, Done.

```Bash
task change 1 Doing
task update 5 Done
```
View Details
See full details for a specific task, including relative creation time.

```Bash
task get 1
```
Delete Task
Permanently remove a task.

```Bash
task delete 1
```
## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.