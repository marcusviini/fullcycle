const mysql = require('mysql');

const config = {
  config: {
    host: 'mysql',
    user: 'test',
    password: 'test',
    database: 'test'
  }
}

const connection = mysql.createConnection(config.config)

connection.connect(function(error){
  if (error) {
    console.error('❌ Query failed: ', error);
    throw error;
  }
  console.log('✅ Connection successful');
});

module.exports = connection