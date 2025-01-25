#! /bin/bash

# go to go.mod file and change this line: "replace github.com/Mau-MD/SSEngine/libs/proto => ../../libs/proto" to be "replace github.com/Mau-MD/SSEngine/libs/proto => ./output/libs/proto"
sed -i 's|replace github.com/Mau-MD/SSEngine/libs/proto => ../../libs/proto|replace github.com/Mau-MD/SSEngine/libs/proto => ./output/libs/proto|' ./go.mod
