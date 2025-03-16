echo "Deploment status: --------------------------------------------------------"
echo " "
sudo kubectl get deployments
echo " "
echo "Managed resources: -------------------------------------------------------"
echo " "
sudo kubectl get managed
echo "-------------------"
sudo kubectl get objects
echo " "

echo "Pods status: -------------------------------------------------------------"
echo " "
sudo kubectl get pods -A
echo " "

echo "Providers status: --------------------------------------------------------"
echo " "
sudo kubectl get providers
echo " "

echo "Deployment status: -------------------------------------------------------"
echo " "
sudo kubectl describe deployment nginx-deployment
echo " "

echo "Crossplane logs: ---------------------------------------------------------"
echo " "
sudo kubectl logs -n crossplane-system deployment/crossplane
echo " "
