# Google Cloud Run Proxy

Proxy TCP simple para Google Cloud Run que redirige tráfico a un servidor v2ray externo.

## 🎯 Descripción

Este proxy se ejecuta en Google Cloud Run y hace proxy del tráfico TCP a un servidor v2ray configurado. Es útil para:

- ✅ Usar Google Cloud Run como proxy gratuito
- ✅ Obtener HTTPS automático sin configurar certificados
- ✅ Escalar automáticamente según el tráfico
- ✅ No depender de un servidor propio para el proxy

## 🏗️ Arquitectura

```
Cliente → Google Cloud Run (Proxy) → VPS con v2ray (V2RAY_SERVER_IP:80)
```

## 📋 Requisitos

- Cuenta de Google Cloud Platform
- Proyecto de Google Cloud configurado
- Google Cloud SDK (`gcloud`) instalado
- Servidor v2ray funcionando (IP conocida)

## 🚀 Despliegue Rápido

### Opción 1: Usar Imagen Docker Pre-construida

```bash
# Desplegar en Cloud Run
gcloud run deploy cloud-run-proxy \
  --image docker.io/praveenkarunarathne/google-cloud-run-proxy \
  --set-env-vars V2RAY_SERVER_IP=TU_IP_VPS \
  --allow-unauthenticated \
  --region us-central1 \
  --platform managed
```

### Opción 2: Build y Deploy Local

```bash
# 1. Build la imagen
gcloud builds submit --tag gcr.io/TU_PROYECTO_ID/cloud-run-proxy

# 2. Desplegar
gcloud run deploy cloud-run-proxy \
  --image gcr.io/TU_PROYECTO_ID/cloud-run-proxy \
  --set-env-vars V2RAY_SERVER_IP=TU_IP_VPS \
  --allow-unauthenticated \
  --region us-central1 \
  --platform managed
```

## ⚙️ Variables de Entorno

### Requeridas:

- **`V2RAY_SERVER_IP`**: IP del servidor v2ray destino (ej: `102.129.137.225`)

### Automáticas (Cloud Run):

- **`PORT`**: Puerto asignado por Cloud Run (automático)

## 📝 Ejemplo de Uso

```bash
# Desplegar con tu IP de VPS
gcloud run deploy cloud-run-proxy \
  --image gcr.io/TU_PROYECTO_ID/cloud-run-proxy \
  --set-env-vars V2RAY_SERVER_IP=102.129.137.225 \
  --allow-unauthenticated \
  --region us-central1

# Obtener URL
gcloud run services describe cloud-run-proxy --region us-central1 --format 'value(status.url)'
```

## 🔧 Desarrollo Local

### Build Local:

```bash
# Build
go build -o proxy

# Ejecutar (requiere variables de entorno)
export PORT=8080
export V2RAY_SERVER_IP=TU_IP_VPS
./proxy
```

### Test Local:

```bash
# Con Docker
docker build -t cloud-run-proxy .
docker run -p 8080:8080 -e PORT=8080 -e V2RAY_SERVER_IP=TU_IP_VPS cloud-run-proxy
```

## 📊 Características

- ✅ **Proxy TCP simple**: Redirige tráfico TCP directamente
- ✅ **Escalable**: Cloud Run escala automáticamente
- ✅ **HTTPS automático**: Cloud Run maneja SSL/TLS
- ✅ **Tier gratuito**: 2 millones de requests/mes gratis
- ✅ **Sin gestión**: No necesitas mantener servidor

## ⚠️ Limitaciones

- ⚠️ **Cold starts**: Puede haber latencia en el primer request
- ⚠️ **Timeout**: Cloud Run tiene límites de timeout (60 min máximo)
- ⚠️ **Costo**: Después del tier gratuito puede ser costoso
- ⚠️ **Protocolo**: Solo TCP, no HTTP/HTTPS específico

## 💰 Costos

### Tier Gratuito:
- ✅ 2 millones de requests/mes
- ✅ 360,000 GB-segundos
- ✅ 180,000 vCPU-segundos

### Después del Tier Gratuito:
- 💰 $0.40 por millón de requests
- 💰 $0.0000025 por GB-segundo
- 💰 $0.0000100 por vCPU-segundo

## 🔍 Verificación

```bash
# Ver logs
gcloud run services logs read cloud-run-proxy --region us-central1

# Ver detalles
gcloud run services describe cloud-run-proxy --region us-central1

# Ver métricas
gcloud run services describe cloud-run-proxy --region us-central1 --format 'value(status)'
```

## 📚 Recursos

- [Google Cloud Run Documentation](https://cloud.google.com/run/docs)
- [Cloud Run Pricing](https://cloud.google.com/run/pricing)
- [v2ray Documentation](https://www.v2fly.org/)

## 📄 Licencia

GPL-3.0 License

## 🤝 Contribuciones

Este repositorio es un fork independiente. Siéntete libre de hacer mejoras y contribuciones.

---

**Desarrollado para uso independiente - No depende del repositorio original**

