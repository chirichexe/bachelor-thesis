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
### 1. Deployment
Un Deployment è una risorsa che gestisce il lifecycle di uno o più Pod in Kubernetes. Permette di:
- Creare e aggiornare gruppi di Pod.
- Gestire il numero di repliche (scalabilità).
- Assicurare che l’applicazione sia sempre disponibile.
- Fare rollback a una versione precedente se necessario.

#### Esempio
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: [nome-app]
  labels:
    app: [nome-app]
spec:
  replicas: [n-pods]  # 🔹 Numero di repliche dei Pod
  selector:
    matchLabels:
      app: [nome-app]  # 🔹 Kubernetes sa che questi Pod appartengono a questo Deployment
  template:
    metadata:
      labels:
        app: [nome-app]  # 🔹 Etichetta assegnata ai Pod per il Service
    spec:
      containers:
      - name: [nome-app]
        image: [nome-immagine]:[versione]  # 🔹 L'immagine da usare
        ports:
        - containerPort: [porta]  # 🔹 La porta che il container espone

```
> ***IMPORTANTE***
> Dato che k3s usa containerd invece di Docker, devi caricare l'immagine nel registry interno.

Kubernetes crea 3 Pod con l’immagine nginx:latest.
- Se un Pod si arresta per errore, Kubernetes lo ricrea automaticamente.
- Il Deployment si assicura che ci siano sempre 3 Pod attivi.

#### Prova
- Creiamo ora una prima applicazione, creando un file ```nginx-deploy.yaml```

- Dopodichè la applico con:
```sh
kubectl apply -f nginx-deploy.yaml
kubectl get pods    # Ottengo i due pod attivi che avevo configurato nel file yaml
```

### 2. Service

- Per poter esporre un pod all'esterno devo creare un service,  I Pod, infatti, hanno un IP volatile, quindi se uno viene riavviato cambia IP e l’applicazione potrebbe non trovarlo più.

#### Esempio
```yaml
apiVersion: v1
kind: Service
metadata:
  name: [nome-servizio]
spec:
  selector:
    app: [nome-servizio]
  ports:
    - protocol: TCP
      port: [port]
      targetPort: [port]
      nodePort: 30007  # 🔹 Kubernetes assegna una porta tra 30000-32767
  type: NodePort

```
#### Prova

- Creo un file ```nginx-service.yaml```, poi lo applico
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

**Soluzioni**:
1. Usare un LoadBalancer locale con klipper-lb
2. Cambiare il servizio in NodePort per accedere direttamente al nodo
```localhost:8080 ``` -> mi fa vedere la pagina di default di Nginx 
3. Usare un ingress controller 

## Chiamate da remoto
facendo ```ip a``` noto che il mio Ip Pubblico è

```
inet 192.168.178.42/24 brd 192.168.178.255 scope global dynamic noprefixroute wlo1
```
Perciò alla porta 30080 di quell'indirizzo trovo la pagina che ho hostato

### Load Balancing

