package main
import (
        "io"
        "net"
        "os"
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
        ln, err := net.Listen("tcp", listenAddr)
        if err != nil {
                return
        }
        for {
                conn, err := ln.Accept()
                if err != nil {
                        continue
                }
                go handleConnection(conn, targetAddr)
        }
}
func handleConnection(src net.Conn, targetAddr string) {
        dst, err := net.Dial("tcp", targetAddr)
        if err != nil {
                src.Close()
                return
        }
        go io.Copy(dst, src)
        go io.Copy(src, dst)
}

