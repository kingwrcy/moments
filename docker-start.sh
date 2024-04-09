mkdir -p ${pwd}/moments
docker run --name moments -d -v ${pwd}/moments:/app/data -v ${pwd}/moments/upload:/app/.output/public/upload -p 3000:3000 moments