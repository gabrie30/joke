## joke

> guaranteed to make your day better, or your money back!

## use

```bash
# installs joke in your $HOME/go/bin, make sure this directory is in your $PATH
$ go get github.com/gabrie30/joke
# sets up the sqlite database in $HOME/.joke.db
$ joke db setup
# tell one joke
$ joke
# tell many jokes
$ joke --count=2
# get help
$ joke --help
```

> for best results add `joke` to .zshrc, .bashrc, etc.  :trollface:

## getting new jokes

New jokes are fetched only once per day, this is to increase performance which helps if joke is added to .zshrc or similar. However, you can fetch jokes manually by running `joke fetch`

## database

The datastore is an SQLite database located at `$HOME/.joke.db`. It's is created upon running `joke db setup`.

### viewing data

Open the database
```
$ sqlite3 $HOME/.jokes.db
sqlite> select * from jokes;
```

> TIP: add the following to your `$HOME/.sqliterc` for easier to read queries
```
.mode column
.headers on
.separator ROW "\n"
.nullvalue NULL
```

## troubleshooting

- Make sure `$HOME/go/bin` is in your $PATH `go env | grep GOBIN` if not, you'll need to set it or put `$HOME/go/bin/joke` somewhere in your $PATH
    - To add `$HOME/go/bin` to your path `export PATH="$PATH:$HOME/go/bin"`
