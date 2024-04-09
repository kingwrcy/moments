basepath=$(cd `dirname $0`; pwd)
mkdir -p ${basepath}/moments
docker run --name moments -d -v ${basepath}/moments:/app/data -v ${basepath}/moments/upload:/app/.output/public/upload -p 3000:3000 moments