Go Profile-ify
==============

A script that automates the profiling of Go programs. It does all of the following:

- Installs [Dave Cheney](https://twitter.com/davecheney)'s [profile package](https://github.com/pkg/profile)
- Adds instrumentation code in your main function
- Runs your program
- Collects the pprof file containing profiling info
- Runs the go pprof program, to generate a list of functions sorted by execution time

Can be modified to output other info that the pprof tool supports, or just
allow more configuration options in general.

Installation
------------

```bash
mkdir ~/opt
cd ~/opt
git clone https://github.com/samuell/goprofileify.git
cd goprofileify
echo 'export PATH=~/opt/goprofileify/:$PATH' >> ~/.bashrc
source ~/.bashrc
```

Usage
-----
```
goprofileify examplegoapp.go
```
