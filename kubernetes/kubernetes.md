# KUBERNETES
Orchestratore di container, design: accoppiamento debole tra container, indipendente dall'infrastruttura.
> Può fare:
> - Storage discovery e load balancing
> - Autoscaling, rollback e rollout automatici
> - Orchestrazione
> Non può fare
> - Fornire architettura a livello applicazione
> - Deployare / Buildare un'applicazione
 

- Repository: https://github.com/kubernetes/kubernetes
- Componenti: https://kubernetes.io/docs/concepts/overview/components/

Struttura principale: [ Cluster [ Node [ Pod ] ] ]

# K3S (versione consigliata dal professore per IoT)
Versione leggera di kubernetes con poco overhead.
## Architettura:
### k3s Server
equivalente al Control Plane di Kubernetes, crea un processo con **SQLite DB**

Ha un tunnel proxy per comunicare con ...

### k3s Agent
Esegue i container sui nodi worker
- c'è un kublet che comunica con ...
- containerd utilizzato al _posto di docker_ per gestire dei Pod[1 ... N]
- 

## API k3s
L'API sever espone una REST API

Il "desired state" è espresso tramite un file YAML

Il CLI () di kubernetes è chiamato kubectl comunica con l'API server e le sue informazioni di connessione
sono nella cartella ```~/.kube/config``` 

Il "contesto" è un grupo di parametri di accesso a un cluster k3s. 
Contiene un cluster k3s, utente e un namespace.


## Installazione e configurazione locale di k3s

```sh
curl -sfL https://get.k3s.io | sh - # installa 
kubectl get nodes                   # visualizza nodi attivi
```
## Creazione prima applicazione

- Creiamo ora una prima applicazione, creando un file ```nginx-deploy.yaml```
- Dopodichè la applico con:
```sh
kubectl apply -f nginx-deploy.yaml
kubectl get pods    # Ottengo i due pod attivi che avevo configurato nel file yaml
```
- Per poter esporre un pod all'esterno devo creare un service, creando un file ```nginx-service.yaml```, poi lo applico
```sh
kubectl apply -f nginx-service.yaml
kubectl get service    # Ottengo i servizi per i pod attivi con anche un ip e porta associati
```
## Esporre il servizio
Possiamo conoscere l'IP del servizio creato di nome **nginx-service** con
```sh
kubectl get svc nginx-service
```
L'output è:
```console
NAME            TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
nginx-service   LoadBalancer   10.43.177.157   <pending>     80:31066/TCP   2m39s
```
In Kubernetes, un servizio di tipo **LoadBalancer** crea automaticamente un bilanciatore di carico esterno per instradare il traffico ai pod interni.


L'external IP è in pending perchè k3s non ha un LoadBalancer esterno configurato, ovvero non c'è nulla che fornisce un IP pubblico (ricorda che k3s è progettato per piccoli server / IoT)

*Soluzioni*:
1. Usare un LoadBalancer locale con klipper-lb
2. Cambiare il servizio in NodePort per accedere direttamente al nodo
```kubectl port-forward svc/nginx-service 8080:80``` -> localhost:8080 mi fa vedere la pagina di default di Nginx 
3. Usare un ingress controller 

# Bibliografia

- docker: video 
- k3s: corso ufficiale rancher.academy
https://www.youtube.com/watch?v=QDwhbMvikGQ

