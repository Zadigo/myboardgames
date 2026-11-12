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

```Shell

# Profile CPU usage
go test -run=^$ -bench=BenchmarkGames -benchmem -cpuprofile=cpu.out -memprofile=mem.out

go tool pprof -top mem.out

# Save your current output to a text file...
go test -bench=. -benchmem > old.txt

# ... make a code optimization, save the new output
go test -bench=. -benchmem > new.txt

# compare them using Go's official benchstat tool (golang.org/x/perf/cmd/benchstat) 
# to verify whether improvements are statistically significant.

# 
go build -gcflags="-m" ./...
```
