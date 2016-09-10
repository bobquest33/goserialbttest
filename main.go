package main

import (
        "log"
		"os"
        "github.com/tarm/serial"
)

func main() {
		serial_port := os.Args[2]
        c := &serial.Config{Name: serial_port, Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                log.Fatal(err)
        }
		defer s.Close()
        _, err = s.Write([]byte(os.Args[1]))
        if err != nil {
                log.Fatal(err)
        }

        
		buf := make([]byte, 128)
        n, err := s.Read(buf)
        if err != nil {
                log.Fatal(err)
        }
        log.Printf("%q", buf[:n])
		
}