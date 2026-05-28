# GO Haber Scraper - Yavuzlar

Golang ile geliştirilmiş, haber sitelerinden anlık olarak veri çeken, filtreleme ve dosyaya yazma yeteneklerine sahip dinamik bir Komut Satırı Arayüzü (CLI) web scraping aracıdır.

## 🛠️ Desteklenen Siteler

1. **The Hacker News** (`-1`) - Siber Güvenlik Haberleri
2. **Technopat** (`-2`) - Güncel Teknoloji Akışı (RSS Feed kanalı üzerinden)
3. **DonanımHaber** (`-3`) - Bilim ve Teknoloji Gelişmeleri

## 💻 Kullanım Rehberi

Projeyi yerelde çalıştırmak veya test etmek için aşağıdaki komut kombinasyonlarını kullanabilirsiniz.

### Temel Çalıştırma
```bash
# Birinci siteyi (The Hacker News) tetikler ve verileri kaydeder
go run main.go -1

# İkinci siteyi (Technopat) tetikler ve verileri kaydeder
go run main.go -2

# Üçüncü siteyi (DonanımHaber) tetikler ve verileri kaydeder
go run main.go -3

```

### Filtreleme Kombinasyonları
```
# Sadece içerik başlıklarını ve açıklamalarını getirir, tarihleri gizler
go run main.go -1 -date

# Sadece içerik başlıklarını ve tarihlerini getirir, açıklamaları gizler
go run main.go -2 -description

# Aynı anda her iki filtreyi de uygulayarak sadece başlıkları listeler
go run main.go -3 -date -description
```
### Uygulamadan Çıkış
```
# Güvenli çıkış modunu tetikler
go run main.go -4
```

### Kurulum
```
git add README.md
git commit -m "Docs: Add comprehensive README.md"
git push origin main```
