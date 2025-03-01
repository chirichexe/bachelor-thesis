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

# k3s (versione consigliata dal professore per IoT)
Versione leggera di kubernetes con poco overhead.
*Architettura*:
- k3s Server
equivalente al Control Plane di Kubernetes.

- k3s Agent
Esegue i container sui nodi worker

- containerd utilizzato al _posto di docker_ per gestire dei Pod[1 ... N]


# API k3s
L'API sever espone una REST API

Il "desired state" è espresso tramite un file YAML

Il CLI () di kubernetes è chiamato kubectl comunica con l'API server e le sue informazioni di connessione
sono nella cartella ```~/.kube/config``` 

Il "contesto" è un grupo di parametri di accesso a un cluster k8ù3s. 
Contiene un cluster k3s, utente e un namespace.


# Installazione locale di k3s

```sh
curl -sfL https://get.k3s.io | sh - 

```



