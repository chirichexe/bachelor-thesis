# DOCKER

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
docker run --publish LISTENPORT:CONTAINERPORT --name NOMECONTAINER IMMAGINE
```
Esempio:
```sh
sudo docker run -it nginx /bin/bash  # Esegue un container interattivo con Bash
```

## Avvio di un server Nginx
### Operazioni sulle immagini
```sh
sudo docker run -d -p 8080:80 --name MIO_NGINX nginx  # Avvia un container in background sulla porta 8080
sudo docker images                                    # Elenca tutte le immagini Docker disponibili
sudo docker image rm [ID]                             # Rimuove un'immagine specificata
```

### Operazioni sui container
```sh
sudo docker ps                                       # Elenca i container in esecuzione
sudo docker container exec -it MIO_NGINX bash        # Apre una shell all'interno del container
sudo docker stop NOMECONTAINER                      # Ferma un container
sudo docker rm NOMECONTAINER                        # Rimuove un container
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
   sudo docker build -t MIA_IMMAGINE .
   ```
   > L'immagine viene creata con successo e può essere verificata con `sudo docker images`

4. **Eseguire il container:**
   ```sh
   sudo docker run --name MIO_CONTAINER -p 3000:3000 MIA_IMMAGINE
   ```
   Per avviare un container in modalità interattiva:
   ```sh
   sudo docker run -it --rm MIA_IMMAGINE sh
   ```

5. **Ricostruire il container con un'installazione pulita:**
   ```sh
   sudo docker stop MIO_CONTAINER
   sudo docker rm MIO_CONTAINER
   sudo docker rmi MIA_IMMAGINE  # Rimuove l'immagine precedente
   ```

## Concetti chiave di Docker
### Immagine
Un'immagine Docker è un file immutabile che include tutto il necessario per eseguire un'applicazione: codice, librerie, dipendenze, file di configurazione e variabili d'ambiente. Le immagini sono la base per la creazione dei container.

### Container
Un container è un'istanza in esecuzione di un'immagine. È un ambiente isolato che esegue il software definito nell'immagine, condividendo il kernel del sistema operativo host. I container sono leggeri e isolati tra loro e dal sistema host.

### Differenza tra immagine e container
- **Immagine** → È un template immutabile che contiene tutto il necessario per eseguire un'applicazione.
- **Container** → È un'istanza eseguibile di un'immagine, un ambiente isolato e indipendente.
