# Go CLI To-Do App

A simple command-line interface (CLI) application for managing your to-do tasks, written in Go.

-----

## Features

  * **Add** new tasks.
  * **Update** existing tasks by ID.
  * **List** all tasks or filter them by status.
  * **Mark** tasks as **Done** or **In-Progress**.

-----

## 🛠️ Installation

### Prerequisites

You need to have **Go** installed on your system.

### Build from source

1.  Clone the repository:
    ```bash
    git clone https://github.com/mu7ammad1951/task-tracker-cli 
    cd task-tracker-cli
    ```
2.  Build the executable:
    ```bash
    go build -o todo
    ```
3.  (Optional) Move the executable to a directory in your system's PATH to run it from anywhere (e.g., `/usr/local/bin` on Linux/macOS):
    ```bash
    mv todo /usr/local/bin/
    ```

-----

## 📝 Usage

The main executable is assumed to be named `todo` (or the name you chose during the build, like `task-tracker-cli`).

### Commands

| Command | Usage | Description |
| :--- | :--- | :--- |
| **`add`** | `todo add "task description"` | Adds a new task with the given description. |
| **`update`** | `todo update <id> "new description"` | Updates the description of the task with the specified `<id>`. |
| **`list`** | `todo list` | Lists all tasks and their current status. |
| **`list`** | `todo list [filter]` | Lists tasks filtered by status (e.g., `todo`, `done`, or `in-progress`). |
| **`mark-done`** | `todo mark-done <id>` | Marks the task with the specified `<id>` as **Done**. |
| **`mark-in-progress`** | `todo mark-in-progress <id>` | Marks the task with the specified `<id>` as **In-Progress**. |

### Examples

Add a task:

```bash
$ todo add "Complete the project documentation"
```

List all tasks:

```bash
$ todo list
```

Mark a task as done (assuming it has ID `1`):

```bash
$ todo mark-done 1
```

List only in-progress tasks:

```bash
$ todo list in-progress
```

Update an existing task:

```bash
$ todo update 1 "Finalize the main logic"
```