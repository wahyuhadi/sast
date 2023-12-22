## Ho to run 

```sh
    $ supply-chain -h 
```

## with config 

scan dir
```yaml
    lang: go
    rules: auto
    scan_type: dir
    projects:
    - /home/Sha256/my-space/Tools/supply-chain
    - /home/Sha256/my-space/Tools/nuc-fuzzing-template

```

scan remote repo 

```yaml
    lang: go
    rules: auto
    scan_type: repo
    projects:
    - github.com/sirupsen/logrus

```