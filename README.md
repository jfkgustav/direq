# DIREQ

DIREQ: direct_request! Låt publiken avgöra vilken låt ni spelar härnäst!🎶🎉✨

## Usage

Klona repot och kör sedan make:

```bash
    make
```

Som kommer att bygga tailwind, generera templ-filer och till slut sätta igång go-servern.
Besök sedan [http://localhost:7331](http://localhost:7331) där sidan visas. Om ändringar görs i någon av templ-filerna kommer det att uppdateras så länge servern övervakas.

# [Figjam med idéer om funktionalitet, flödesdiagram och user scenarios](https://www.figma.com/board/quZ5O9BUbecfrgayWHqxO9/direq?node-id=2-179&t=45BrxA5xuUB6XvIK-1)
# [Figma design-fil med mockups av användargränssnittet](https://www.figma.com/design/Yh4yYD5mQgCwycVn0cGmpz/direq-design?node-id=0-1&t=ZXFgOCBAfx9nv8Wp-1)

🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶

Inspirerat av: [https://www.demoll.se/onskalataranders](https://www.demoll.se/onskalataranders)

🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶🎶
# Funktionalitet

Applikationen består av 4 huvudsakliga vyer. Dessa kallas:
  - Publik (Mobil)
  - Musiker Live (Mobil eller helst surfplatta)
  - Musiker Redigera (Mobil)
  - Musiker Redigera (Dator)

Utöver dessa kommer även en vy för att logga in och sätta upp synk med kalkylarket också krävas.

## Publik funktioner: 📲
I vyn "Publik" kommer dessa funktioner att finnas. Vyn ska utvecklas för att kunna fungera på smarta telefoner främst och behöver inte vara anpassad för datorskärmar då den väldigt sällan kommer att visas på dessa.
  - Sök
  - Filter
  - Välj låt(ar)
  - Önska icke-existerande låt
  - Skicka
  - Mottagandebekräftelse?
  - Ange namn/bordsnummer

## Musiker Live: 📲
I vyn "Musiker Live" kommer dessa funktioner att finnas. Vyn ska utvecklas för att kunna fungera på surfplattor främst men också smarta telefoner. Behöver inte vara anpassad för datorskärmar då den väldigt sällan kommer att visas på dessa.
  - Visa önskemål
  - Sortera önskemål
  - Bocka för att ett önskemål har/ska spela(t)s
  - Godkänn önskelåt som redan spelats under kvällen
  - Skapa session
  - Avsluta session
  - Blockera användare :/

## Musiker redigera 📲 💻
I vyn "Musiker Live" kommer dessa funktioner att finnas. Denna vy ska gå att använda med såval surfplatta och smart telefon som dator för att kunna vara anpassningsbar för olika användares vanor.
  - Synka med kalkylarksapplikation
  - Få feedback när tillagda låtar inte följer korrekt format
  - Lägga till / ta bort låtar
  - Lägga till från tidigare önskemål
  - Se historik / statistik från önskemål

En låt består av fälten "Artist" och "Titel" och kan ha en eller flera taggar som kan användas för att kategorisera låtar senare, antingen för publiken när de ska söka på låt eller för musikern för att snabbt kunna bygga en spelordning utifrån kategori/er.


## Oklart ⁉⁉⁉❓❓❓
Funktioner som kanske kan finnas men som är oklar var den ska finnas:
  - Visa QR-kod för sessionen

---

Bild från Figjam med ovanstående information x)

![image](https://github.com/user-attachments/assets/86e44242-d59d-4454-b05a-ce8c659d6268)
