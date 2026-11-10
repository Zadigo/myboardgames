## Commands

### Testing

```Shell
go test -coverprofile=.coverage.html -coverpkg=./... ./...


# Treemap
# Url to visualize: https://nikolaydubina.github.io/go-cover-treemap/ (use .coverage.html)

go install github.com/nikolaydubina/go-cover-treemap@latest

go-cover-treemap -coverprofile .coverage.html > tree-map-out.svg


# If error: zsh: command not found: go-cover-treemap
# then run the command again 

echo 'export GOPATH=$HOME/go' >> ~/.zshrc

echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc

source ~/.zshrc
```
