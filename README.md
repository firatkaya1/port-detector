# port-detector
A simple CLI for detect the open ports and reporting

# Nedir ? 
port-detector terminal üzerinde çalışan,işletim sistemindeki açık olan portların listelenip, rapor oluşturulduğu bir terminal uygulamasıdır.

# Nasıl Kurulur ?

### Adım 1
Github repositoriden proje linkini kopyalayın ve kendi bilgisayarınıza indirin.  

```
git clone https://github.com/firatkaya1/port-detector.git
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
```
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
### Adım 5 (Opsiyonel)
Her defasında komuta ulaşabilmek için build ettiğimiz dosyayı *./find-duplicates* diyerek çalıştırmak zahmetli gelebilir. Linux kullanıcıları için elde ettiğiniz build dosyasını /bin dosyasının içine kopyalamanız gerekmektedir.
```
sudo cp find-duplicates /bin  
```
### Bitti. 
Artık terminal'den dosyanın kendisi olmadan çalıştırabilir ve kullanabilirsiniz. 

# Nasıl Kullanılır ? 

Bir kullanıcı senaryosuna göre işletim sisteminizde açık olan portlar bulunmakta ve bu portların nasıl kapatılacağınızı bilmediğini farz edelim. Bu portlar üzerinden, yada bu portlar dışındaki portları kapatmak istiyorsunuz ve bunu nasıl yapacağınızı bilmiyorsunuz. Sizler için işletim sisteminizi tarayıp, ufak bir rapor çıkartan bu terminal tabanlı uygulama mükemmel bir çözüm olacaktır.

### Adım 1 
Aşağıda bulunan komut bir'den fazla flag almaktadır. 
**--language**; oluşturmak istediğiniz raporun ana dilini temsil eder. İki farklı değer alabilir. **tr** ve **en**.  Kısa kullanımı -l şeklinde mümkündür.
**--name**; çıkarmak istediğiniz rapor dosyasının ismidir. Opsiyoneldir. Kısa kullanımı **-n** şeklinde mümkündür.
**--path**; çıkarmak istediğiniz raporun işletim sisteminde nereye çıkartılacağını belirtir. Kısa kullanımı **-p** şeklindedir. 

```
./port-detector detect -l tr --name port-detector -p /home/kaya/Desktop
```
Yukarıdaki komutun çalıştırıldığı takdirde aşağıdaki şekilde örnek bir çıktı oluşturucaktır:
```
Start to detecting...
Port Detect : 0.0.0.0:631
Port Detect : 0.0.0.0:5939
Port Detect : 0.0.0.0:36477
Port Detect : 0.0.0.0:57621
Dil : Turkish
Dosya'nın çıkarıldığı yer : /home/kaya/Desktop
Dosya Adı : port-detector
Tam adresi : /home/kaya/Desktop/port-detector.pdf
Rapor oluşturma süresi :  160.964941ms
```
---
#### Author

Mevcut author bilgisini görebilmek istiyorsanız, aşağıdaki komutu çalıştırın. 
```
port-detector author
```

#### Version

Mevcut version bilgisini görebilmek istiyorsanız, aşağıdaki komutu çalıştırın. 
```
port-detector version
```

#### license

Mevcut version bilgisini görebilmek istiyorsanız, aşağıdaki komutu çalıştırın. 
```
port-detector license
```

#### Help

Mevcut tüm komutları görebilmek istiyorsanız help komutunu kullanabilirsiniz.
```
port-detector help
```

# Komutu bilgisayarımdan nasıl kaldırabilirim ? 
Build ettiğimiz dosyayı /bin klasörüne kopyalamıştık. Yapmanız gereken tek şey /bin klasörü altındaki find-duplicates dosyasını bulup silmektir. Bunu terminal ekranından aşağıdaki komut ile yapabilirsiniz. Aynı zaman da klonladığımız dosyayı silmeniz yeterlidir. 

```
sudo rm /bin/port-detector
```
---
Bir problemle karşılaşırsanız bana yazabilirsiniz. 

[me@kayafirat.com](mailto:me@kayafirat.com?subject=[GitHub]%20port-detector)

