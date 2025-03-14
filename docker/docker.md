# DOCKER - GENERALE

## Installazione
Per installare Docker, eseguire lo script `docker_install.sh` oppure seguire la guida ufficiale disponibile su [Docker Docs](https://docs.docker.com/get-docker/).

## Comandi base
```sh
docker info      # Mostra informazioni dettagliate sulla configurazione di Docker
docker version   # Mostra la versione di Docker installata
docker login     # Effettua il login su Docker Hub o altro registry
```

## Eseguire un container
Per avviare un container Docker, usare il comando:
```sh
docker run --publish [LISTENPORT]:[CONTAINERPORT] --name [NOME_CONTAINER] [NOME_IMMAGINE]
```
Esempio:
```sh
docker run -it nginx /bin/bash  # Esegue un container interattivo con Bash
```

## Avvio di un server Nginx
### Operazioni sulle immagini
```sh
docker run -d -p 8080:80 --name MIO_NGINX nginx  # Avvia un container in background sulla porta 8080
docker images                                    # Elenca tutte le immagini Docker disponibili
docker image rm [ID]                             # Rimuove un'immagine specificata
```

### Operazioni sui container
```sh
docker ps                                         # Elenca i container in esecuzione
docker container exec -it [MIO_NGINX] bash        # Apre una shell all'interno del container
docker stop [NOME_CONTAINER]                      # Ferma un container
docker rm [NOME_CONTAINER]                        # Rimuove un container
```

Per uscire da un terminale attachato al container senza fermarlo:
```
CTRL+P CTRL+Q
```

## Esempio: Avvio di un server Node.js con Docker
1. **Inizializzare il progetto Node.js:**
    ```sh
    npm init -y   # Crea il file package.json
    npm install   # Installa le dipendenze
    ```

2. **Configurare Docker in VS Code:**
    - Aprire la Command Palette (`CTRL+SHIFT+P`)
    - Cercare `Docker: Add Docker Files to Workspace`
    - Selezionare `Node.js`

3. **Creare l'immagine del container:**
    ```sh
    docker build -t [NOME_IMMAGINE] .
    ```
    > L'immagine viene creata con successo e può essere verificata con ` docker images`

4. **Eseguire il container:**
    ```sh
    #detached -------|
    #		         V
    docker run -d --name [NOME_CONTAINER] -p 3000:3000 [NOME_IMMAGINE]

    #log accessibili mediante
    docker logs -f [NOME_CONTAINER]

    #per accedere all'interno del container
    docker exec -it [NOME_CONTAINER] sh

    ```
        
    Per avviare un container in modalità interattiva:
    ```sh
    docker run -it --rm [NOME_IMMAGINE] sh
    ```

5. **Ricostruire il container con un'installazione pulita:**
    ```sh
    docker stop [NOME_CONTAINER]
    docker rm [NOME_CONTAINER]
    docker rmi [NOME_IMMAGINE]  # Rimuove l'immagine precedente
    ```

## Concetti chiave di Docker
### Immagine
Un'immagine Docker è un file immutabile che include tutto il necessario per eseguire un'applicazione: codice, librerie, dipendenze, file di configurazione e variabili d'ambiente. Le immagini sono la base per la creazione dei container.

### Container
Un container è un'istanza in esecuzione di un'immagine. È un ambiente isolato che esegue il software definito nell'immagine, condividendo il kernel del sistema operativo host. I container sono leggeri e isolati tra loro e dal sistema host.

### Differenza tra immagine e container
- **Immagine** → È un template immutabile che contiene tutto il necessario per eseguire un'applicazione.
- **Container** → È un'istanza eseguibile di un'immagine, un ambiente isolato e indipendente.

# DOCKER - PERSISTENZA

### Volumi
```sh
docker create volume [NOME_VOLUME]
docker volume ls
docker create inspect [NOME_VOLUME]
docker create rm [NOME_VOLUME] # prune per eliminare tutti
```
È conveniente avviare il test con

```docker run -d --name [NOME_CONTAINER] -v [NOME_VOLUME]:[DIRECTORY_DEL_CONTAINER] [...] ```
Così abbiamo direttamente montato il volume in una directory

Vantaggi:
- Persistenza: Anche se elimini il container, i dati nel volume rimangono.
- Isolamento: Il volume non dipende dalla struttura delle directory locali.
- Portabilità: Puoi montare lo stesso volume in più container senza problemi.


# DOCKER - COMPOSE
Creare applicazioni multi-container (magari uno per ogni "ruolo" specifico es. backend, frontend ...) 
definendo dei file YAML.
Utile per piccoli progetti e test, meno completo di kubernetes.

## Linguaggio YAML
Utilizzato come file di configurazione ```docker-compose.yml``` 

```yml
key: value
map:
    map-key: value
sequence:
    - el1
    - el2
json-map: {"key" : "val"}
```

Esempio file di configurazione docker:

```yml
services:
  node-app:
    build:
      context: ./node_container
      dockerfile: Dockerfile
    container_name: node-container
    ports:
      - "3000:3000"
    volumes:
      - node-volume:/usr/src/app/logs
    restart: unless-stopped

volumes:
  node-volume:
    driver: local

```
## Compose
Elenco la serie di passaggi per avviare il docker compose
```sh
docker compose build # Builda l'immagine
docker compose up -d # avvia in background
# docker compose down  - per stoppare
docker compose exec [NOME_APP] sh #per accedere al container
```

### Avvio docker compose

```sh
docker compose up -d # partito correttamente
```

Provando ora ad eseguire

```sh
docker compose -p [NOME_PROGETTO] up -d
```
Esso fallirà perchè la porta 3000 risulta già in uso.

Soluzione: sul file di configurazione:

```yml
services:
  node-app:
    build:
      context: ./node_container
      dockerfile: Dockerfile
    container_name: node-container
    ports:
      - "3001:3000" # Incremento la porta esterna
    volumes:
      - node-volume:/usr/src/app/logs
    restart: unless-stopped
```

Output:
```console
❯ sudo docker compose ls
NAME                STATUS              CONFIG FILES
docker              running(1)          /home/davide/Desktop/bachelor-thesis/docker/docker-compose.yml
test                running(1)          /home/davide/Desktop/bachelor-thesis/docker/docker-compose.yml
```

# DOCKER - CONTAINER REGISTRIES
è una repository centrale per le immagini Docker. Posso fare push e store del binario dell'immagine su un container registry.
Docker Hub (ufficiale) o altri repository privati

```sh
docker images # per vedere le immagini
docker login    
docker tag node-image-test:latest chirichexe/node-image-test:latest # associo all'immagine un tag
sudo docker push chirichexe/node-image-test:latest # pusho l'immagine sul container registry
```

# Documentazione consultata
- https://docs.docker.com/
- https://youtu.be/kTp5xUtcalw?si=vxgi2Le5tnYmP8HX