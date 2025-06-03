
<div align="center">
    <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.svg" alt="Go Gopher" width="120"/>
</div>

# I'm learning GO!

This repository contains all the files I'm using to learn Go.  
Nothing fancy... just a tutorial! 🚀

---

## 📦 About

This project is a collection of code samples, experiments, and tutorials as I explore the Go programming language. It's a work in progress and meant for learning purposes.

## 🛠️ Features

- Simple Go code examples
- API experiments
- Notes and learning resources

## 📚 Resources

- [Go Official Website](https://golang.org/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)

## Database Migrations

This project uses [golang-migrate](https://github.com/golang-migrate/migrate) to manage database schema changes. Install the CLI with:

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Migration scripts live in the `migrations` directory. The initial migration sets up the `books` and `users` tables with seed data.

