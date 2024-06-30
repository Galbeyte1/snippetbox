# Connecting to Database

---

## Connecting Locally

### Verify you have MySQL installed locally

```zsh
mysql --version
```

![mysql version](docs/imgs/mysql-version.png)

---

### Download MySQL locally

```zsh
brew install mysql
```

Once you've installed and verified MySQL you can connect as a user

(DB User creation coming soon)

```zsh
mysql -D snippetbox -u web -p
```

Yu have total control over which database is used at runtime, just by using
the -dsn command-line flag

```zsh
go run cmd/web/main.go -dsn <USER:PASSWORD@/snippetbox?parseTime=true>
```

Where `USER` is a database use and `PASSWORD` is substitute for the password.
