echo "Applying Crossplane resources--------------------------------"
echo " "
sudo kubectl apply -f provider-kubernetes.yaml
sudo kubectl apply -f provider-config.yaml
sudo kubectl apply -f nginx-deployment-crossplane.yaml
echo " "

echo "Crossplane resources applied successfully----------------------"
echo " "
sudo kubectl get deployments
echo " "

echo "Checking pods status: -----------------------------------------"
echo " "
sudo kubectl get pods
echo " "
