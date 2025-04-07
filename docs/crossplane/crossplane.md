# CROSSPLANE
Crossplane è un'estensione open-source per Kubernetes che permette al cluster di supportare il provisioning e la gestione di infrastrutture cloud, trasformando il cluster in uno **universal control plane** , ovvero il controllo continuo che la propria risorsa esista e funzioni come desiderati. È quello che faceva anche Kubernetes, ma in questo caso viene esteso a tutte le risorse.

Qualsiasi risorsa che usa un API è messa sotto lo stesso tetto. Crossplane si comporta come un classico **Controller Kubernetes**, quindi se qualcosa modifica o elimina a risorsa esterna, Crossplane si occupa di riportare lo stato a quello desiderato tramite un **riconciliation loop**.

Consente ai team di piattaforma di definire e offrire API per l'infrastruttura in modalità self-service ai team di sviluppo, sfruttando strumenti e pratiche nativi di Kubernetes.

## Funzionalità chiave

- **Infrastructure as Code (IaC)**: Definisci e gestisci l'infrastruttura utilizzando il manifest YAML Kubernetes.
- **Modularità**: Definisci e gestisci l'infrastruttura utilizzando manifest Kubernetes.
- **Estendibilità**: Estendi Crossplane con Custom Resource Definitions (CRD) e controller personalizzati.
- **Multi-Cloud**:  Gestisci risorse su più provider cloud da un unico control plane universale.

Tutto ciò che devo fare è applicare il manifest, dimenticarmi delle API.

## Architettura

