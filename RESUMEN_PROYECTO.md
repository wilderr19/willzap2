# 📋 Resumen del Proyecto: Google Cloud Run Proxy

## ✅ Archivos Incluidos

### Archivos del Repositorio Original:

1. **`main.go`** - Código fuente del proxy (Go)
   - Proxy TCP simple
   - Redirige tráfico a `V2RAY_SERVER_IP:80`
   - Maneja conexiones concurrentes

2. **`Dockerfile`** - Imagen Docker
   - Build multi-stage
   - Imagen final minimalista (distroless)
   - Optimizado para Cloud Run

3. **`go.mod`** - Dependencias Go
   - Módulo: `tcp-proxy`
   - Go 1.21

4. **`README.md`** - Documentación mejorada
   - Instrucciones de despliegue
   - Ejemplos de uso
   - Información de costos

### Archivos Adicionales Creados:

5. **`.gitignore`** - Archivos a ignorar en Git
6. **`INSTRUCCIONES_SUBIR_GITHUB.md`** - Guía para subir a GitHub
7. **`subir-a-github.sh`** - Script automatizado
8. **`RESUMEN_PROYECTO.md`** - Este archivo

---

## 🎯 Funcionalidad

### ¿Qué Hace?

El proxy recibe conexiones TCP en Google Cloud Run y las redirige a tu servidor v2ray.

### Flujo:

```
Cliente → Cloud Run (puerto automático) → VPS (V2RAY_SERVER_IP:80)
```

### Variables de Entorno:

- **`V2RAY_SERVER_IP`**: IP de tu VPS con v2ray (requerida)
- **`PORT`**: Puerto asignado por Cloud Run (automática)

---

## 🚀 Cómo Usar

### 1. Desplegar en Cloud Run:

```bash
gcloud run deploy cloud-run-proxy \
  --source . \
  --set-env-vars V2RAY_SERVER_IP=TU_IP_VPS \
  --allow-unauthenticated \
  --region us-central1
```

### 2. O Build y Deploy:

```bash
# Build
gcloud builds submit --tag gcr.io/TU_PROYECTO/cloud-run-proxy

# Deploy
gcloud run deploy cloud-run-proxy \
  --image gcr.io/TU_PROYECTO/cloud-run-proxy \
  --set-env-vars V2RAY_SERVER_IP=TU_IP_VPS \
  --allow-unauthenticated
```

---

## 📤 Subir a GitHub

### Opción 1: Script Automatizado

```bash
cd google-cloud-run-proxy
chmod +x subir-a-github.sh
./subir-a-github.sh
```

### Opción 2: Manual

```bash
cd google-cloud-run-proxy
git init
git add .
git commit -m "Initial commit: Google Cloud Run Proxy"
git remote add origin https://github.com/wilderr19/willzap2.git
git push -u origin main
```

---

## 📊 Comparación con Original

| Aspecto | Original | Tu Versión |
|---------|----------|------------|
| **Repositorio** | RicardoB3/Google_Cloud_Run_Proxy | wilderr19/willzap2 |
| **Dependencia** | Fork del original | Independiente |
| **README** | Básico | Mejorado con ejemplos |
| **Documentación** | Mínima | Completa |
| **Scripts** | No | Sí (subir-a-github.sh) |

---

## ✅ Ventajas de Tener Tu Propio Repositorio

1. ✅ **Independencia**: No dependes del repositorio original
2. ✅ **Control total**: Puedes modificar sin restricciones
3. ✅ **Personalización**: Agrega tus propias mejoras
4. ✅ **Documentación**: README mejorado con ejemplos
5. ✅ **Scripts**: Herramientas para facilitar el uso

---

## 🔧 Próximos Pasos

1. **Subir a GitHub**: Usa el script o comandos manuales
2. **Desplegar en Cloud Run**: Prueba el despliegue
3. **Personalizar**: Agrega mejoras según necesites
4. **Documentar**: Actualiza README con tus cambios

---

## 📝 Notas

- El código es simple pero efectivo
- Funciona como proxy TCP directo
- Compatible con v2ray y otros servicios TCP
- Optimizado para Cloud Run

---

**¡Todo listo para subir a tu repositorio!**

