const express = require('express');
const connection = require('./connection');
const app = express();
const port = 3000;

app.get('/', async (req, res) => {
 const title = '<h1>Full Cycle Rocks!</h1>'
  const people = await query()
  if(!people) {
    res.send(`${title}<h2>No people found</h2>`)
    return
  }
  const list = people.map(person => `<li>${person.name}</li>`).join('')
  const ul = `<ul>${list}</ul>`
  res.send(`${title}${ul}`)
})

app.listen(port, async () => {
  console.log(`App listening at http://localhost:${port}`)
  createTable()
  insert()
})

// private

const query = async () => {
  const select = 'SELECT * FROM people'
  const promise = new Promise((resolve, reject) => {
    connection.query(select, (error, results, fields) => {
      if (error) {
        reject(error)
      }
      resolve(results)
    })
  })
  return promise
}

const insert = () => {
  const insert = 'INSERT INTO people (name) values ("Marcus"), ("Aline"), ("Maria"), ("Hitallo"), ("Larissa")'
  return connection.query(insert)
}

const createTable = () => {
  const createTable = 'CREATE TABLE IF NOT EXISTS people(id int not null auto_increment, name varchar(255), primary key(id))'
  return connection.query(createTable)
}