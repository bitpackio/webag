Web-AG README
=============

Git Befehle
-----------

- Lokales Git Repository personalisieren:  
	`git config --global user.name "DEIN GITHUB-NAME"`  
	`git config --global user.email EMAIL-ADRESSE@DOMAIN.COM`

- Status des lokalen git Repository anzeigen lassen  
	`git status`

- Für eine Stunde (= 3600 Sekunden) keine Zugangsdaten bei pull oder push eingeben müssen:  
	`git config --global credential.helper 'cache --timeout=3600'`

- Anlegen eines neuen (lokalen) Git-Repositorys aus einem bestehenden Repository von github:  
	`git clone https://github.com/URL/ZUM/GIT-REPOSITORY.git`
	- Link zu diesem Git-Repository: https://github.com/bitpackio/webag.git

- Dateien oder Ordner entfernen oder umbenennen: siehe Befehle zur Dateiverwaltung, Bewehlen jeweils `git` vorstellen

- Wenn neue Datei angelegt oder eine bestehende Datei lokal geändert wurde:
	- hinzufügen zum lokalen Repository:  
	`git add ORDNER oder PFAD/ZUR/DATEI` (added den Ordner oder die Datei im Ordner, diese müssen aber sichtbar sein)  
	`git add *` (added alle Ordner und Dateien im aktuellen Ordner)  
	`git add .` (added den aktuellen Ordner, quasi gleichbedeutend mit `git add *`)

	- dann (zur Sicherheit) das lokale Repository nochmal auf den aktuellen Stand bringen:  
	`git pull` ("saugt" alle inzwischen auf github aktualisierten Dateien – du hast nun die aktuelle Version)

	- dann committen, damit deine Mitprogrammierer wissen was genau du geändert hast:  
	`git commit -m "NACHRICHT ZUR ÄNDERUNG"`

	- dann in das gemeinsame Repository auf github pushen:  
	`git push -u origin main`

Hinweise:
- Bevor du lokal an einer Datei arbeitest oder Änderungen committest: Immer erst die aktuelle Version mit `git pull` pullen!
- Bei mehreren, voneinander unabhängigen Änderungen, z.B. in verschiedenen Dateien:  
	Mehrere Commits machen, damit man die Änderungen einzeln sieht und ggf. eine rückgängig machen kann, ohne die anderen zu beeinflussen


Konsolen-Befehle zur Dateiverwaltung
------------------------------------

- Einen neuen Ordner anlegen: `mkdir ORDNER`
- Eine neue Datei erstellen: `touch DATEI.txt`
- Eine Datei in der Console bearbeiten oder erstellen: `vim DATEI.txt`

Außerhalb eines Git-Repositorys: (ansonsten vor den Befehl `git` schreiben, z.B. `git rm DATEI.txt`)
- Eine Datei löschen: `rm DATEI.txt`
- Einen Ordner und seinen Inhalt löschen: `rm -r ORDNER`
- Eine Datei umbenennen und/oder verschieben: `mv ALTER/PFAD/ALTERDATEINAME.txt NEUER/PFAD/NEUERDATEINAME.txt`
- Einen Ordner umbenennen und/oder verschieben: `mv ALTER/PFAD/ALTERORDNERNAME NEUER/PFAD/NEUERORDNERNAME`

Übersicht wichtiger Konsolen-Befehle: https://wiki.ubuntuusers.de/Shell/Befehls%C3%BCbersicht/#Grundkommandos


Konsolen-Editor VIM bedienen
----------------------------

- Datei mit VIM öffnen: `vim DATEI.txt`

Zwei Modi (wird unten links angezeigt): `-- EINFÜGEN --` und Speichern/Schließen/Kopieren.  
`-- EINFÜGEN --` ist der Eingabemodus: Man kann nur im Eingabemodus schreiben  
Wenn man nicht im Eingabemodus ist, kann man speichern, schließen, kopieren, löschen usw.

- In den Eingabemodus gelangen: Taste `i` drücken (jetzt steht unten links `-- EINFÜGEN --`)
- Den Eingabemodus verlassen: Taste `Esc` drücken (`-- EINFÜGEN --` ist unten links verschwunden)

Außerhalb des Eingabemodus:
- Speichern: `:w` eingeben, mit `Enter` bestätigen
- Schließen und VIM verlassen: `:q`, mit `Enter` bestätigen
- Speichern und Schließen: `:wq`, mit `Enter` bestätigen
- Rückgängig machen: `dd`
- Eine Zeile kopieren: `yy`
- Eine Zeile einfügen: `p`

Übersicht zu VIM-Befehlen: https://wiki.ubuntuusers.de/VIM/#VIM-Modi


Website vom Handy aus aufrufen
------------------------------

- in `website/main.go` die Zeile `serverPort string = ":8080"` in `serverPort string = "0.0.0.0:8080"` ändern

- mit root-Rechten mit `service firewalld stop` die Firewall des Fedora-Sticks deaktivieren

- über `ifconfig` die lokale IP-Adresse (beginnend mit 192.168...) herausfinden

- auf dem Handy die IP-Adresse im Browser aufrufen (im lokalen WLAN): `http://192.168.178.XX:8080`


HTML-Klassen
------------

- einzelne Bilder:
```html
<figure>
	<img src="..." />
	[<figcaption>...</figcaption>]
</figure>
```
- mehrere Bilder: mehrere `html <figure>` tags in einem `html <div class="gallery">`
- Slider: mehrere `html <figure>` tags in einem `html <div class="slider">` 
