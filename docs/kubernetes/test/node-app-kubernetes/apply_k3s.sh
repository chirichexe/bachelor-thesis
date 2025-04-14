echo "Status: ------------------------------------------------------ "
echo " "
sudo kubectl apply -f node-deploy.yaml
sudo kubectl apply -f node-service.yaml
echo " "

echo "Pods: ------------------------------------------------------- "
echo " "
sudo kubectl get pods
echo " "

echo "Deployments: ------------------------------------------------ "
echo " "
sudo kubectl get deployments
echo " "

echo "Services: --------------------------------------------------- "
echo " "
sudo kubectl get services
echo " "

echo "------------------------------------------------------------- "
echo "Provo a raggiungere la risorsa..."
curl http://localhost:30080
echo "------------------------------------------------------------- "


