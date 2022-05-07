echo "work in "
pwd

cd ./service
# excute goctl for api
echo "excute api goctl ..."
goctl api go -api ./api/*.api -dir ./api -style gozero

# excute goctl for models
# for model_dir in `ls ./model`
# do
#     echo "excute model goctl - $model_dir ..."
#     goctl model mysql ddl -src ./model/$model_dir/*.sql -dir ./model/$model_dir -c
# done