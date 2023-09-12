# HashBash
A tool to experiment with diffrent hash functions.

## Why
The tools are made to experiment with hash functions.
Online tools don't let you hash binary value, so it is hard 
to test diffrent points of intrest against each other

And the padding makes it so that some points of intrest ar not able to compare to each other.
If you want to compare 512 zeros to 1024 zeros the padding adds message length, so that the 
comparising isn't usefull
## How
The tools work by running go build /toolname/ then use ./toolname to run tool

## What
The different tools currently availible are:

|tool|use|
|----|-------|
| binary input | This tool takes binary input and gives a hash back, both in binary as in a hexidecimal represitation. |
| pad | this tool takes binary input and gives the padding that would be used in the sha256 hashfunction |
| zeros| gives a string of zeros with a user-chosen length|
| ones | gives a string of ones with a user-chosen length |
