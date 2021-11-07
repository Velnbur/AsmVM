# AsmVM
> My lab work for Computing Systems Architecture(CSA) subject in my university.
> Also was made as a simple example of small **Golang** _CLI_ tool

AsmVM - is a simple interpretation for nasm kind of files. 
Now it has eight registers from `R1` to `R8` and two operators
for manipulating them `mov` and `sub`.

For example, `mov R2, 2` will save `2` in `R2` register. And `sub R3, 1`
will subtract value in `R3` and `1`. Result of `sub` saves in accumulator
register - `R1`. You can use `example.iasm` as a remainder

## Prerequisites

- golang >=1.17.2

## Build and Install
Clone repository:
> git clone TODO

Go to src dir:
> cd AsmVM/src/

Build:
> go build -o iasm

Use:
> ./iasm <path_to_file>

