# A Git repository monitor , status all your repo in one command

# Install
```shell
go get github.com/StevenZack/gs
```

# Usage

## Add repository you want to monitor into the configure file

```shell
gs -a ./mygitrepo
```

```shell
gs -a ~/go/src/github.com/StevenZack/*
```

## Run `git status` in all your repository
```shell
gs
```

## Run `git pull` in all your repository
```shell
gs -p
```

## Remove repository in configure file
```shell
gs -r ./mygitrepo
```
