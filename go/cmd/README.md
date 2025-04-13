## Report cmd

I want to build out a report command, that will:
 - run tests (if possible; might need to do this from os/exec, capture output and parse ...)
 - run benchmarks (tbd; need to write and see how I can access)
 - spit out a nice table broken down by package name


This will change the workflow, to:
 - solve exercise, proven with tests
 - write benchmark method
 - run report command to get the table output
 - (optional) maybe have report write that to a section of README.md


### Approach

- Use cobra to turn this repo into a CLI tool
- Set up run command to run tests
- Set up report command to rnu tests, capture output, programmatically access and format.
