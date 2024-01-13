# Go-Templates
Repository of common Go project layout templates used to generate new projects with the `go-new` tool.

### Getting started
Verify that you have installed [go new]() by first running:

`$ gonew` 

If it is already installed you should see something similar to:

```
usage: gonew srcmod[@version] [dstmod [dir]]
See https://pkg.go.dev/golang.org/x/tools/cmd/gonew.
```

Otherwise install **gonew** by running:

`$ go install golang.org/x/tools/cmd/gonew@latest`


### Usage
To use the templates here simply navigate to a parent directory you wish to clone the project/template in
and then run the following command:

`$ go new github.com/masonictemple4/go-templates/template-name github.com/exampleuser/modulename`


That's it, now `$ cd into modulename` and you're good to go.
