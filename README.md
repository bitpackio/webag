# webag

GIT BFEFEHLE

-> Lokales Git Repository personalisieren:
	git config git config --global user.name <"Dein github Name">

-> Status des lokalen git Repository anzeigen lassen
	git status

-> Anlagen eines neuen (lokalen) git repositorys aus einem bestehenden Repository von github:
	git clone <URL des Git-Repositorys auf github>

-> Wenn neue Datei angelegt oder eine bestehende Datei im lokal geändert wurde: 	- hinzufügen zum lokalen Repository:
	git add <Bezeichnung der Datei/ des Ordners die/der hinzugefügt werden soll> (added die datei den Ordner, diese müssen aber sichtbar sein)
	git add * (added einfach alle Änderungen)
	git add . (added alle änderungen des aktuellen ordners)

	- dann (zur sicherheit) das lokale Repository auf den aktuellen stand bringen:
	git pull (Saugt alle inzwischen auf github aktualisierten Dateien - du hast nun die aktuelle Version)

	- dann committen damit deine Mitprogrammierer wissen was genau du geändert hast
	git commit -m <"änderungs Nachricht">

	- dann in das gemeinsame Repository auf github pushen
	git push -u origin main

-> Bevor du lokal loslegst und an einer Datei arbeitest: Immer erst die aktuelle Verion mit "git pull" pullen!

CONSOLEN BEFEHLE ZUR DATEIVERWALTUNG

-> Einen neuen Ordner anlegen: mkdir <Ordnername>
-> Eine neue Datei erstellen: touch <dateiname.txt>
-> Eine Datei in der Console bearbeiten oder erstellen: vim <dateiname.text>
-> Eine Datei löschen: rm <Dateiname>
-> Einen Ordner löschen: rm -r <Ordnername>



CONSOLEN EDITOR VIM BEDIENEN

Zwei Modi (werden unten links angezeigt): -- EINFÜGEN -- und Speichern/Schließen/Kopieren ...
-- EINFÜGEN -- ist der Eingabemodus: Man kann nur im Eingabemodus schreiben
Wenn man nicht im Eingabemosud ist, kann man Speichern, Schließen, kopieren usw.

-> In den Eingabemodus: Taste i drücken (jetzt steht unten links -- EINFÜGEN --)
-> Den Eingabemodus verlassen: Taste Ecs drücken (-- EINFÜGEN -- ist unten links verschwunden)
-> SPEICHERN: 1) Eingabemodus verlassen 2) :w eingeben, mit Enter bestätigen
-> Schließen und VIM verlassen: 1) Eingabemodus verlassen 2) :q, mit Enter bestätigen
-> Speichern und Schließen: 1) Eingabemodus verlassen 2) :wq, mit Enter bestätigen


