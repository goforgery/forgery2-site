/*
    To run this you'll need to have nodejs installed.

    npm i express
    node init.js
    curl -i http://127.0.0.1:3000/
*/

var express = require('express');
var app = express();

app.get('/', function(req, res) {
    res.send('Hello World');
});

app.listen(3000);