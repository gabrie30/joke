## joke

> guaranteed to make your day better, or your money back!

### use

```bash
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

> sqlite uses the file $HOME/joke.db that is created on your behalf after running `joke db setup`

## troubleshooting

- Make sure `$HOME/go/bin` is in your $PATH `go env | grep GOBIN` if not, you need to set it
