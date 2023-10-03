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

### Tools
|tool|use|
|----|-------|
| binary input | This tool takes binary input and gives a hash back, both in binary as in a hexidecimal represitation. |
| pad | this tool takes binary input and gives the padding that would be used in the sha256 hashfunction |
| zeros| gives a string of zeros with a user-chosen length|
| ones | gives a string of ones with a user-chosen length |

### Examples

Example|Use|Why|
|-------|---|-----|
| No skew hash| build with *go build NoSkewHash.go*| Aims to help prove *[A]* |
| Skew hash | build with *go build NoSkewHash.go*|Aims to help prove *[A]*  |
| Not all outputs | build with *go build NoSkewHash.go*|Aims to help prove *[A]*  |
| Hash.go| run with *go run Hash.go* | Baseline for custom SHA256 research implementations|

#### *A*
mb = message block
h = compression function 
H = hash function
O = Output
l = length
h(mb)=h(mb)
h(mb)^(Ol+1)≠h(mb)
proves 
H:A→B 
∃x ∈ [0,2^l)∶ x ∉ B
