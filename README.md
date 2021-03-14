# tool
A tool in golang which automates the entire process of collecting and organizing your code submissions
from Codeforces online-judge in one single Git repository.


## Install

If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):

```
go get -u github.com/shivangraina/cftool
```

## Usage

```
cftool [-c] [cfUsername] [-g] [githubUsername] [-n] [repositoryName]
```

## Example
```
cftool -c sam17 -g sam -n solutions
```
