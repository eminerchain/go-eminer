# go-eminer


### what's the eminer

As an innovative global hash rate service platform, Eminer hash rate service platform aims to reintegrate and redefine the mining industrial ecology based on hash rate with thoughts and finally build into a hash rate financial service platform integrating hash rate service,financial services and information service.

### Building the source

For prerequisites and detailed build instructions please read the [Installation Instructions](https://github.com/ethereum/go-ethereum/wiki/Building-Ethereum) on the wiki.

Building `em` requires both a Go (version 1.13 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make em
```

or, to build the full suite of utilities:

```shell
make all
```

### Deploy

Under Linux or Mac, get the compressed file from release and extract it to get the executable file em.Create the storage directory /data/em, and copy the executable file to the directory. Then execute the startup command, in which '--port' is the customized chain synchronization port, '--rpc' would open RPC service, '--rpcaddr' is the self-defined RPC listening IP which should set to 127.0.0.1 if you do not want the remote RPC connection, --rpcport is the customized RCP listening port.
for example
```
tar zxvf em-linux-amd64-1.1.16-unstable.tar.gz
mkdir –p /data/em
cp em-linux-amd64-1.1.16-unstable/em /data/em/
nohup /data/em/em --datadir /data/em/em-data --port 30303 --rpc --rpcaddr 0.0.0.0 --rpcport 8545 2>> /data/em/em.log &
```
### Attach the console
```
/data/em/em attach /data/em/em-data/em.ipc
```
### Common console commond
```
#get block height
em.blockNumber
#get block info
em.getBlock(blockHashOrBlockNumber)
#get accounts in wallet
em.accounts
#get transaction info
em.getTransaction(transactionHash)
#generate new accounts and store them in the keystore directory, encrypted with passphrase
personal.newAccount(passphrase)
#sent transaction
personal.sendTransaction({from:'affress',to:'address',value:web3.toWei(100,'em'),action:0}, "password")
#start rpc by console
admin.startRPC("0.0.0.0", 8545)
#stop rpc by console
admin.stopRPC()
```
