const express = require('express')
const app = express()
const port = 3000

app.get('/login', (req, res) => {
  console.log('Login called with get');
  console.table(req.headers);
  res.send('Hello World fron auth-service!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})