#!/bin/bash

echo "Current directory contents:"
ls -la

echo "Contents of go.mod before:"
cat go.mod

# Detect OS and use appropriate sed syntax
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    sed -i '' 's|../../libs/proto|./output/libs/proto|g' go.mod
else
    # Linux and others
    sed -i 's|../../libs/proto|./output/libs/proto|g' go.mod
fi

echo "Contents of go.mod after:"
cat go.mod
 