# Chmura Obliczeniowa Lab 6



## Table of Contents
- [Build](#build)
- [Screens](#screens)


## Build

Polecenie budowania obrazu:

```bash
docker build  --secret id=ghlab_access,src=<ścieżka do klucza publicznego> --build-arg VERSION='1.0' --platform linux/amd64,linux/arm64 -t ghcr.io/n33qu/pawcho6:lab6 .  
```

Wynik polecenia:
```bash
[+] Building 778.8s (34/34) FINISHED                                                                                                docker:default
 => [internal] load build definition from Dockerfile                                                                                          0.0s
 => => transferring dockerfile: 1.65kB                                                                                                        0.0s 
 => [linux/arm64 internal] load metadata for docker.io/library/nginx:1.22                                                                     2.3s 
 => [linux/amd64 internal] load metadata for docker.io/library/nginx:1.22                                                                     2.3s 
 => [auth] library/nginx:pull token for registry-1.docker.io                                                                                  0.0s
 => [internal] load .dockerignore                                                                                                             0.0s
 => => transferring context: 2B                                                                                                               0.0s 
 => [internal] load build context                                                                                                             0.2s 
 => => transferring context: 7.75MB                                                                                                           0.2s 
 => [linux/amd64 nginx 1/4] FROM docker.io/library/nginx:1.22@sha256:fc5f5fb7574755c306aaf88456ebfbe0b006420a184d52b923d2f0197108f6b7         5.6s 
 => => resolve docker.io/library/nginx:1.22@sha256:fc5f5fb7574755c306aaf88456ebfbe0b006420a184d52b923d2f0197108f6b7                           0.0s 
 => => sha256:2a9f38700bb5a0462e326fe3541b45f24a677ac3cd386c4922d48da5fbb6f0a8 1.41kB / 1.41kB                                                0.2s 
 => => sha256:fd071922d543e072b21cb41a513634657049d632fe48cfed240be2369f998403 773B / 773B                                                    0.2s 
 => => sha256:ac713a9ef2cca7a82e27f0277e4e3d25c64d1cf31e4acd798562d5532742f5ef 958B / 958B                                                    0.3s
 => => sha256:ef2fc869b944b87eaf25f4c92953dc69736d5d05aa09f66f54b0eea598e13c9c 626B / 626B                                                    0.3s 
 => => sha256:fd03b214f77493ccb73705ac5417f16c7625a7ea7ea997e939c9241a3296763b 25.57MB / 25.57MB                                              2.0s 
 => => sha256:f1f26f5702560b7e591bef5c4d840f76a232bf13fd5aefc4e22077a1ae4440c7 31.41MB / 31.41MB                                              2.2s
 => => extracting sha256:f1f26f5702560b7e591bef5c4d840f76a232bf13fd5aefc4e22077a1ae4440c7                                                     1.7s
 => => extracting sha256:fd03b214f77493ccb73705ac5417f16c7625a7ea7ea997e939c9241a3296763b                                                     1.0s 
 => => extracting sha256:ef2fc869b944b87eaf25f4c92953dc69736d5d05aa09f66f54b0eea598e13c9c                                                     0.0s 
 => => extracting sha256:ac713a9ef2cca7a82e27f0277e4e3d25c64d1cf31e4acd798562d5532742f5ef                                                     0.0s 
 => => extracting sha256:fd071922d543e072b21cb41a513634657049d632fe48cfed240be2369f998403                                                     0.0s 
 => => extracting sha256:2a9f38700bb5a0462e326fe3541b45f24a677ac3cd386c4922d48da5fbb6f0a8                                                     0.0s 
 => [linux/arm64 nginx 1/4] FROM docker.io/library/nginx:1.22@sha256:fc5f5fb7574755c306aaf88456ebfbe0b006420a184d52b923d2f0197108f6b7         7.4s 
 => => resolve docker.io/library/nginx:1.22@sha256:fc5f5fb7574755c306aaf88456ebfbe0b006420a184d52b923d2f0197108f6b7                           0.0s 
 => => sha256:8d9209af03bce9f88f0d4efca543b7e12f811376fe936e71556043585b5b67ca 1.40kB / 1.40kB                                                0.5s 
 => => sha256:9427c5e3c7e65d31829e6abe5a56aed8e39acf52329ec29b42e6b93463105324 772B / 772B                                                    0.4s 
 => => sha256:129a8129e575ba4a03c0fb2d50154236a4b4cef49e30b00ffedcc2cc7f6d637f 958B / 958B                                                    0.2s 
 => => sha256:0137ac0f244275c4d8d8292fcfd0213ede1475a252ebf3d84f6688e8ea6d57b7 627B / 627B                                                    0.2s 
 => => sha256:3dc85e50d06487a804f6b2e69eba13a96eebc5957d415da155b1a9e7f72179be 25.49MB / 25.49MB                                              2.6s 
 => => sha256:fcdb9667c46b09d1c1d058681ea4a1db41e66bbc1a71d873a0c9da4f7a92947d 30.06MB / 30.06MB                                              3.9s 
 => => extracting sha256:fcdb9667c46b09d1c1d058681ea4a1db41e66bbc1a71d873a0c9da4f7a92947d                                                     1.3s 
 => => extracting sha256:3dc85e50d06487a804f6b2e69eba13a96eebc5957d415da155b1a9e7f72179be                                                     0.6s 
 => => extracting sha256:0137ac0f244275c4d8d8292fcfd0213ede1475a252ebf3d84f6688e8ea6d57b7                                                     0.0s 
 => => extracting sha256:129a8129e575ba4a03c0fb2d50154236a4b4cef49e30b00ffedcc2cc7f6d637f                                                     0.0s 
 => => extracting sha256:9427c5e3c7e65d31829e6abe5a56aed8e39acf52329ec29b42e6b93463105324                                                     0.0s 
 => => extracting sha256:8d9209af03bce9f88f0d4efca543b7e12f811376fe936e71556043585b5b67ca                                                     0.0s 
 => [linux/amd64 build  1/10] ADD ./alpine-minirootfs-3.17.3-aarch64.tar /                                                                    0.1s 
 => [build  2/10] RUN apk update && apk upgrade                                                                                              11.5s 
 => [build  2/10] RUN apk update && apk upgrade                                                                                              10.9s 
 => [build  3/10] RUN apk add --no-cache npm nodejs openssh-client git                                                                        7.2s 
 => [build  3/10] RUN apk add --no-cache npm nodejs openssh-client git                                                                        7.0s 
 => [build  4/10] RUN mkdir -p /clonedRepo                                                                                                    1.0s 
 => [build  4/10] RUN mkdir -p /clonedRepo                                                                                                    1.0s 
 => [build  5/10] RUN --mount=type=secret,id=ghlab_access     mkdir -p ~/.ssh &&     cp /run/secrets/ghlab_access ~/.ssh/id_rsa &&     chmod  5.0s 
 => [build  5/10] RUN --mount=type=secret,id=ghlab_access     mkdir -p ~/.ssh &&     cp /run/secrets/ghlab_access ~/.ssh/id_rsa &&     chmod  5.1s 
 => [build  6/10] RUN npx create-react-app react_app_lab6                                                                                   538.6s 
 => [build  6/10] RUN npx create-react-app react_app_lab6                                                                                   540.4s 
 => [build  7/10] RUN mv /clonedRepo/src/App.js /react_app_lab6/src/App.js                                                                    0.8s 
 => [linux/amd64 build  8/10] WORKDIR react_app_lab6                                                                                          0.0s 
 => [build  9/10] RUN npm install                                                                                                            57.7s 
 => [build  7/10] RUN mv /clonedRepo/src/App.js /react_app_lab6/src/App.js                                                                    1.1s 
 => [linux/arm64 build  8/10] WORKDIR react_app_lab6                                                                                          0.0s 
 => [build  9/10] RUN npm install                                                                                                            57.1s 
 => [build 10/10] RUN npm run build                                                                                                         153.4s 
 => [build 10/10] RUN npm run build                                                                                                         150.3s 
 => [linux/arm64 nginx 2/4] COPY --from=build /react_app_lab6/build/. /var/www/html                                                           0.0s 
 => [linux/arm64 nginx 3/4] COPY --from=build clonedRepo/nginx.conf /etc/nginx/conf.d/default.conf                                            0.0s 
 => [linux/arm64 nginx 4/4] COPY --from=build clonedRepo/nginx.conf /etc/nginx/nginx.conf                                                     0.0s 
 => [linux/amd64 nginx 2/4] COPY --from=build /react_app_lab6/build/. /var/www/html                                                           0.1s 
 => [linux/amd64 nginx 3/4] COPY --from=build clonedRepo/nginx.conf /etc/nginx/conf.d/default.conf                                            0.0s 
 => [linux/amd64 nginx 4/4] COPY --from=build clonedRepo/nginx.conf /etc/nginx/nginx.conf                                                     0.0s 
 => exporting to image                                                                                                                        0.4s 
 => => exporting manifest sha256:cb34c16cb0e095d8b0d4d7ee4447ea406601735b94884b9f8b724ab7312114f8                                             0.0s 
 => => exporting config sha256:c52832d29b856f07c9d6ad2aefaa0539800fbb64d3c7409cf09a1f79371b7082                                               0.0s 
 => => exporting attestation manifest sha256:8dd8476d0c3f661b4e571d123450bdd6af31abf0755d61a55bf622c91edb8012                                 0.0s 
 => => exporting manifest sha256:2e72ef269195b1c8c5e44300095a492f47bc36d8bce447c285bc1f297a8c4697                                             0.0s 
 => => exporting config sha256:027b9a90f9ff9c8cba68227c8b9398282f30638510a94610fca2fcbe02f3a25f                                               0.0s 
 => => exporting attestation manifest sha256:e8f52ae12f48adf84b8b08fa3eab3653c9fb7078bbf1743b5f116195b945ad47                                 0.0s 
 => => exporting manifest list sha256:e523d4854f7f3c64aa62d13b2da4e933da10b312ebf5792e3ee127e213a43439                                        0.0s 
 => => naming to ghcr.io/n33qu/pawcho6:lab6                                                                                                   0.0s 
 => => unpacking to ghcr.io/n33qu/pawcho6:lab6                                                                                                0.1s 

View build details: docker-desktop://dashboard/build/default/default/uii837wc4en2pgphppmxm69ys```
```

Przesłanie obrazu na githuba (po zalogowaniu z tokenem):
```bash
docker push ghcr.io/n33qu/pawcho6:lab6
```

Wynik polecenia:
```bash
The push refers to repository [ghcr.io/n33qu/pawcho6]
096cb04b1015: Waiting                                                                                                                              
0137ac0f2442: Waiting                                                                                                                              
fcdb9667c46b: Pushed
0ef66d146302: Pushed
3dc85e50d064: Pushed
fd03b214f774: Pushed
2a9f38700bb5: Pushed
ef2fc869b944: Pushed
129a8129e575: Pushed
fd071922d543: Pushed
f1f26f570256: Pushed
7a7ecd9d31fd: Pushed
8d9209af03bc: Pushed
236248f32d7c: Pushed
5b911a51b168: Pushed
ceceece9eff7: Pushed
bc8e23dfe884: Pushed
15e254f26d56: Pushed
lab6: digest: sha256:e523d4854f7f3c64aa62d13b2da4e933da10b312ebf5792e3ee127e213a43439 size: 1609
```
Link do package:
https://github.com/users/N33Qu/packages/container/package/pawcho6

Uruchamianie kontenera:
```bash
docker run -p 80:80 --name=<name> <image_name>
```

Sprawdzenie kontenera:
```bash
docker ps -a  
```


## Screens

Efekt działania aplikacji na nginx

![image](https://github.com/N33Qu/pawcho6/assets/116431498/3ffed679-c1af-478e-a076-c89514afa96f)

