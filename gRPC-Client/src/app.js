
require('dotenv').config();
const express = require("express");
const morgan = require("morgan");
const cors = require("cors");
const client = require('./gRPC_client')

let random = require('random-name')


const app = express();

//Settings
const port = process.env.PORT || 3000;
app.set('PORT', port);

//Middls
app.use(express.json());
app.use(morgan("dev"));
app.use(cors());


app.post('/sendGame', (req, res)=>{
    const {game_id, players} = req.body;
    
    client.SendResultGame({game_id, players}, function(err, response) {
        if(err) {
            res.status(400).json('Error on sending result');
        }
        res.status(200).json(response.response_Game)
    });
});

app.get('/', (req, res) => {
    res.status(200).json({name: random(), msg: 'Hello World'});
})

app.listen(app.get('PORT'), ()=>{
    console.log("Server on port: ", app.get('PORT'))
})


