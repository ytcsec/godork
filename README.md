# GoDork

Hedefinizdeki domaine ait oluşturulan dorklarla hedef üzerinde istediğiniz kategoriye ait sonuçları getiren tool

## Kullanım

1. Uygulamayı başlatın: `./godork.exe`
2. Search yerine hedef domaini girin (örnek: `parslar.org`)
3. **Go** butonuna basın
4. Dork sonuçlarını alın

### Örnek site

```
parslar.org
```

### Örnek sonuç

```
site:parslar.org filetype:pdf
site:parslar.org intitle:"index of"
site:parslar.org inurl:admin/login
```

## Kategoriler

| Kod | Arayüz Adı | Amaç |
|-----|------------|------|
| HassasDosyalar | Gizli Dosyalar | env ve pdf vb. dışa açılmayan dosyalar |
| DizinListeleme | Dizinler | Açık bırakılmış dizin listelemeleri |
| LoginSayfalari | Loginler | Admin ve login panelleri |
| HataMesajlari | Errorlar | SQL gibi hatalar |
| Paneller | Paneller | phpMyAdmin, cPanel, wp-admin vb. yönetim panelleri |
| ConfigDosyalari | Configler | `.env`, `wp-config`, `.git` gibi config dosyalar |
| YedekDosyalar | Yedekler | `.bak`, `.old`, `backup` içeren yedek dosyalar |
| KullaniciBilgileri | Kullanıcılar | Kullanıcı adı, parola, API anahtarı içerebilecek dosyalar |
| KameraIoT | IoT | Public kamera ve IoT cihazlar |

## Dork Listesini Genişletmek için

Yeni bir dork eklemek için `main.go` içindeki `dorkListesi` dizisine `"kategori|sorgu"` formatında bir satır ekleyin:

```go
var dorkListesi = []string{
    // Mevcut dorklar...
    "HassasDosyalar|filetype:kdbx",
    "LFI|inurl:?file=",
}
```

Yeni bir kategori eklediyseniz, kullanıcıya gösterilecek adı da `kategoriIsimleri` mapine ekleyin:

```go
var kategoriIsimleri = map[string]string{
    "LFI": "LFI Tespit Dorku",
}
```
