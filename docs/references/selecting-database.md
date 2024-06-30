# Selecting a Database

### Connecting at runtime

There is flexibility over which database is used at runtime, just by using the `-dsn` command-line flag when starting the golang server

```zsh
go run cmd/web/main.go -dsn <USER:PASSWORD@/snippetbox?parseTime=true>
```

Where `USER` is a database use and `PASSWORD` is substitute for the password.

---

### Connecting to Development Database

---

### Connecting to QA Database Database

---

### Connecting to Production Database
