#!/bin/bash
# Set working directory
pushd `dirname "$0"` > /dev/null

# Rebuild
echo "Rebuilding..."
pushd .. > /dev/null
go build

# Clear previous output
popd > /dev/null
mkdir -p actual
rm -Rf actual/*.out

run_test () {
  ../toyrobot < input/$1.in > actual/$1.out
  if diff expected/$1.out actual/$1.out ; then
    echo Passed test: $1
  else
    echo Failed test: $1
    exit 1
  fi
}

# Test all the input files
for f in input/*.in; do
  input_file=${f##*/}
  run_test ${input_file%.*}
done

echo All tests passed.
popd > /dev/null
exit 0
