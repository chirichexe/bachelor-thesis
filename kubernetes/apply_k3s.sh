sudo kubectl apply -f nginx-deploy.yaml
sudo kubectl apply -f nginx-service.yaml

echo "------------------------------------------------------- "
echo "Status: "
sudo kubectl get pods
sudo kubectl get deployments
sudo kubectl get services

echo "------------------------------------------------------- "
echo "Provo a raggiungere la risorsa..."
curl http://localhost:30080
echo "------------------------------------------------------- "


