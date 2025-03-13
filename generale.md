# 1. Architettura a microservizi

## Generalità e differenza con architettura monlitica
Variante della SOA (service oriented) -> accoppiamento debole tra elementi
Ognuno ha la propria responsabilità

Esempio architetture monolitiche:
Ogni unità è scritta, costruita e compilata come un elemento a sè stante
Se vogliamo "scalare" il sistema dobbiamo clonare l'entità
*Esempio: Modello a 3 tier*
WEB SERVER - BUSINESS - DATA 
*Problemi*
Aggiunta di elementi complicata, i processi sono collegati , se un solo elemento ha un "picco" tutta l'architettura va ridimensionata


Nei microservizi ogni parte si occupa del proprio ruolo e ognuno si occupa di una parte indipendentemente
Essi comunicano tra di loro tramite interfaccia definita da API

 - Link: https://aws.amazon.com/it/microservices/ 

## Trasformazione 
Creare una "facciata" che permetta di aggiornare il codice legacy (?)

## Benefici e rischi
Benefits:
Fault isolation
Open source :)
Deploy veloce and easy scaling

Rischi: 
importante le !! retry strategies!!

# 2. Cloud Native services
Mentalità: pets vs. cattle 
Infrastruttura immutabile ma elementi distrutti su richiesya, non AGGIORNATI e SOSTITUITI, non aggiornati
Il container è un 

## 2.1 Containerizzare applicazione
1. deploy e monitorare container
2. Continuous integration
3. Orchestratore 
4. Osservabilità (cosa avviene nel cluster? ) Prometheus
5. Proxy e sicurezza

- Link: https://www.cncf.io/projects/
- https://github.com/kubernetes/kubernetes

# 3. Container
## Generalità
Cos'è un container? È un'unità di deployment e ha tutto quello che serve al codice per essere eseguito (codice, librerie, tools ...)
Usa meno risorse e rende frammentabile, portabile isolato etc...

Cosa si intende per *Virtualizzazione* ?
VM: Ideali per "long running tasks", lente da bootare e "pesanti" al livello di memoria
CONTAINER: Leggerissime, veloci da avviare (non serve boot)

VM -> ha un'app
CONTAINER -> hanno tante VM

**Continuazione appunti, vai sulla cartella -> DOCKER**
**Continuazione appunti, vai sulla cartella -> KUBERNETES**
