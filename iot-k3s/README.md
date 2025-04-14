# Progetto k3s - IoT

Il progetto ha come sfida principale il Provisioning dinamico di Workload su dispositivi IoT, utilizzando Crossplane e operatori Kubernetes.

## Architettura
1. Creazione di una macchina virtuale Vagrant per il nodo master del k3s e di una per il nodo worker. Ho utilizzato IP statici per facilitare la comunicazione tra essi.

2. Creazione di un namespace **iotdevices** per i dispositivi IoT 

3. Creazione di un'immagine docker per i pod di lavoro, che contengono un'applicazione (per ora test) in Go per la gestione dei dispositivi IoT. L'immagine è stata creata con Dockerfile e successivamente caricata su un repository Docker tramite i comandi:

```sh
GOOS=linux GOARCH=amd64 go build -o main . # compilazione dell'app

docker build -t chirichexe/device-agent:latest .
docker push chirichexe/device-agent:latest
```

4. Definizone di una **CustomResourceDefinition** per descrivere i dispositivi IoT. I parametri principali sono:

- **ip**: un campo stringa per l'indirizzo IP.
- **status**: uno stato che può essere "available", "assigned", "offline" o "errored".
- **lastStatusChange**: una data in formato date-time che indica l’ultima modifica dello stato.
- **assignedNamespace**: il nome del namespace a cui il dispositivo è assegnato.
- **capabilities**: un array di stringhe che descrivono le capacità del dispositivo.
- **lastSeen**: la data e ora dell’ultimo contatto con il sistema.
- **expirationTime**: la data e ora in cui il dispositivo scadrà o non sarà più considerato valido.
[...]

5. Definizione di un Deployment statico che crea due pod che eseguono l'immagine Docker.

```sh
kubectl rollout restart deployment iot-devices-deployment # per riavviare il deployment
```

6. Definizione di un Service di Kubernetes, ( ovvero un'astrazione per esporre i Pod nella rete ) 

Comandi utili:
```sh
kubectl get pods -l app=iot-device -o wide
kubectl get svc iot-device-service
kubectl get nodes -o wide

kubectl exec -it <nome-pod> -- sh

```

4. Creazione di un API Rest per esporre funzionalità di:
- **registrazione** di un dispositivo IoT.
- **deregistrazione** di un dispositivo IoT.
[...]

5. Sviluppo dell'operatore Kubernetes

l’Operatore sarà composto da due controller, ciascuno con responsabilità distinte:   
- **Controller Dispositivo** (Workload Manager), l'obiettivo è di monitorare il carico e scalare autonomamente
    - Osserva risorse IoTDevice
    - Crea e riconcilia Pod o Deployment per il workload associato al dispositivo

        i Pod sono agenti software che girano su nodi "worker" e rappresentano l’interfaccia logica tra il cluster e i dispositivi
        
        obiettivi: Connettersi al dispositivo, rispondere a eventi esterni, esporre metriche e stato
    
    - Crea un deployment da HPA (Horizontal Pod Autoscaler) per il carico (oppure analisi di metriche da Prometheus, scalando in base al numero di richieste)
- **Controller Admin**
    - Tiene traccia dell’inventario di dispositivi
    - Gestisce assegnazioni a namespace
    - Aggiorna lo stato dei dispositivi

## Flow operativo

1. Admin registra dispositivo -> API crea CR IoTDevice
2. Controller IoTDevice:

    - Legge spec
    - Crea Deployment + ConfigMap
    - Crea HPA (?)

3. Pod parte e si connette al dispositivo fisico
4. Se il carico cresce:

    - HPA scala i pod oppure
    - Controller aggiorna spec.replicas dinamicamente

5. Controller admin aggiorna status del dispositivo

<!--- **Controller 1: Deployment di Workload su dispositivi IoT**!--->
- **Controller 2: Un admin gestisce un pool di dispositivi IoT da assegnare a un namespace**

1. Creazione del [Namespace](src/namespace.yaml) per i dispositivi IoT
```kubectl get namespaces``` per visualizzarli
2. Il namespace esiste. Ora definisco un [CRD](src/iotdevices-crd.yaml) utilizzato per descivere un oggetto "iotDevice" collegato al namespace. In questo modo introduco il concetto di "dispositivo IoT" come risorsa di Kubernetes gestibile tramite kubectl. i dispositivi hanno:


Elenco di comandi eseguibili:
```sh
kubectl get iotdevices -n iot-devices # per avere una visione d'insieme.
kubectl delete iotdevice <nome> -n iot-devices # eliminazione dispositivo dal namespace
kubectl describe iotdevice <nome> -n iot-devices # per dettagli approfonditi.
kubectl get events -n iot-devices # per monitorare gli eventi.
kubectl get pods -n iot-devices e kubectl logs <pod> -n # iot-devices per controllare lo stato e i log dei workload connessi.
```

Aggiornamento tramite kubectl

```kubectl patch iotdevice device-001 -n iot-devices --type='merge' -p '{"spec": {"status": "assigned"}}'```

Cambio lo stato ad "assigned"

. Creazione del controller in python, che segue i tre passaggi: 1. Watch, 2. Analyze, 3. Act. Il controller è in grado di gestire i dispositivi IoT e di assegnarli a un namespace.

```sh
docker build -t myrepo/iot-controller:latest .
docker push myrepo/iot-controller:latest
```

## DA STUDIARE
- RBAC per gestire gli accessi
- Crossplane