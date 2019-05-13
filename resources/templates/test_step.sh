#!/usr/bin/env bash



TestStep_N() {
echo "#####################"
echo "----- Test Step - N -----"
echo "TEST STEP - N "
echo "API NAME: name"
echo "URL: url"

response_code=2001

if [ $response_code = "200" ]; then
    echo "pass"
   return 1
else
  echo "fail"
  return 0
fi
echo "----- --- -----"

#<<SCRIPT_PLACEHOLDER>>


}

TestStep_N