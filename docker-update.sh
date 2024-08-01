basepath=$(cd `dirname $0`; pwd)
mkdir -p ${basepath}/moments

docker stop moments && docker rm moments
docker pull kingwrcy/moments:latest
docker run --name moments -e JWT_KEY=cfqYVP6CZm9mSqLVGlmL -d -v ${basepath}/moments:/app/data -p 3000:3000 kingwrcy/moments:latest