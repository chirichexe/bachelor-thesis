sudo docker build -t node-image-test node_container/.
sudo docker stop node-container
sudo docker rm node-container
sudo docker run -d --name node-container -p 3000:3000 -v node-volume:/usr/src/app/logs node-image-test
sudo docker exec -it node-container sh
