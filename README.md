
# badgeify

Create a badge (SVG) by coverage results, or by custom color, label and value.

```sh
# Usage
$ badgeify > coverage.svg # find coverage results and create badge
$ badgeify --value "80%" > coverage-master.svg # provide value
$ badgeify --color "cc00cc" --label "works" --value "well" > well-done.svg

# Development
$ make test # run testsuite
$ make run # run in development
$ make # build executable
$ ./bin/badgeify
```
