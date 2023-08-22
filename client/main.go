package main

import (
        "bufio"
        "fmt"
        "io"
        "log"
        "net"
        "os"
        "strings"
)

func logFatal(err error) {
        if err != nil {
                log.Fatal(err)
        }
}

func main() {
        connection, err := net.Dial("tcp", "localhost:8080")
        logFatal(err)

        defer connection.Close()

        fmt.Println("Enteryour username : ")

        reader := bufio.NewReader(os.Stdin)
        username, err := reader.ReadString('\n')

        logFatal(err)

        username = strings.Trim(username, " \r\n")

        welcomeMsg := fmt.Sprintf("Welcome %s, say hi to your friend.\n", username)

        fmt.Println(welcomeMsg)

        go read(connection)
        write(connection, username)

}

func read(connection net.Conn) {
        for {
                reader := bufio.NewReader(connection)
                message, err := reader.ReadString('\n')

                if err == io.EOF {
                        connection.Close()
                        fmt.Println("Connection closed.")
                        os.Exit(0)
                }

                fmt.Println(message)
                fmt.Println("---------------------------------")

        }
}

func write(connection net.Conn, username string) {
        for {

                reader := bufio.NewReader(os.Stdin)
                message, err := reader.ReadString('\n')

                if err != nil {

                        break
                }

                // username: - massage

                message = fmt.Sprintf("%s:- %s\n", username, strings.Trim(message, " \r\n"))
                connection.Write([]byte(message))

        }
}
