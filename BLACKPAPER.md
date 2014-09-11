# BLACKPAPER
### Here will be all that stuff that I could forget
====
## 1. Filesystem

All dns records are stored in usual fs in the backward direction with .endp in the end of name. 

For example: test.example.hype would turn in hype/example/test.endp. All records are stored in format 
recordtype	record 

The headers like dnsclass and name of record are extracted from filename.

### 1.1 Name types

All names are either zone either final node. 

Zone representation is a directory with endpoint and service files

Endpoit representation is a single file with some records in it

### 1.2 Service files

Service files are begin with ';' and written with CAPS. Currently there 3 service files:

- ;SIG -- Here is detached OpenPGP signature of serialized zone name

- ;DGATE -- Here is a table with format of <zonename>:<keyname>

- ;MIRROR -- List of nodes that are holding that zone

## 2. Signing and authority

By the design full zone name is spelled like: test.example.hype~keyname. Where keyname is a unique in network (rather by convention) name of maintainer of root directory or a key signed by getable key.

> getable key -- is a key that could be verified by signature checks

Desired format for GLOBAL keynames is root.nickname.top_level_domain. For example: root.valdek.hype

### 2.1 Delegating

The owner of upperlaying zone is writing to DGATE information: dzone:nickname~parent.zone(~...)

The ~ symbol means that the owner of parent.zone signed the 'nickname' key. When someone is mirroring a zone he asks other nodes for copy of a key and a parent signature and checks it.

## 3 Networking

Every node has an opened websocket connections to another nodes and either listens for update messages, either sending them.

### 3.1 Message strucure

Message consists from Header and body. Header districted from body with ';' symbol

#### 3.1.1 Header

Header consists of: 

* emmitter, the hostname that changed a zone

* revision, a GMT+0 time and date to minute

* type, a symbol that signals about contents message

Name of the zone is derived from URL, but functions that casting the package have need of name to hash

All fields are distrected by ':'

#### 3.1.2 Message

Can be either an update, either data message.

Update message can be either a request (type U), either answer (type u). 

In first case there is '?' and hash of current zone. 

Data message can be either a signature (type S), either serialised zone (type s)