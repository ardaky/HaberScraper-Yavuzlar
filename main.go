package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ilkSite(aciklamaGizlenecekMi bool, tarihGizlenecekMi bool) {
	istek, hata := http.Get("https://www.thehackernews.com/")
	if hata != nil {
		fmt.Println("Hata:", hata)
		return
	}
	defer istek.Body.Close()

	if istek.StatusCode != 200 {
		fmt.Println("Hata: İstek başarısız oldu. Durum Kodu:", istek.StatusCode)
		return
	}

	icerik, hata := goquery.NewDocumentFromReader(istek.Body)
	if hata != nil {
		fmt.Println("Hata: Belge oluşturulamadı.", hata)
		return
	}

	dosya, hata := os.OpenFile("sonuc.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if hata != nil {
		fmt.Println("Hata: Dosya açılamadı.", hata)
		return
	}

	defer dosya.Close()

	icerik.Find(".body-post").Each(func(i int, s *goquery.Selection) {
		baslik := s.Find(".home-title").Text()
		aciklama := s.Find(".home-desc").Text()

		tarihRaw := s.Find(".item-label").Text()
		tarihTemiz := strings.ReplaceAll(tarihRaw, "", "")
		tarih := strings.TrimSpace(tarihTemiz)

		fmt.Printf("%d. Haber:\n%s\n", i+1, baslik)
		fmt.Fprintf(dosya, "%d. Haber:\n%s\n", i+1, baslik)

		if aciklamaGizlenecekMi == false {
			fmt.Printf("Aciklama: %s\n\n", aciklama)
			fmt.Fprintf(dosya, "Aciklama: %s\n\n", aciklama)
		}

		if tarihGizlenecekMi == false {
			fmt.Printf("Tarih: %s\n\n", tarih)
			fmt.Fprintf(dosya, "Tarih: %s\n\n", tarih)
		}
	})
}

func ikinciSite(aciklamaGizlenecekMi bool, tarihGizlenecekMi bool) {

	istek, hata := http.Get("https://feedpress.me/technopat")
	if hata != nil {
		fmt.Println("Hata:", hata)
		return
	}
	defer istek.Body.Close()
	if istek.StatusCode != 200 {
		fmt.Println("Hata: İstek başarısız oldu. Durum Kodu:", istek.StatusCode)
		return
	}

	if istek.StatusCode != 200 {
		fmt.Println("Hata: İstek başarısız oldu. Durum Kodu:", istek.StatusCode)
		return
	}

	icerik, hata := goquery.NewDocumentFromReader(istek.Body)
	if hata != nil {
		fmt.Println("Hata: Belge oluşturulamadı.", hata)
		return
	}

	dosya, hata := os.OpenFile("sonuc.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if hata != nil {
		fmt.Println("Hata: Dosya açılamadı.", hata)
		return
	}
	defer dosya.Close()

	icerik.Find("item").Each(func(i int, s *goquery.Selection) {
		baslik := strings.TrimSpace(s.Find("title").Text())
		aciklama := strings.TrimSpace(s.Find("description").Find("p").First().Text())
		tarih := strings.TrimSpace(s.Find("pubDate").Text())

		if baslik == "" {
			return
		}

		fmt.Printf("%d. Haber:\n%s\n", i+1, baslik)
		fmt.Fprintf(dosya, "%d. Haber:\n%s\n", i+1, baslik)

		if aciklamaGizlenecekMi == false {
			fmt.Printf("Aciklama: %s\n", aciklama)
			fmt.Fprintf(dosya, "Aciklama: %s\n", aciklama)
		}

		if tarihGizlenecekMi == false {
			fmt.Printf("Tarih: %s\n", tarih)
			fmt.Fprintf(dosya, "Tarih: %s\n", tarih)
		}
		fmt.Println()
		fmt.Fprintln(dosya)
	})
	fmt.Println("Veriler 'sonuc.txt' dosyasina basariyla kaydedildi!")
}

func ucuncuSite(aciklamaGizlenecekMi bool, tarihGizlenecekMi bool) {

	istemci := &http.Client{}
	istek, hata := http.NewRequest("GET", "https://www.donanimhaber.com/teknoloji-haberleri", nil)

	istek.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	istek.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	yanit, hata := istemci.Do(istek)

	if hata != nil {
		fmt.Println("Hata:", hata)
		return
	}
	defer yanit.Body.Close()

	if yanit.StatusCode != 200 {
		fmt.Println("Hata: İstek başarısız oldu. Durum Kodu:", yanit.StatusCode)
		return
	}

	icerik, hata := goquery.NewDocumentFromReader(yanit.Body)
	if hata != nil {
		fmt.Println("Hata: Belge oluşturulamadı.", hata)
		return
	}

	dosya, hata := os.OpenFile("sonuc.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if hata != nil {
		fmt.Println("Hata: Dosya açılamadı.", hata)
		return
	}
	defer dosya.Close()

	icerik.Find("article.blogItem").Each(func(i int, s *goquery.Selection) {
		baslik := strings.TrimSpace(s.Find("a.baslik").Text())
		aciklama := strings.TrimSpace(s.Find("div.aciklama").Text())
		tarih := strings.TrimSpace(s.Find("bilgi. span").Text())

		if baslik == "" {
			return
		}

		fmt.Printf("%d. Haber:\n%s\n", i+1, baslik)
		fmt.Fprintf(dosya, "%d. Haber:\n%s\n", i+1, baslik)

		if aciklamaGizlenecekMi == false {
			fmt.Printf("Aciklama: %s\n", aciklama)
			fmt.Fprintf(dosya, "Aciklama: %s\n", aciklama)
		}

		if tarihGizlenecekMi == false {
			fmt.Printf("Tarih: %s\n", tarih)
			fmt.Fprintf(dosya, "Tarih: %s\n", tarih)
		}
		fmt.Println()
		fmt.Fprintln(dosya)
	})
	fmt.Println("Veriler 'sonuc.txt' dosyasina basariyla kaydedildi!")
}
func main() {

	site1 := flag.Bool("1", false, "İlk sitenin baslik bilgisini getirir. thehackernews.com")
	site2 := flag.Bool("2", false, "İkinci sitenin baslik bilgisini getirir. technopat.net")
	site3 := flag.Bool("3", false, "Üçüncü sitenin baslik bilgisini getirir. webtekno.com")

	cikis := flag.Bool("4", false, "Programdan çıkış yapar.")

	tarihGizlemece := flag.Bool("date", false, "Tarih bilgilerini gizler.")
	aciklamaGizlemece := flag.Bool("description", false, "Aciklama bilgilerini gizler.")

	flag.Parse()

	if *site1 {
		ilkSite(*aciklamaGizlemece, *tarihGizlemece)
	} else if *site2 {
		ikinciSite(*aciklamaGizlemece, *tarihGizlemece)
	} else if *site3 {
		ucuncuSite(*aciklamaGizlemece, *tarihGizlemece)
	} else if *cikis {
		fmt.Println("Programdan çıkış yapılıyor... Allaha emanetsiniz")
		return
	}

}
