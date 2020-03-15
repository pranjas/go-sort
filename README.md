# go-sort
A collection of sort algorithms written in go as I learn
more about Golang!

About go modules,
In order to manage code it's important to keep related files together,
therefore I used go mod init to do this.

At first there were errors with go mod init, I tried the following
* go mod init
* go mod init $(pwd) => Says something about empty path (No idea wtf that means)

then I just decided to give it a name so the following worked
* go mod init ~sort~ **pks_sort**
The above package seems to exist already. So changed the name to **pks_sort**.

I also created a new directory to hold some random generator functions.

** It's important to note that you need to name the package to be name of the directory
   in which the go file exists. And there shouldn't be any file containing "main"
   package as that causes an error while importing package which says that package is
   a binary not an importable module. **

* Removed sort_main.go which was actually an executable containing each algo's run.
* Added benchmark functions for each algorithm.

To run bechmarks, use the following
**go test -bench=.**
