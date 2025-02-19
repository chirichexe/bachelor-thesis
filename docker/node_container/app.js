// index.js
const express = require('express');
const app = express();
const port = 3000;

// Endpoint che restituisce "Hello, World!"
app.get('/', (req, res) => {
  res.send('Hello, World!');
});

// Avvia il server
app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
