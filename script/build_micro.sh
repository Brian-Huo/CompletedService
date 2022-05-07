base_path=$1 #"./service/address"
echo "work in $base_path"

cd $base_path
rm -rf model/
rm -rf api/
rm -rf rpc/

# build sub folders
mkdir model
mkdir api
mkdir rpc

# move files to folder
cp *.sql model/
cp *.api api/
cp *.proto rpc/

# excute goctl for model/api/rpc
echo "excute model goctl ..."
goctl model mysql ddl -src ./model/*.sql -dir ./model -c
echo "excute api goctl ..."
goctl api go -api ./api/*.api -dir ./api -style gozero
echo "excute proto goctl ..."
goctl rpc protoc ./rpc/*.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc