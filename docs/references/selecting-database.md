# Selecting a Database

### Connecting at runtime

There is flexibility over which database is used at runtime, just by using the `-dsn` command-line flag when starting the golang server

```zsh
go run cmd/web/main.go -dsn <USER:PASSWORD@/snippetbox?parseTime=true>
```

Where `USER` is a database use and `PASSWORD` is substitute for the password.

A quirk of our MySQL driver is that we need to use the parseTime=true parameter in our DSN to force it to convert TIME and DATE fields to time.Time. Otherwise it returns these as []byte objects.

---

### Connecting to Development Database

---

### Connecting to QA Database Database

---

### Connecting to Production Database
