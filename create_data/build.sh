# DELETE OUT
# shellcheck disable=SC2034
build_out=$(dirname $(readlink -f "$0"))/out/create_data
# delete output
# if [ ! -f "$build_out" ];then
#     rm "$build_out"
# fi
# BUILD
# shellcheck disable=SC2035
go build -o "$build_out"   *.go
