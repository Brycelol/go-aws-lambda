#!/usr/bin/env bash

BUILD_DIR="`pwd`/build"

echo "Building functions..."

echo "Cleaning build directory $BUILD_DIR..."
rm -rf "$BUILD_DIR"

for file in cmd/functions/*/*; do
    echo "Found Lambda Function File at $file..."
    func_dir="$(dirname $file)"
    func_build_dir="$BUILD_DIR/$func_dir"
    mkdir -p "$func_build_dir"

    echo "Compiling Lambda function to binary..."
    GOOS=linux go build -o "$func_build_dir/bootstrap" "$file"

    zip_output="$func_build_dir/lambda.zip"
    echo "Zipping up Lambda Function to $zip_output..."
    zip "$zip_output" "$func_build_dir/bootstrap"
done

echo "Finished!"