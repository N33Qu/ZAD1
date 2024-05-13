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
