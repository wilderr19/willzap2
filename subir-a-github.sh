#!/bin/bash

# Script para subir el código a GitHub
# Uso: ./subir-a-github.sh

set -e

REPO_URL="https://github.com/wilderr19/willzap2.git"
BRANCH="main"

echo "=========================================="
echo "Subiendo a GitHub: willzap2"
echo "=========================================="
echo ""

# Verificar que estamos en el directorio correcto
if [ ! -f "main.go" ]; then
    echo "❌ Error: No se encuentra main.go"
    echo "   Ejecuta este script desde el directorio google-cloud-run-proxy"
    exit 1
fi

# Inicializar git si no está inicializado
if [ ! -d ".git" ]; then
    echo "[1/5] Inicializando repositorio Git..."
    git init
    echo "✅ Git inicializado"
else
    echo "[1/5] Git ya está inicializado"
fi

# Agregar archivos
echo ""
echo "[2/5] Agregando archivos..."
git add .
echo "✅ Archivos agregados"

# Commit
echo ""
echo "[3/5] Creando commit..."
git commit -m "Initial commit: Google Cloud Run Proxy para v2ray" || echo "⚠️  No hay cambios para commitear"
echo "✅ Commit creado"

# Configurar remote
echo ""
echo "[4/5] Configurando remote..."
if git remote | grep -q "origin"; then
    echo "⚠️  Remote 'origin' ya existe, actualizando..."
    git remote set-url origin "$REPO_URL"
else
    git remote add origin "$REPO_URL"
fi
echo "✅ Remote configurado: $REPO_URL"

# Push
echo ""
echo "[5/5] Subiendo a GitHub..."
echo ""

# Intentar push
if git push -u origin "$BRANCH" 2>&1 | grep -q "error\|fatal"; then
    echo "⚠️  Error al hacer push. Intentando con --force..."
    read -p "¿Sobrescribir contenido existente? (s/N): " confirm
    if [[ $confirm =~ ^[Ss]$ ]]; then
        git push -u origin "$BRANCH" --force
        echo "✅ Push completado (forzado)"
    else
        echo "❌ Push cancelado"
        exit 1
    fi
else
    echo "✅ Push completado"
fi

echo ""
echo "=========================================="
echo "✅ Código subido exitosamente"
echo "=========================================="
echo ""
echo "Repositorio: $REPO_URL"
echo "Branch: $BRANCH"
echo ""
echo "Verifica en: https://github.com/wilderr19/willzap2"
echo ""

