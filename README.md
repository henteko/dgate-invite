# dgate-invite

dgate-invite is [DeployGate Invite API](https://deploygate.com/docs/api) CLI tool

## How to use
### login
```
$ dgate-invite -login
# Owner Name:[input your owner name]
# token:[input your deploygate api key]
```

### Get App Member List
```
$ dgate-invite -g -p [input target app]
```

#### ex
```
$ dgate-invite -g -p com.henteko07.assisthack
```

### Invite User for App
```
$ dgate-invite -i -p [input target app] [input user list]
```

#### ex
```
$ dgate-invite -i -p com.henteko07.assisthack henteko2 henteko3
```

### Delete User for App
```
$ dgate-invite -d -p [input target app] [input user list]
```

#### ex
```
$ dgate-invite -d -p com.henteko07.assisthack henteko2 henteko3
```

### Logout
```
$ dgate-invite -logout
```


## License
MIT  
Copyright 2014, henteko
