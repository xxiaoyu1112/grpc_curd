#!/bin/bash
WS=$(pwd)

REPO_PATH=$(dirname $(readlink -f "$0"))

modules=("create_data" "get_data" "update_data" "delete_data")

for module in "${modules[@]}"
do
    echo $module
    cd $REPO_PATH/$module
    source build.sh 
   # do whatever on "$i" here
done
cd $WS
