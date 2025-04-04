# etap budowania aplikacji
FROM scratch AS build

# dodanie systemu bazowego (Alpine)
ADD alpine-minirootfs-3.21.3-x86_64.tar /

# zmienna wersji aplikacji
ARG VERSION="bombardilo_crocodilo_1.2"

# ustawienie powłoki
SHELL ["/bin/sh", "-c"]

# instalacja zależności do budowy aplikacji
RUN apk update && apk add --no-cache go git build-base

# ustawienie katalogu roboczego
WORKDIR /app

# skopiowanie pliku źródłowego
COPY main.go .

# ustawienie zmiennej środowiskowej
ENV VERSION=${VERSION}

# kompilacja i generowanie pliku HTML
RUN go build -o builder main.go && \
    ./builder && \
    cp index.html /tmp/index.html

# etap końcowy z serwerem nginx
FROM nginx:alpine

# skopiowanie statycznego pliku HTML do katalogu serwowanego przez nginx
COPY --from=build /tmp/index.html /usr/share/nginx/html/index.html

# dodanie healthchecka
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget -q -O- http://localhost:80 || exit 1

# uruchomienie nginx w trybie pierwszoplanowym
CMD ["nginx", "-g", "daemon off;"]
