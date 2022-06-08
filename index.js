var express = require('express')
var fs = require('fs')
var axios = require('axios')
var moment = require('moment-timezone')
var child_process = require('child_process')

var app = express()
app.use(express.urlencoded({extended:false}))
app.use(express.json())

moment.tz.setDefault('Asia/Shanghai');

var json_server = 'http://localhost:5555/words/';

var port = process.env.PORT || 3000;

var token = process.env.TOKEN;
var secret = process.env.SECRET;

var pull = async () => {
  var p = new Promise((t,c) => {
    child_process.exec('cd ./words-db && git pull origin master && git config --global user.email "sunwu51@126.com" && git config --global user.name "frank"', (e, stdout, stderr) =>{
      if(!e){
        t(stdout)
      }else {
        c(e)
      }
    })
  })
  await p;
}

var push = async (words) => {
  var p = new Promise((t,c) => {
    child_process.exec(`cd ./words-db && git add words.json && git commit -m "add word ${words}" && git push https://${token}@github.com/sunwu51/words-db.git`, (e, stdout, stderr) =>{
      if(!e){
        t(stdout)
      }else {
        c(e)
      }
    })
  })
  await p;
}

var getCurMonday = () => {
  return moment().weekday(1).format('YYYY-MM-DD');
}


pull();

var wordsQueue = [];


setInterval(async ()=>{
  if (wordsQueue.length > 0){
    var str = wordsQueue.toString()
    wordsQueue = [];
    console.log( str + ' uploading')
    try {
      await push(str);
    } catch(e) {
      console.error(e)
    }
  }
}, 1000 * 60);


app.get('/', async (req,res)=>{
  try{
    var r = await axios({url: json_server})
    res.json(r.data)
  }catch(e){ res.sendStatus(404)}
});

app.get('/:id', async (req,res)=>{
  var id = req.params.id;
  try{
    var r = await axios({url: json_server + id})
    res.json(r.data)
  }catch(e){ res.sendStatus(404)}
});

var mutex = 0;
app.post('/add', async (req, res)=>{
  var word = req.body.word;
  var s = req.body.secret;

  if(s != secret){
    res.send(403);
    return;
  }
  if(!word){
    res.sendStatus(403);
    return;
  }
  if(mutex == 1){
    res.json({code: -1, msg: 'another process running, plz wait for a minute'});
    return;
  }
  try{
    mutex = 1;
    var id = getCurMonday();
    var item = await new Promise( (t,c) => {
      axios({url: json_server + id}).then(xx=>t(xx.data))
      .catch(e=>t(null))
    });
    console.log(`add word ${word}, the item in db is ${JSON.stringify(item)}`)
    var r;
    if(!item) {
      item = {id, list: [word]}
      r = await axios({url: json_server,  method:'POST', data: item})
    }else {
      item = {id, list: [word, ...item.list]}
      r = await axios({url: json_server + id,  method:'PUT', data: item})
    }
    wordsQueue.push(word)
    res.json(r.data)
  }catch(e){
    console.error(e)
    res.sendStatus(500)
  }finally{
    mutex = 0;
  }

})

app.listen(port, ()=>{console.log("app start")})