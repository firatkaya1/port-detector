# port-detector
A simple CLI for detect the open ports and reporting

# Nedir ? 
port-detector terminal üzerinde çalışan,işletim sistemindeki açık olan portların listelenip, rapor oluşturulduğu bir terminal uygulamasıdır.

# Nasıl Kurulur ?

### Adım 1
Github repositoriden proje linkini kopyalayın ve kendi bilgisayarınıza indirin.  

```
https://github.com/firatkaya1/port-detector.git
```

### Adım 2

Klonlanan dosyayı açın. 
```
cd port-detector
```
### Adım 3 
Projeyi build edin.
```
go build -o port-detector
```
#### Dipnot
Eğer makinanız  **go commond not found** gibi bir hata alırsanız aşağıdaki komutu çalıştırın. Ardından projeyi tekrar build edin. 
```
sudo apt install golang
```
### Adım 4 
Projeyi çalıştırın. 
```
./port-detector
```
Örnek bir çıktısı : 
port-detector an command based interface which is help you to find opened ports and report them easily.
Start to using, just enter the detect command and watch the magic. Enjoy...

Usage:
  port-detector [command]

Available Commands:
  author      This command return author name.
  detect      Best port detector CLI
  help        Help about any command
  license     This command return license type.
  version     This command return version number.

Flags:
  -h, --help      help for port-detector
  -v, --version   version for port-detector

Use "port-detector [command] --help" for more information about a command.



```
