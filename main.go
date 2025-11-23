package main
import (
        "io"
        "log"
        "net"
        "os"
        "time"
)
func main() {
        listenAddr := ":" + os.Getenv("PORT")
        
        // Obtener host destino (DHOST o V2RAY_SERVER_IP como fallback)
        targetHost := os.Getenv("DHOST")
        if targetHost == "" {
                targetHost = os.Getenv("V2RAY_SERVER_IP")
        }
        
        // Obtener puerto destino (DPORT o 443 por defecto para v2ray)
        targetPort := os.Getenv("DPORT")
        if targetPort == "" {
                targetPort = "443"
        }
        
        targetAddr := targetHost + ":" + targetPort
        
        log.Printf("🚀 Proxy iniciado - Escuchando en: %s", listenAddr)
        log.Printf("🎯 Redirigiendo a: %s", targetAddr)
        
        ln, err := net.Listen("tcp", listenAddr)
        if err != nil {
                log.Fatalf("❌ Error al escuchar: %v", err)
                return
        }
        
        log.Printf("✅ Proxy listo y escuchando conexiones...")
        
        for {
                conn, err := ln.Accept()
                if err != nil {
                        log.Printf("⚠️ Error al aceptar conexión: %v", err)
                        continue
                }
                go handleConnection(conn, targetAddr)
        }
}
func handleConnection(src net.Conn, targetAddr string) {
        clientAddr := src.RemoteAddr().String()
        log.Printf("📥 Nueva conexión TCP desde: %s → %s", clientAddr, targetAddr)
        
        dst, err := net.Dial("tcp", targetAddr)
        if err != nil {
                log.Printf("❌ Error al conectar con destino %s: %v", targetAddr, err)
                src.Close()
                return
        }
        
        log.Printf("✅ Conexión establecida: %s ↔ %s", clientAddr, targetAddr)
        startTime := time.Now()
        
        // Copiar datos en ambas direcciones
        done := make(chan bool, 2)
        
        go func() {
                io.Copy(dst, src)
                done <- true
        }()
        
        go func() {
                io.Copy(src, dst)
                done <- true
        }()
        
        // Esperar a que termine una de las copias
        <-done
        
        duration := time.Since(startTime)
        log.Printf("🔌 Conexión cerrada: %s (duración: %v)", clientAddr, duration)
        
        src.Close()
        dst.Close()
}

