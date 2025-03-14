# TIROCINIO
**Argomento: Studio e implementazione di Provisioning e Deployment di Workload su Dispositivi IoT attraverso Kubernetes Operators**

## FASE 0: Documentazione generale e comprensione delle API di docker e di kubernetes:

> Resoconto: 
> - [Resoconto docum. generale](generale.md)
> - [Resoconto docker](docker/docker.md)

## FASE 1: Installazione e configurazione di un cluster k3s su una o più macchine, creando un ambiente Kubernetes leggero per gestire applicazioni distribuite.

### Installazione (su uno o più nodi), comprensione dell'architettura e delle API 

Ho installato un nodo k3s su un'unica macchina, comprendendo come avrei potuto estenderlo ad altri nodi.

Ho compreso il funzionamento dell'architettura e come si coordinano i vari agenti per offrire il servizio

### Creazione app nginx con Deployment, Service e IngressController
1. Deployment (nginx-deployment.yaml)

Mediante il Deployment ho definito il modo in cui il pod Nginx viene gestito in Kubernetes. Nel mio caso avvia 3 repliche (pods) dell'immagine nginx:latest. Ogni replica è una copia indipendente dell'applicazione. Anche se uno dei pod dovesse fallire, Kubernetes creerà un nuovo pod per mantenerne il numero di repliche definito.

2. Service (nginx-service.yaml)

Ho configurato il Service per fungere da bilanciatore di carico per i pod Nginx (il type è quindi "LoadBalancer"), distribuendo il traffico in ingresso verso i vari pod di Nginx. Essi saranno accessibili tramite un IP e una porta fissa.

3. Ingress (nginx-ingress.yaml)

L'ingress controller "traefik" fornisce un hostname (nel mio caso, nginx:local) per accedere al servizio.

> [Resoconto](kubernetes/kubernetes.md)

### Provisioning mediante Crossplane

Ho installato Crossplane mediante helm

> [Resoconto](crossplane/crossplane.md)

**Visualizzazione**

## FASE 2: Installare e configurare strumenti di monitoraggio come Prometheus, integrandoli con Kubernetes e configurando gli operatori per una gestione automatica.

## FASE 3: Creare e configurare un operatore Kubernetes per il deployment e la gestione automatica dei workload che eseguono su dispositivi IoT, verificando che l'operatore possa scalare e gestire i container in base alle esigenze di carico.

## FASE 4: Simulare flussi di dati provenienti da dispositivi IoT per testare il sistema, creando applicazioni containerizzate che elaborano i flussi di dati e monitorano le loro performance.

## FASE 5: Implementare meccanismi di scalabilità automatica dei container e configurare il bilanciamento del carico tra i nodi del cluster per garantire efficienza e ridondanza.

## FASE 6: Simulare guasti nel sistema per testare la resilienza del cluster Kubernetes, implementando strategie di recovery e failover per garantire la disponibilità continua dei workload sui dispositivi IoT.

