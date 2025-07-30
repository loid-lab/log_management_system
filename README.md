# 📝 Log Management System

A simple interactive CLI-based log management tool written in Go. This application allows users to manage log directories and files, including reading, writing, and organizing logs directly from the terminal.

## 📦 Features

- Create log directories
- Change working directory
- Print current working directory
- List contents of the current directory
- Create log files
- Read log files
- Write log entries

## ✅ Requirements

- Go (version 1.18 or higher)

## 🚀 Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/loid-lab/log_management_system.git
   cd log_management_system
   ```

2. **Build the application**:
   ```bash
   go build -o log-manager
   ```

3. **Run the application**:
   ```bash
   ./log-manager
   ```

## 🖥️ Usage

When you start the program, you'll be greeted with an interactive menu:

```
Welcome to the Log Manager!
Choose an operation:
1. Create Log Directory
2. Change Log Directory
3. Print Current Working Directory
4. List Directory Contents
5. Create Log File
6. Read Log File
7. Write Log Entry
8. Exit
```

Enter the number corresponding to the action you want to perform and follow the prompts.

## 📁 Project Structure

```
log_management_system/
├── main.go
├── go.mod
├── directory_operations.go
├── file_operations.go

```

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## 📄 License

This project is licensed under the MIT License.

## ✨ Author

**Loid Lab** – [GitHub Profile](https://github.com/loid-lab)
