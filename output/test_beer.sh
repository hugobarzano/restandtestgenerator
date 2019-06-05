#!/usr/bin/env bash


TestStep_1() {
echo "----- Test Step - 1 -----"
echo "TEST STEP - 1 "
echo "API NAME: Beer for you"
echo "URL: http://localhost:8080/beer"
echo "Scripting field..."
#<<SCRIPT_PLACEHOLDER>>

response_code=$(curl -XPOST -i -k "http://localhost:8080/beer" --write-out %{http_code} --output /dev/null -d  '{"grados":"los suficientes","ingradientes":"cosas buenas","name":"Alhambra","sabores":"ricoo"}' )

if [ $response_code = "201" ]; then
    echo "STEP - 1: PASS"
    echo "----- --- -----"
    return 0
else
  echo "STEP - 1: FAIL"
  return 1
fi

}




TestStep_2() {
echo "----- Test Step - 2 -----"
echo "TEST STEP - 2 "
echo "API NAME: Beer for you"
echo "URL: http://localhost:8080/beer"
echo "Scripting field..."
get=$(curl -sb -H "Accept: application/json" "http://localhost:8080/beer" | jq '.[0]._id')  
  export ID=$get

response_code=$(curl -XGET -i -k --write-out %{http_code} --output /dev/null http://localhost:8080/beer)

if [ $response_code = "200" ]; then
    echo "STEP - 2: PASS"
    echo "----- --- -----"
    return 0
else
  echo "STEP - 2: FAIL"
  return 1
fi

}




TestStep_3() {
echo "----- Test Step - 3 -----"
echo "TEST STEP - 3 "
echo "API NAME: Beer for you"
echo "URL: http://localhost:8080/beer"
echo "Scripting field..."
#<<SCRIPT_PLACEHOLDER>>

response_code=$(curl -XPUT -i -k "http://localhost:8080/beer/${ID//\"}" --write-out %{http_code} --output /dev/null -d  '{"grados":"los suficientes_Update","ingradientes":"cosas buenas_Update","name":"Alhambra_Update","sabores":"ricoo_Update"}' )

if [ $response_code = "200" ]; then
    echo "STEP - 3: PASS"
    echo "----- --- -----"
    return 0
else
  echo "STEP - 3: FAIL"
  return 1
fi

}




TestStep_4() {
echo "----- Test Step - 4 -----"
echo "TEST STEP - 4 "
echo "API NAME: Beer for you"
echo "URL: http://localhost:8080/beer"
echo "Scripting field..."
echo $ID

response_code=$(curl -XDELETE -i -k --write-out %{http_code} --output /dev/null http://localhost:8080/beer/${ID//\"})

if [ $response_code = "200" ]; then
    echo "STEP - 4: PASS"
    echo "----- --- -----"
    return 0
else
  echo "STEP - 4: FAIL"
  return 1
fi

}






TEST_PASS=0
TEST_FAIL=0
TOTAL_TEST=0

declare -a arr=("TestStep_1" "TestStep_2" "TestStep_3" "TestStep_4" )

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