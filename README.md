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
* go mod init sort

But all my files were still at top level therefore I moved the algos in a separate
algo directory leaving me with only sort_main.go in the top level directory.

I also created a new directory to hold some random generator functions.

** It's important to note that you need to name the package to be name of the directory
   in which the go file exists. And there shouldn't be any file containing "main"
   package as that causes an error while importing package which says that package is
   a binary not an importable module. **

Post that I was able to build sort_main without any issues.
