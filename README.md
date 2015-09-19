# timestamping

```
"command | timestamping regexp [regexp] ..."
```

Act as a tee command except that when the given regexps are detected in a line, it show a timestamp and the difference with the previous timestamp mark or with the start of the execution.

It can be used to obtain the time consumed by each step of a command line tool that show it progress at the stdout.
