# timestamping

```
"command | timestamping regexp [regexp] ..."
```

Act as a tee command except that when the given regexps are detected in a line, it show a timestamp and the difference with the previous timestamp mark or with the start of the execution.

It can be used to obtain the time consumed by each step of a command line tool that show it progress at the stdout.

# Notes
EOF not defected ok using go version prior to 1.5

(https://github.com/golang/go/commit/fd5b8aa7999e6710e14f4798dcb9e9387247511d)[Bugfix at 1.5]