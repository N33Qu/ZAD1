# Opis

Ten projekt zawiera kod źródłowy i plik Dockerfile potrzebny do zbudowania obrazu kontenera serwera HTTP.

## Instrukcje

###Aby zbudować obraz kontenera, wykonaj następujące polecenie w katalogu głównym projektu:

```bash
docker buildx build -f Dockerfile -t neeqoo/lab:zad1_Go --sbom=true --provenance=mode=max --push .
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/0a675f1a-79f3-4a34-9e3d-060ee0776b27)

###Aby uruchomić kontener na podstawie zbudowanego obrazu, należy wykonać polecenie:

```bash
docker run -d -p 8080:8080 --name nazwa_kontenera neeqoo/lab:zad1_Go
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/d259941b-c1c8-4ce7-9e92-de4b9c042e5a)

###Aby uzyskać informacje wygenerowane przez serwer w trtakcie uruchomienia, należy użyć polecenia:

```bash
docker logs nazwa_kontenera
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/351035ac-3e1c-4b0e-b14c-809a7defa019)

###Aby sprawdzić ile warst posiada zbudowany obraz, należy użyć polecenia:

```bash
docker history neeqoo/lab:zad1_Go
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/85c3d5bd-b79a-444a-a4ad-2cdb06a2e53a)

###Sprawdzenie czy żadne składow obrazu nie uzysjkują oceny CVSS większej od 7.0:

```bash
docker scout cves --platform linux neeqoo/lab:zad1_Go
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/fd7eb7c7-9fff-47b4-b953-e5542e8d540f)

###Działanie uruchomionej aplikacji:

![image](https://github.com/N33Qu/ZAD1/assets/116431498/d8a3ce8c-7eb5-4a24-995a-1899e895ab89)


###Link do Docker Hub z obrazem:

https://hub.docker.com/layers/neeqoo/lab/zad1_Go/images/sha256-0f6d2466e878b70de3d0511ac0412c351dde1f4ac38c6fd7ffae3334bcab53a8?context=repo
