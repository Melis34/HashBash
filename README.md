# HashBash
A tool to experiment with diffrent hash functions.

## Why
The tools are made to experiment with hash functions.
Online tools don't let you hash binary value, so it is hard 
to test diffrent points of intrest against each other

## How
The tools work by running go build /toolname/ then use ./toolname to run tool

## What
The diffrent tools currently availible are:

### binary input
This tool takes binary input and gives a hash back, both in binary as in a hexidecimal represitation.
### pad
this tool takes binary input and gives the padding back that the sha256 function will add before it runs
### zeros
This tool gives a string of all zeros