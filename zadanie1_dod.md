# Opis
Projekt realizuje zadania 1 - 3, punkt czwartu powoduje nieoczekiwane problemy.

## Instrukcja

### Stworzenie obrazu na platformy linux/amd64 i linux/arm64, a także zapisanie danycgh cache w repozytorium na Docker Hub.
```bash
docker buildx build --no-cache --cache-to=type=registry,ref=nazwa:tag,max-size=1GB --platform linux/amd64,linux/arm64 -t nazwa:tag --push .
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/ed685ac5-eee6-403a-909f-687ff95ebc34)

### Stworzenie obrazu z użyciem cache z Docker Hub:

```bash
docker buildx build --cache-from=type=registry,ref=nazwa:tag --platform linux/amd64,linux/arm64 -t nazwa:tag --push .
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/c861cb36-2bb3-4a49-90dc-1df3d3450700)


### Stworzenie kontenera na podstawie obrazu

```bash
docker run -d -p 8080:8080 --name nazwa_kontenera nazwa:tag
```

![image](https://github.com/N33Qu/ZAD1/assets/116431498/8b2ff7d3-7e32-4d51-9eeb-3a8fc91ecfba)

### Działanie aplikacji:

![image](https://github.com/N33Qu/ZAD1/assets/116431498/5504a02f-d4ef-40c1-a7d4-cc2bb1a76007)

![image](https://github.com/N33Qu/ZAD1/assets/116431498/ccf01f44-5485-4989-8124-e689f89f736c)

![image](https://github.com/N33Qu/ZAD1/assets/116431498/c036a737-b50a-4d97-86d5-2b5e155864ea)

### Sprawdzenie docker scout:

![image](https://github.com/N33Qu/ZAD1/assets/116431498/ed83e65a-4434-4300-90d2-9b3e8847b698)

### Linki do Docker Hub-a:

Zad1_dodL:
https://hub.docker.com/layers/neeqoo/lab/zad1_dod/images/sha256-fb815d4a700d270c8cccb4eefc8fd47cdf2e68e35c17f053a4b3cf16ad7b300f?context=repo

Zad1_cache:
https://hub.docker.com/layers/neeqoo/lab/zad1_cache/images/sha256-d4ed9230d7cf2756e03ad117668a032d017f3d2baeaa94a9c3cf23044d4475af?context=repo
