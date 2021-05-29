This repo is inspired by spf13's [Cobra](https://github.com/spf13/cobra). Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.
In this repo, I have created a simple Calculator (only addition and multiplication is implemented, you can add the rest :wink:) using Cobra.

- Clone this repository into desired directory using this command `git clone https://github.com/FedTheBug/calculator-cli.git`.
- Go into the directory where you have cloned the repo, and run `go mod vendor -v`
- After that, to get the executable binary file, run `go install calculator-cli`

You're finally done. 
To test the calculator: 
- To add numbers (integers)
```
    calculator-cli add 1200 500
```
Response should be like this
```
    Addition of [1200 500] is 1700
```
- To add numbers (floats)
```
    calculator-cli add 2.5 3.5 -f
```
Response should be like this
```
    Addition of [2.5 3.5] is 6.000
```
