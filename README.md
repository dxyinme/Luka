![luka](https://i.loli.net/2020/06/08/Sng2LXTsPUD6aod.jpg)

<p></p>

[![Build Status](https://travis-ci.com/dxyinme/Luka.svg?branch=dxyinme)](https://travis-ci.com/dxyinme/Luka)

<h3>a golang IM service</h3>

#### compile

```
<in linux>
export GO111MODULE=on
export GOPROXY=https://goproxy.io
make keeperD
make assigneerD
```

#### use 
```
<in linux, start simple cluster>
[compile]
cd bin
cd AssigneerDeployment
[change conf/assign.conf]
bash assigneer/start.sh
cd ..
cd KeeperDeployment
bash keeper/new_keeper [keeper_name] [machineIP:ListeningPort] [keeperID]
cd [keeper_name]
bash start_ICC.sh
```

#### about config

<h4>cluster.conf</h4>
the host of this service serve for.
```batch
[host] host keeperID
```
<h4>assign.conf</h4>
the config for machines we prepare to use.
```batch
IP Password
...
```



#### Luka Wiki
[LukaWiki](https://github.com/dxyinme/Luka/wiki)

#### about LukaMsg
[LukaComm](https://github.com/dxyinme/LukaComm)

#### about client
[LukaClient](https://github.com/dxyinme/LukaClient)

#### communication
QQ: 252896124 </p>
mail: ProjectLuka@yandex.com 