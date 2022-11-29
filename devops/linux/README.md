## Linux commands every developer should know


#### Find the name(-name) from current path(.) with matching name as .png
```bash
find . -name '*.jpg'
```

#### Find the name(-name) from current path(.) with matching name as .png with case insensitivity
```bash
find . -iname '*.jpg'
```

### Get the CPU core

```bash
nproc
```

### Replace some text in file using sed from input.txt i.e replace $version to 1.0.1 in file input.txt
```bash
sed -i -e 's/$version/1.0.1/g' input.txt
```

### Find the os-release
cat /etc/os-relese

### Find Local time
```bash
cat -l /etc/localtime
```

### Find the os-release
cat /etc/os-release

### Find Local time
```bash
cat -l /etc/localtime
```