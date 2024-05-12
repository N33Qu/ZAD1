#Wybrany obraz
FROM python:3.5-alpine3.12

#Dodanie etykiet
LABEL author="Mikolaj Stasz"


# Kopiowanie plików aplikacji do kontenera
COPY server.py /app/

# Ustawienie katalogu roboczego
WORKDIR /app

# Uruchomienie serwera po uruchomieniu kontenera
CMD ["python", "server.py"]
