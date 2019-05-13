#!/usr/bin/env bash

echo
TestStep_0() {
echo "----- Test Step - 0 -----"
echo "TEST STEP - 0 "
echo "API NAME: basic web app"
echo "URL: http://localhost:8080/"

response_code=$(curl -XGET -i -k --write-out %{http_code} --output /dev/null http://localhost:8080/)

if [ $response_code = "200" ]; then
    echo "STEP - 0: PASS"
    echo "----- --- -----"
    return 0
else
  echo "STEP - 0: FAIL"
  return 1
fi
}
echo "----- --- -----"
echo

echo
TestStep_1() {
echo "----- Test Step - 1 -----"
echo "TEST STEP - 1 "
echo "API NAME: basic web app"
echo "URL: localhost:8080/users"

response_code=$(curl -XPOST -i -k --write-out %{http_code} --output /dev/null localhost:8080/users)

if [ $response_code = "" ]; then
    echo "STEP - 1: PASS"
    echo "----- --- -----"
    return 0
else
  echo "STEP - 1: FAIL"
  return 1
fi
}
echo "----- --- -----"
echo




TEST_PASS=0
TEST_FAIL=0
TOTAL_TEST=0

declare -a arr=("TestStep_0" "TestStep_1" )

for i in "${arr[@]}"
do
    if $i; then
        TEST_PASS=$((TEST_PASS+1));
    else
        TEST_FAIL=$((TEST_FAIL+1));
fi
done

echo
echo

echo "--- TEST CASE REPORT ---"
echo "TEST PASS: " $TEST_PASS
echo "TEST FAIL: " $TEST_FAIL
echo "TOTAL EXECUTED: " ${#arr[@]}
echo "--- --- --- --- --- ---"