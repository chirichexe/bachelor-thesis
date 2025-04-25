# Immagine ufficiale di node.js come base
FROM node:lts-alpine

ENV NODE_ENV=production

# Directory di lavoro dentro il container
WORKDIR /usr/src/app

# Copia dei .json di configurazione nel container
COPY ["package.json", "package-lock.json*", "npm-shrinkwrap.json*", "./"]

# Operazioni: 
# 1. installa moduli ( Rimosso && mv node_modules ../ )
# 2. copia tutto il resto del progetto
RUN npm install --production --silent
COPY . .

# Porta esposta
EXPOSE 3000

# Imposta i permessi
#RUN chown -R node /usr/src/app
#USER node

# Applicazione avviata
CMD ["npm", "start"]
