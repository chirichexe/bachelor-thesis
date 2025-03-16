# CROSSPLANE
Crossplane è un'estensione open-source per Kubernetes che permette al cluster di supportare il provisioning e la gestione di infrastrutture cloud, trasformando il cluster in uno **universal control plane** , ovvero il controllo continuo che la propria risorsa esista e funzioni come desiderati. È quello che faceva anche Kubernetes, ma in questo caso viene esteso a tutte le risorse.

Qualsiasi risorsa che usa un API è messa sotto lo stesso tetto. Crossplane si comporta come un classico **Controller Kubernetes**, quindi se qualcosa modifica o elimina a risorsa esterna, Crossplane si occupa di riportare lo stato a quello desiderato.

Consente ai team di piattaforma di definire e offrire API per l'infrastruttura in modalità self-service ai team di sviluppo, sfruttando strumenti e pratiche nativi di Kubernetes.

## Funzionalità chiave

- **Infrastructure as Code (IaC)**: Definisci e gestisci l'infrastruttura utilizzando il manifest YAML Kubernetes.
- **Modularità**: Definisci e gestisci l'infrastruttura utilizzando manifest Kubernetes.
- **Estendibilità**: Estendi Crossplane con Custom Resource Definitions (CRD) e controller personalizzati.
- **Multi-Cloud**:  Gestisci risorse su più provider cloud da un unico control plane universale.

Tutto ciò che devo fare è applicare il manifest, dimenticarmi delle API.

## Installazione
```sh
helm install crossplane crossplane-stable/crossplane --namespace crossplane-system --create-namespace

# Lo troverò poi disponibile con:
kubectl get pods -n crossplane-system
```

Visualizzo i due pod di crossplane:
- ```crossplane```
È il controller principale che gestisce le risorse (core). Installa i provider e orchestra le risorse monitorando gli Object creati. È lui che controlla continuamente lo stato delle risorse.

- ```crossplane-rbac-manager```
Gestisce i permessi, controlla che i controller abbiano le autorizzazioni per interagire con il cluster. 

[Documentazione sui pod](https://docs.crossplane.io/v1.19/concepts/pods/)


## Custom Resource Definition
![image](/assets/crossplane-components.png)
Componenti crossplane

I CRDs (Custom Resource Definition) Rappresentano la risorsa esterna come elemento nativo di Kubernetes, così da poter usare le classiche api ```kubectl create```, ```describe```...

### 1. Provider

Poichè Crossplane non è gestito direttamente da kubectl, ma usa le API di Kubernetes per installare i provider, creo una risorsa **Provider** mediante il file [provider-kubernetes.yaml](./nginx-app-crossplane/provider-kubernetes.yaml), poi la applico.
Ora Crossplane sa come interagire con il cluster Kubernetes.

### 2. ProviderConfig
Per dare l'accesso al cluster locale

```sh
# Creo un service account per crossplane e gli do i permessi di cluster-admin 

kubectl create namespace crossplane-system
kubectl create serviceaccount crossplane-admin -n crossplane-system
kubectl create clusterrolebinding crossplane-admin-binding --clusterrole=cluster-admin --serviceaccount=crossplane-system:crossplane-admin

# Ottengo il token per il service account
SECRET_NAME=$(kubectl get serviceaccount crossplane-admin -n crossplane-system -o jsonpath='{.secrets[0].name}')
TOKEN=$(kubectl get secret $SECRET_NAME -n crossplane-system -o jsonpath='{.data.token}' | base64 --decode)
```

Creo il file [provider-config-kubernetes.yaml](./nginx-app-crossplane/provider-config.yaml), fornisce le credenziali e la configurazione necessaria per connettersi al cluster Kubernetes.

Creo il secret contenente il kubeconfig:
```sh
kubectl create secret generic crossplane-admin-secret -n crossplane-system \
  --from-literal=kubeconfig="$(kubectl config view --raw)"
```
Poi applico il provider-config.

### 3. Creazione manifest

Creo il file [nginx-deployment-crossplane.yaml](./nginx-app-crossplane/nginx-deployment-crossplane.yaml), che rappresenta il deployment di nginx, e lo applico.





## Comandi utilizzati

Crossplane provides a command-line interface (CLI) to interact with and manage Crossplane resources. Below are some of the key CLI commands and their meanings:

- **`kubectl crossplane install`**: Installs Crossplane into your Kubernetes cluster.
- **`kubectl crossplane uninstall`**: Uninstalls Crossplane from your Kubernetes cluster.
- **`kubectl crossplane provider install <provider>`**: Installs a specific cloud provider into Crossplane.
- **`kubectl crossplane provider uninstall <provider>`**: Uninstalls a specific cloud provider from Crossplane.
- **`kubectl crossplane configuration package install <package>`**: Installs a configuration package into Crossplane.
- **`kubectl crossplane configuration package uninstall <package>`**: Uninstalls a configuration package from Crossplane.
- **`kubectl crossplane composition create -f <file>`**: Creates a new composition from a YAML file.
- **`kubectl crossplane composition delete <name>`**: Deletes an existing composition by name.
- **`kubectl crossplane resource claim create -f <file>`**: Creates a new resource claim from a YAML file.
- **`kubectl crossplane resource claim delete <name>`**: Deletes an existing resource claim by name.

These commands help you manage Crossplane installations, providers, configurations, compositions, and resource claims efficiently.

# Documentazione consultata
- [official Crossplane documentation](https://crossplane.io/docs/).
- https://www.youtube.com/watch?v=AtbS1u2j7po&list=PLyicRj904Z9_X62k6_XM_xlJkSyoQDkS2
kubectl get pods -n crossplane-system
