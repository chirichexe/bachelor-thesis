# Progetto k3s - IoT

Il progetto ha come sfida principale il Provisioning dinamico di Workload su dispositivi IoT, utilizzando Crossplane e operatori Kubernetes. La creazione di due controller: 
<!--- **Controller 1: Deployment di Workload su dispositivi IoT**!--->
- **Controller 2: Un admin gestisce un pool di dispositivi IoT da assegnare a un namespace**

1. Creazione del [Namespace](src/namespace.yaml) per i dispositivi IoT
2. Creazione di una CRD che verrà utilizzata per richiedere i dispositivi IoT
3. Creazione di un claim per richiedere un dispositivo IoT