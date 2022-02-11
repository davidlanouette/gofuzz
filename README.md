# Fuzzing in Go 1.18

## What is it?

"Fuzzing is a type of automated testing which continuously manipulates inputs to a program to find bugs. Go fuzzing uses coverage guidance to intelligently walk through the code being fuzzed to find and report failures to the user. Since it can reach edge cases which humans often miss, fuzz testing can be particularly valuable for finding security exploits and vulnerabilities."

## How does it work?

Demo time!

## When does it make sense for you?

### Good uses
Testing corner cases that you never thought of.

### Limitations
You can't control the input, so it can be hard to test what should be an error, and what shouldn't.



## Appendix

Tutorial
https://tip.golang.org/doc/tutorial/fuzz

Installing go 1.18
https://tip.golang.org/doc/tutorial/fuzz#installing_beta

Slides
https://docs.google.com/presentation/d/1v4Suf9dVbDZkjQsLMSrIX_MPXwFSXUyCyLv3sbjc968/edit?usp=sharing
