Go Profile-ify
==============

A script that automates the profiling of Go programs. It does all of the following:

- Installs [Dave Cheney](https://twitter.com/davecheney)'s [profile package](https://github.com/pkg/profile)
- Adds instrumentation code in your main function
- Runs your program
- Collects the pprof file containing profiling info
- Runs the go pprof program, to generate a list of functions sorted by execution time

Can (and probably should) be modified to output other info that the pprof tool
supports, or just allow more configuration options in general.

The script is ugly and could use some polish, but it works!

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
```bash
goprofileify examplegoapp.go
```

Example usage session
---------------------

### View the example program before instrumentation

```bash
$ cat examplegoapp.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Writing 1M points:")
	for i := 0; i <= 1e6; i++ {
		fmt.Print(".")
	}
}
```

### Run the goprofileify script on the example program

```bash
$ goprofileify examplegoapp.go
github.com/pkg/profile
github.com/pkg/profile already installed!
Found temporary pprof file in /tmp/profile939290290/cpu.pprof
Wrote the slowest (cumulative) functions to perftest_20160819_161619_slowfuncs.log ...
Dropped 2 nodes (cum <= 0.01s)
      flat  flat%   sum%        cum   cum%
         0     0%     0%      2.81s   100%  main.main
         0     0%     0%      2.81s   100%  runtime.goexit
         0     0%     0%      2.81s   100%  runtime.main
     0.02s  0.71%  0.71%      2.68s 95.37%  fmt.Print
     0.05s  1.78%  2.49%      2.66s 94.66%  fmt.Fprint
     0.04s  1.42%  3.91%      2.32s 82.56%  os.(*File).Write
     0.07s  2.49%  6.41%      2.25s 80.07%  os.(*File).write
     0.01s  0.36%  6.76%      2.17s 77.22%  syscall.Write
     0.03s  1.07%  7.83%      2.16s 76.87%  syscall.write
     1.94s 69.04% 76.87%      2.13s 75.80%  syscall.Syscall
     0.03s  1.07% 77.94%      0.13s  4.63%  fmt.(*pp).doPrint
         0     0% 77.94%      0.13s  4.63%  runtime.convT2E
     0.05s  1.78% 79.72%      0.12s  4.27%  fmt.newPrinter
     0.05s  1.78% 81.49%      0.10s  3.56%  fmt.(*pp).printArg
     0.01s  0.36% 81.85%      0.10s  3.56%  runtime.entersyscall
     0.05s  1.78% 83.63%      0.09s  3.20%  runtime.exitsyscall
     0.02s  0.71% 84.34%      0.09s  3.20%  runtime.newobject
     0.07s  2.49% 86.83%      0.07s  2.49%  runtime.mallocgc
     0.03s  1.07% 87.90%      0.07s  2.49%  runtime.reentersyscall
     0.07s  2.49% 90.39%      0.07s  2.49%  runtime/internal/atomic.Cas
     0.03s  1.07% 91.46%      0.07s  2.49%  sync.(*Pool).Get
     0.03s  1.07% 92.53%      0.07s  2.49%  sync.(*Pool).pin
         0     0% 92.53%      0.05s  1.78%  runtime.casgstatus
     0.02s  0.71% 93.24%      0.05s  1.78%  runtime.typedmemmove
     0.01s  0.36% 93.59%      0.04s  1.42%  fmt.(*pp).fmtString
     0.01s  0.36% 93.95%      0.04s  1.42%  fmt.(*pp).free
     0.04s  1.42% 95.37%      0.04s  1.42%  runtime.memmove
     0.04s  1.42% 96.80%      0.04s  1.42%  sync.runtime_procPin
         0     0% 96.80%      0.03s  1.07%  fmt.(*fmt).fmt_s
     0.02s  0.71% 97.51%      0.03s  1.07%  os.epipecheck
     0.01s  0.36% 97.86%      0.03s  1.07%  runtime.exitsyscallfast
         0     0% 97.86%      0.03s  1.07%  sync.(*Pool).Put
     0.02s  0.71% 98.58%      0.02s  0.71%  fmt.(*fmt).truncate
     0.02s  0.71% 99.29%      0.02s  0.71%  runtime.assertI2T2
     0.02s  0.71%   100%      0.02s  0.71%  runtime.getcallerpc
```

### View the example program after instrumentation

Note the import of the github.com/pkg/profile package, and the defer statement
on the first line in `main()`.

```bash
$ cat examplegoapp.go
package main

import (
	"fmt"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	fmt.Println("Writing 1M points:")
	for i := 0; i <= 1e6; i++ {
		fmt.Print(".")
	}
}
```
