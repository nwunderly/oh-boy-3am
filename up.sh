git pull
docker build . --tag 3am:latest
docker run -d \
 --name 3am \
 --network prod \
 -v $PWD/data:/3am/data \
 --restart unless-stopped \
 3am --prod