[ Management Kubernetes Cluster [ Crossplane [ Core ] [ Provider -> (comunica con l'esterno) ] ]]

## 0.Installazione e Crossplane pods
```sh
helm install crossplane crossplane-stable/crossplane --namespace crossplane-system --create-namespace

# Lo troverò poi disponibile con:
kubectl get pods -n crossplane-system
```

Visualizzo i due pod di crossplane che si occupano di gestire tutti i componenti e le risorse:
- ```crossplane```
È il controller principale che gestisce le risorse (core). Installa i provider e orchestra le risorse monitorando gli Object creati. È lui che controlla continuamente lo stato delle risorse.

- ```crossplane-rbac-manager```
Gestisce i permessi, controlla che i controller abbiano le autorizzazioni per interagire con il cluster. 

[Documentazione sui pod](https://docs.crossplane.io/v1.19/concepts/pods/)


## Custom Resource Definition
![image](/assets/crossplane-components.png)
Componenti crossplane

I CRDs (Custom Resource Definition) Rappresentano la risorsa esterna come elemento nativo di Kubernetes, così da poter usare le classiche api ```kubectl create```, ```describe```...

### 1. Providers
Un Provider comnnete Kubernetes a un servizio esterno. È responsabile della creazione e gestione del lifetime delle 
risorse esterne ad esso collegate, controllandole costantemente. 

Esso **traduce** le API esterne in API Kubernetes, permettendo di gestire le risorse esterne come se fossero native di Kubernetes.

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

### 3. Managed Resources

Rappresenta un servizio esterno in un Provider [...](https://docs.crossplane.io/v1.19/concepts/managed-resources/)
> Meglio utilizzare un **CompositeResourceDefinition** per definire un servizio complesso.

### 4. Composition

Architettura generale
![image](../../assets/composition-how-it-works.svg)

Un Composition è un template per creare più ManagedResources come un singolo oggetto.
Esso descrive deployment più complessi, combinando più ManagedResources e le loro customizzazioni
in un unico oggetto.

Due modalità di composizione:

- **Resources** (deprecata)
- **Pipeline**: definisce una pipeline di steps, ognuno dei quali chiama una funzione il cui elenco è
visualizzabile con ``` kubectl get functions```. QUando una composizione ha una pipeline, il controller di composizione esegue ogni step in sequenza.

Esempio di Composition con Pipeline:

```yaml
apiVersion: apiextensions.crossplane.io/v1.​
kind: Composition.​
metadata: Standardmetadata.​
spec:
    compositeTypeRef: References the XRD by specifying its apiVersion and kind.
    mode: Determines the composition mode, such as Pipeline.
    pipeline: Defines a sequence of functions to process the composition.
      step: Name of the step, e.g., patch-and-transform.
      functionRef: References the function to execute:
          name: Name of the function.
      input: Specifies input parameters:
          apiVersion: Function's API version.
          kind: Type of input, e.g., Resources.
          resources: List of resource templates: 
              name: Identifier for the resource.
              base: Base configuration of the resource, including apiVersion, kind, and spec.
```
## 5. CompositeResourceDefinition (XRD) 

Un XRD è un CRD che definisce una API custom (utilizzata da developers o utenti finali) per interagire con risorse esterne. 

Esse utilizzano un openAPIV3 schema per estendere Kubernetes con nuove risorse.

Sono visualizzabili con ```kubectl get xrd```.

> Users create composite resources **(XRs)** and Claims **(XCs)** using the API schema defined by an **XRD**.

Creare una Custom Resource Definition comporta la definizione di 
- **group** 
  di solito un nome di dominio, definisce una collezione di API correlate
- **names**
  specificano come referenziare la risorsa (kind: UpperCamelCased, consigliabile iniziare con "x", plural: lowercase)
- **version**
  sistema di versioning per la risorsa, indica quanto è stabile 
- **schema**
  definisce nome, tipo e l'"opzionalità" dei parametri. Ogni versione a uno schema unico

```yaml
apiVersion: apiextensions.crossplane.io/v1
kind: CompositeResourceDefinition
metadata:
  name: xdatabases.custom-api.example.org
spec:
  group: custom-api.example.org
  names:
    kind: xDatabase
    plural: xdatabases
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              region:
                type: string 
              size:
                type: string  
              name: #opzionale
                type: string  
            required: 
              - region 
              - size
    # Removed for brevity
```
Dopo aver applicato il file, crossplane crea una nuova custom resource definition che matcha la API definita

**Abilitare i claim** ad usare l'XRD:
> Common Crossplane convention is to use claimNames that match the XRD names, but without the beginning “x.”

```yaml
apiVersion: apiextensions.crossplane.io/v1
kind: CompositeResourceDefinition
metadata:
  name: xdatabases.custom-api.example.org
spec:
  group: custom-api.example.org
  names:
    kind: xDatabase
    plural: xdatabases
  claimNames:
    kind: Database    #unici, non possono essere uguali ad altri claim o ad altri kind di XRD
    plural: databases
  versions:
  # Removed for brevity
```

**Gestione di segreti**:
Possibilità di definire un ```connectionSecretKeys```
```yaml
 connectionSecretKeys:
    - username
    - password
    - address
```

[_**Altre policy sulla documentazione**_](https://docs.crossplane.io/v1.19/concepts/composite-resource-definition/)

## 6. Composite Resource (XR): . 

Sono le risorse create dopo che un utente ha chiamato la API personalizzata definita in un XRD.
Reminder: le composite resources rappresentano un insieme di MR come un singolo oggetto Kubernetes

Ogni volta che un utente alla API custom, Crossplane crea un XR e linka tutte le relative ManagedResources.

## 7. Claims

Claims sono simili alle Composite Resources, ma esistono all'interno di un namespace in Kubernetes. 
Ogni Claim è collegato a una singola Composite Resource con scope a livello di cluster. 
Gli utenti creano dei Claims nei loro namespace specifici, garantendo così l'isolamento 
delle risorse XRD rispetto ad altri team che operano in namespace diversi.

```kubectl describe``` sul claim mostra le informazioni sulla Composite Resource a cui è collegato.

```yaml
apiVersion: apiextensions.crossplane.io/v1  #| apiVersion: example.org/v1alpha1
kind: CompositeResourceDefinition           #| kind: database
metadata:                                   #| metadata:
  name: xmydatabases.example.org            #|   name: my-claimed-database
spec:                                       #| spec: [ importante il ResourceRef, che collega il claim alla Composite Resource ]     
  group: example.org                        # Lo trovo con kubectl describe database.example.org/my-claimed-database
  names:
    kind: XMyDatabase
    plural: xmydatabases
  claimNames:
    kind: Database
    plural: databases
  # Removed for brevity
```

È possibile definire un nome del secret object dove Crossplane salverà i dettagli della connessione.

 writeConnectionSecretToRef:
    name: my-claim-secret


# Documentazione consultata
- [official Crossplane documentation](https://crossplane.io/docs/).
- https://www.youtube.com/watch?v=tbMCWp7rsk8
- https://www.youtube.com/watch?v=2l8j_yxJbow
- https://www.youtube.com/watch?v=AtbS1u2j7po&list=PLyicRj904Z9_X62k6_XM_xlJkSyoQDkS2