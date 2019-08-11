## joke

> guaranteed to make your day better, or your money back!

## use

```bash
# installs joke in your $HOME/go/bin, make sure this directory is in your $PATH
$ go get github.com/gabrie30/joke
# sets up the sqlite database in $HOME/joke.db
$ joke db setup
# tell one joke
$ joke
# tell many jokes
$ joke --count=2
# get help
$ joke --help
```

> for best results add `joke` to .zshrc, .bashrc, etc.  :trollface:

## datastore

- SQLite uses the file `$HOME/joke.db` that is created on your behalf after running `joke db setup`. The database is only updated the first time you run the `joke` command, for that day. This is to increase performance.

> add the following to your `$HOME/.sqliterc` for easier to read queries
```
.mode column
.headers on
.separator ROW "\n"
.nullvalue NULL
```
then
```
$ sqlite $HOME/jokes.db
sqlite> select * from jokes;
```

## troubleshooting

- Make sure `$HOME/go/bin` is in your $PATH `go env | grep GOBIN` if not, you'll need to set it or put `$HOME/go/bin/joke` somewhere in your $PATH
