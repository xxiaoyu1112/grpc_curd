source build_all.sh
ports=(50001 50002 50003 50004)

modules=("create_data" "get_data" "update_data" "delete_data")

echo "stop old service"

for port in "${ports[@]}"
do
pIDa=`lsof -i :$port | grep -v "PID" `
if [ "$pIDa" != "" ];
then
   fuser -k -n tcp $port   
   echo "close port: $port"
fi
 
done


for module in "${modules[@]}"
do
    $REPO_PATH/$module/out/$module &
   # do whatever on "$i" here
done