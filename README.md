# Gotogo

Gotogo is a command-line tool for managing your to-do list. It checks all completed items and removes those older than one day, while keeping items completed today as encouragement.

![img](https://github.com/user-attachments/assets/99c32e7f-11c9-48d9-ac5e-b36487fb98eb)


## Features

- Automatically removes completed items older than one day
- Keeps items completed today

## Installation

To install Gotogo, use `go get`:

```sh
go install github.com/Kaya-Sem/gotogo@latest
```
Alternatively, you can download the latest release from GitHub:

## Configuration

Gotogo requires a CSV file to store your to-do items. 
This file should be located in your XDG configuration home directory under `.config/gotodo/todo.csv`.

### Example

On most Linux systems, this would be:

```sh
~/.config/gotodo/todo.csv
```

### CSV File Format

The CSV file should have the following headers:

```csv
id,title,completed,timestamp
```

An example of a to-do item:

```csv
1,Finish writing documentation,false,2024-10-01
```

The id, completed status and timestamp will be automatically generated when
you create new items through the CLI interface.

## Command Line Usage

to use Gotogo from the command line, you can run the following commands:


**List all to-do items:**

```sh
gotogo
```
**Add a new to-do item:**

```sh
gotogo add "Your new to-do item
```


**Mark a to-do item as completed:**

```sh
gotogo done <item_id>
```

## Contributing

1. Fork the repository
2. Create a new branch (git checkout -b feature-branch)
3. Commit your changes (git commit -am 'Add new feature')
4. Push to the branch (git push origin feature-branch)
5. Create a new Pull Request

# License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
