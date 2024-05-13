
```bash
docker buildx build --no-cache --cache-to=type=registry,ref=nazwa:tag,max-size=1GB --platform linux/amd64,linux/arm64 -t nazwa:tag --push .
```
