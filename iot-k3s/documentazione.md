# Progetto k3s - IoT

Il progetto ha come sfida principale il Provisioning dinamico di Workload su dispositivi IoT, utilizzando Crossplane e operatori Kubernetes. La creazione di due controller: 


<!--- **Controller 1: Deployment di Workload su dispositivi IoT**!--->
- **Controller 2: Un admin gestisce un pool di dispositivi IoT da assegnare a un namespace**

1. Creazione del [Namespace](src/namespace.yaml) per i dispositivi IoT
```kubectl get namespaces``` per visualizzarli
2. Il namespace esiste. Ora definisco un [CRD](src/iotdevices-crd.yaml) utilizzato per descivere un oggetto "iotDevice" collegato al namespace. In questo modo introduco il concetto di "dispositivo IoT" come risorsa di Kubernetes gestibile tramite kubectl. i dispositivi hanno:

- **ip**: un campo stringa per l'indirizzo IP.

- **status**: uno stato che può essere "available", "assigned", "offline" o "errored".

- **lastStatusChange**: una data in formato date-time che indica l’ultima modifica dello stato.

- **assignedNamespace**: il nome del namespace a cui il dispositivo è assegnato.

- **capabilities**: un array di stringhe che descrivono le capacità del dispositivo.

- **lastSeen**: la data e ora dell’ultimo contatto con il sistema.

- **expirationTime**: la data e ora in cui il dispositivo scadrà o non sarà più considerato valido.

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

3. Sviluppo del controller in Go. Esso verrà eseguito nei pod all'interno del cluster. Ogni pod avrà un container che eseguirà il controller. Il controller è responsabile della gestione dei dispositivi IoT e della loro assegnazione a un namespace specifico. Utilizza le API di Kubernetes per interagire con il cluster e monitorare lo stato dei dispositivi.

4. Utilizzo di crossplane per
- Creare composizioni che definiscono come un dispositivo iot deve essere configurato  nel cluster.

- Quando un utente richiede un dispositivo, Crossplane avvia automaticamente il suo provisioning. (es. se diventa errored, ne assegna un altro, se diventa offline avvia una procedura di riavvio o sostituzione).

- Riassegnazione automatica dei dispositivi inattivi e allocazione automatica delle risorse (se quelli attivi sono pochi, Crossplane ne crea di nuovi o libera quelli inattivi).

Cosa deve fare l'admin:
- Stabilire le regole di gestione

Cosa deve fare l'user
- Possibilità di claim di un dispositivo 
