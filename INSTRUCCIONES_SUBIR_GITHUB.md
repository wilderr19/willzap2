# 📤 Instrucciones: Subir a GitHub

## 🎯 Objetivo

Subir este código a tu repositorio: `https://github.com/wilderr19/willzap2`

---

## 📋 PASO A PASO

### Paso 1: Inicializar Git (Si no está inicializado)

```bash
cd google-cloud-run-proxy
git init
```

### Paso 2: Agregar Archivos

```bash
git add .
```

### Paso 3: Commit Inicial

```bash
git commit -m "Initial commit: Google Cloud Run Proxy para v2ray"
```

### Paso 4: Agregar Remote

```bash
git remote add origin https://github.com/wilderr19/willzap2.git
```

### Paso 5: Verificar Remote

```bash
git remote -v
```

### Paso 6: Push al Repositorio

```bash
# Si el repositorio está vacío, usar:
git push -u origin main

# O si usa master:
git push -u origin master
```

---

## 🔄 Si el Repositorio Ya Tiene Contenido

### Opción A: Forzar Push (Sobrescribir)

```bash
git push -u origin main --force
```

**⚠️ ADVERTENCIA:** Esto eliminará todo el contenido anterior del repositorio.

### Opción B: Merge (Mantener Contenido Existente)

```bash
git pull origin main --allow-unrelated-histories
git push -u origin main
```

---

## 📝 COMANDOS COMPLETOS (Copia y Pega)

```bash
# Ir al directorio
cd google-cloud-run-proxy

# Inicializar git (si no está)
git init

# Agregar todos los archivos
git add .

# Commit
git commit -m "Initial commit: Google Cloud Run Proxy para v2ray"

# Agregar remote
git remote add origin https://github.com/wilderr19/willzap2.git

# Push
git push -u origin main
```

---

## 🔧 Si Tienes Problemas

### Error: "remote origin already exists"

```bash
# Eliminar remote existente
git remote remove origin

# Agregar de nuevo
git remote add origin https://github.com/wilderr19/willzap2.git
```

### Error: "failed to push some refs"

```bash
# Hacer pull primero
git pull origin main --allow-unrelated-histories

# O forzar push (si estás seguro)
git push -u origin main --force
```

### Error: "authentication failed"

```bash
# Usar token de GitHub o SSH
# Opción 1: Token
git remote set-url origin https://TU_TOKEN@github.com/wilderr19/willzap2.git

# Opción 2: SSH
git remote set-url origin git@github.com:wilderr19/willzap2.git
```

---

## ✅ Verificación

Después del push, verifica en GitHub:

1. Ve a: `https://github.com/wilderr19/willzap2`
2. Deberías ver los archivos:
   - `main.go`
   - `Dockerfile`
   - `go.mod`
   - `README.md`
   - `.gitignore`

---

## 🎯 Resumen

**Comandos finales:**
```bash
cd google-cloud-run-proxy
git init
git add .
git commit -m "Initial commit: Google Cloud Run Proxy"
git remote add origin https://github.com/wilderr19/willzap2.git
git push -u origin main
```

**¡Listo! Tu código estará en tu repositorio.**

