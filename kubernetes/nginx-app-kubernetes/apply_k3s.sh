echo "Status: ------------------------------------------------------------------- "
echo " "
sudo kubectl apply -f nginx-deploy.yaml
sudo kubectl apply -f nginx-service.yaml
sudo kubectl apply -f nginx-ingress.yaml
echo " "

echo "Nodes: ------------------------------------------------------------------- "
echo " "
sudo kubectl get nodes -o wide
echo " "

echo "Pods: -------------------------------------------------------------------- "
echo " "
sudo kubectl get pods
echo " "

echo "Deployments: ------------------------------------------------------------- "
echo " "
sudo kubectl get deployments
echo " "

echo "Services: ---------------------------------------------------------------- "
echo " "
sudo kubectl get services
echo " "

echo "Curl della risorsa: ------------------------------------------------------ "
echo " "
curl -I http://nginx.local
echo "-------------------------------------------------------------------------- "


