# justgo
My collection of go projects


## Setting up Go

Install `git` if you have not already done so.  
```sh
sudo apt install git
```

Most of the instructions in our READMEs apply to Linux, the operating system of choice to run and deploy open source applications.  
Follow the instuctions to download and install Go for your specific operating system.  

https://golang.org/doc/install

## Code organization

### Workspace  
- A workspace contains many version control repositories (managed by Git, for example).
- Each repository contains one or more packages.
- Each package consists of one or more Go source files in a single directory.
- The path to a package's directory determines its import path.

A workspace is a directory hierarchy with two directories at its root:  
- src contains Go source files, including those you `go get`  
- bin contains executable commands  
- pkg contains the libs you `go get`  

Follow the guidelines mentioned in:  
https://golang.org/doc/gopath_code.html

### GOPATH  
The default GOPATH is $HOME/go on Unix.  
If you are setting your GOPATH, follow the guidelines mentioned in above link as well.  

## Git
Below are basic git commands to create your GitHub project for subsequent collaboration:  

```sh
cd ...to your src directory
$ git init
Initialized empty Git repository in /home/...your src directory/.git/
```

Update .git/config file with:
```sh
$ vi .git/config
```

Enter your github email and name accordingly:  
```config
[core]
	repositoryformatversion = 0
	filemode = true
	bare = false
	logallrefupdates = true
[remote "origin"]
	url = https://github.com/.../....git
	fetch = +refs/heads/*:refs/remotes/origin/*
[user]
	email = <your email used in github>
	name = <your github username>
```

Pull latest source codes from GitHub:  
```sh
$ git fetch
Username for 'https://github.com': ...  
Password for 'https://...@github.com': ...
remote: Enumerating objects: 115, done.
remote: Counting objects: 100% (115/115), done.
remote: Compressing objects: 100% (85/85), done.
remote: Total 115 (delta 17), reused 115 (delta 17), pack-reused 0
Receiving objects: 100% (115/115), 1.75 MiB | 1.21 MiB/s, done.
Resolving deltas: 100% (17/17), done.
From https://github.com/...
 * [new branch]      master     -> origin/master

$ git pull origin master
Username for 'https://github.com': ...
Password for 'https://...@github.com': ...
From https://github.com/...
 * branch            master     -> FETCH_HEAD
```





