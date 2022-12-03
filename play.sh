#!/bin/zsh

cp test.yaml test.yaml.tmp
cp saved/$1.yaml test.yaml
go build
./pirsig
cp test.yaml.tmp test.yaml
rm test.yaml.tmp

ffplay -f f32le -ar 44100 -showmode 1 out.bin