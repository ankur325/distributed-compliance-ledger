#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail
source integration_tests/cli/common.sh

# Preparation of Actors

echo "Create regular account"
create_new_account user_account ""

vid=$RANDOM
pid=$RANDOM
vendor_account=vendor_account_$vid
echo "Create Vendor account - $vendor_account"
create_new_vendor_account $vendor_account $vid

# Body

# Ledger side errors

echo "Add Model with VID: $vid PID: $pid: Not Vendor"

result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$user_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 4"
echo "$result"

test_divider

vid1=$RANDOM
echo "Add Model with VID: $vid1 PID: $pid: Vendor ID does not belong to vendor"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid1 --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 4"
echo "$result"

test_divider

echo "Add Model with VID: $vid PID: $pid: Twice"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
check_response_and_report "$result" "\"success\": false"
check_response_and_report "$result" "\"code\": 501"
echo "$result"

test_divider

# CLI side errors

echo "Add Model with VID: $vid PID: $pid: Unknown account"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$vendor_account --yes)
result=$(dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from "Unknown"  2>&1) || true
check_response_and_report "$result" "Key Unknown not found"

test_divider

echo "Add model with invalid VID/PID"
i="0" 
result=$(echo "test1234" | dclcli tx model add-model --vid=$i --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Invalid VID"

result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$i --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Invalid PID"

i="string" 
result=$(echo "test1234" | dclcli tx model add-model --vid=$i --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Parsing Error: VID is set to 'string', but it must be 16 bit unsigned integer"

result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$i --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Parsing Error: PID is set to 'string', but it must be 16 bit unsigned integer"

test_divider

echo "Add model with empty name"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName="" --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid ProductName"

test_divider

echo "Add model with empty description"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="" --partNumber=1 --commissioningCustomFlow=0  --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid ProductLabel"

test_divider

echo "Add model with empty partNumber"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel="Test Label" --partNumber="" --commissioningCustomFlow=0  --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "Code: 6"
check_response_and_report "$result" "Invalid PartNumber"

test_divider

echo "Add model without --from flag"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0  --yes 2>&1) || true
check_response_and_report "$result" "required flag(s) \"from\" not set"

test_divider

echo "Add model without enough parameters"
result=$(echo "test1234" | dclcli tx model add-model --vid=$vid --pid=$pid --deviceTypeID=1 --productName=TestProduct  --partNumber="1" --commissioningCustomFlow=0  --from $vendor_account --yes 2>&1) || true
check_response_and_report "$result" "required flag(s) \"productLabel\" not set"
