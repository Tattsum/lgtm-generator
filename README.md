# lgtm-generator

## deploy
### 1. Dockerイメージをビルドします。

```bash
docker build -t gcr.io/your-project-id/lgtm-generator:latest .
```

### DockerイメージをGoogle Container Registryにプッシュします。

```bash
docker push gcr.io/your-project-id/lgtm-generator:latest
```

### Cloud Runサービスをデプロイします。環境変数も設定します。

```bash
gcloud run deploy lgtm-generator \
  --image gcr.io/your-project-id/lgtm-generator:latest \
  --platform managed \
  --region your-region \
  --allow-unauthenticated \
  --set-env-vars AZURE_OPENAI_ENDPOINT=your_azure_openai_endpoint,AZURE_OPENAI_API_KEY=your_azure_openai_api_key
```
