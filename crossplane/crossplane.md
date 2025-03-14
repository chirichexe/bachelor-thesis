# CROSSPLANE
Crossplane è un'estensione open-source per Kubernetes che permette al tuo cluster di supportare il provisioning e la gestione di infrastrutture cloud, servizi e applicazioni. Consente ai team di piattaforma di definire e offrire API per l'infrastruttura in modalità self-service ai team di sviluppo, sfruttando strumenti e pratiche nativi di Kubernetes.

## Funzionalità chiave

- **Infrastructure as Code (IaC)**: Definisci e gestisci l'infrastruttura utilizzando manifest Kubernetes.
- **Modularità**: Definisci e gestisci l'infrastruttura utilizzando manifest Kubernetes.
- **Estendibilità**: Estendi Crossplane con Custom Resource Definitions (CRD) e controller personalizzati.
- **Multi-Cloud**:  Gestisci risorse su più provider cloud da un unico control plane..

Tutto ciò che devo fare è applicare il manifest, dimenticarmi delle API

## 

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