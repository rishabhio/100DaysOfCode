const pug = require('pug');
const express = require('express');
const app = express();
const options = {};

const locals = {};

app.get('/', ( req, res ) => {
    const data = '<h1> Article Title </h1> ';
    const html = pug.renderFile('templates/article.pug', {article: data, next: 'hello world' } );
    return res.send( html );
});

app.listen( 3000 );