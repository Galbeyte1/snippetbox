# Connecting to Database

---

### Verify you have MySQL installed locally

```zsh
mysql --version
```

![mysql version](/docs/imgs/mysql-version.png)

---

### Download MySQL locally

```zsh
brew install mysql
```

Once you've installed and verified MySQL you can

### Start a MySQL server

```zsh
mysql.server start
```

To stop the MySQL server

```zsh
mysql.server stop
```

To restart the MYSQL server

```zsh
mysql.server restart
```

(DB User creation coming soon)

## Connecting Locally

The default is to connect as the web user

```zsh
mysql -D snippetbox -u web -p
```
