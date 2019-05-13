#!/usr/bin/env bash

TestStep_0() {
echo "----- Test Step - 0 -----"
echo "TEST STEP - 0 "
echo "API NAME: basic web app"
echo "URL: http://localhost:8080/"

response_code=$(curl -XGET -i -k --write-out %{http_code} --output /dev/null http://localhost:8080/)

if [ $response_code = "200" ]; then
    echo "TEST STEP - 0: PASS"
    return 0
else
  echo "TEST STEP - 0: FAIL"
  return 1
fi
echo "----- --- -----"
}

TEST_PASS=0
TEST_FAIL=0
TOTAL_TEST=0

## declare an array variable
declare -a arr=("TestStep_0" "TestStep_0" "TestStep_0")

## now loop through the above array
for i in "${arr[@]}"
do
    if $i; then
        TEST_PASS=$((TEST_PASS+1));
    else
        TEST_FAIL=$((TEST_FAIL+1));
fi
   # or do whatever with individual element of the array
done

# You can access them using echo "${arr[0]}", "${arr[1]}" also


echo "TEST PASS: " $TEST_PASS
echo "TEST FAIL: " $TEST_FAIL
echo "TOTAL EXECUTED: " ${#arr[@]}