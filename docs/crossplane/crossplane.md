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
Un Provider comnnete Kubernetes a un servizio esterno. È responsabile della creazione e gestione del lifetime delle risorse esterne ad esso collegate, controllandole costantemente. 

**Installare un provider**: 

I provider hanno due tipi diversi di configurazione: **ControllerConfig** (deprecato) e **ProviderConfig**.

#### Test con Nginx

Creo una risorsa **Provider** mediante il file [provider-kubernetes.yaml](./nginx-app-crossplane/provider-kubernetes.yaml), poi la applico.
Ora Crossplane sa come interagire con il cluster Kubernetes. 

- Per ottenere i providers attivi: ```kubectl get providers```
- Per debug, ad esempio se un Provider è bloccato: ```kubectl describe providerrevisions```
- Eliminare: ```kubectl delete provider```

### 1.1 DeploymentRuntimeConfigs
Funzione beta, la analizzerò in seguito.

### 2. ProviderConfig
Determina le impostazioni che il Provider utilizza comunicando al Provider esterno. Caso d'uso: configurare le credenziali per accedere a un servizio esterno.

Esempio di ProviderConfig per AWS:
```sh
apiVersion: aws.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: user-keys
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: my-key
      key: secret-key

```

#### Test con Nginx
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

<!--
### 3. Creazione manifest

Creo il file [nginx-deployment-crossplane.yaml](./nginx-app-crossplane/nginx-deployment-crossplane.yaml), che rappresenta il deployment di nginx, e lo applico.
!-->

# 3. Managed Resources

Rappresenta un servizio esterno in un Provider. Quando ne creiamo una, 
> Meglio utilizzare un **CompositeResourceDefinition** per definire un servizio complesso, come un deployment.

# Documentazione consultata
- [official Crossplane documentation](https://crossplane.io/docs/).
- https://www.youtube.com/watch?v=AtbS1u2j7po&list=PLyicRj904Z9_X62k6_XM_xlJkSyoQDkS2
