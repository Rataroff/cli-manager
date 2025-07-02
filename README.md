# cli task manager

A command-line task manager written in Go

## Features

- Add, list, edit, delete, and mark tasks as done
- Persist tasks to a local JSON file
- Clear all completed tasks
- Help command with usage instructions

## Installation

```bash
git clone https://github.com/Rataroff/cli-manager.git
cd cli-manager
go build -o taskmanager
```
## Usage Examples

Add a new task:
```bash
./taskmanager add "Buy groceries"
```
Mark a task as done:
```bash
./taskmanager done 1
```
Edit a task description:
```bash
./taskmanager edit 1 "Buy new PC"
```
Delete a task:
```bash
./taskmanager delete 1
```
Show help:
```bash
./taskmanager help
```


