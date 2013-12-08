package cachetag

import (
    "bufio"
    "net"
    "net/textproto"
)

func StartServer(config Configuration) {
    // Start listening
    ln, err := net.Listen("tcp", config.address+":"+config.port)
    if err != nil {
        // handle error
        out(err.Error())
    }

    // Init storage
    ct_storage = make(map[string]string)

    // Close listener at the end.
    defer ln.Close()

    // Accept connections
    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
            out(err.Error())

            continue
        }

        // Handle the connection in a new goroutine.
        // The loop then returns to accepting, so that
        // multiple connections may be served concurrently.
        go handle_connection(conn)
    }
}

func handle_connection(conn net.Conn) {
    // Shut down the connection at the end
    defer conn.Close()

    // Reader and Writer for this connection
    reader := textproto.NewReader(bufio.NewReader(conn))
    writer := textproto.NewWriter(bufio.NewWriter(conn))

    // Loop through commands
READ:
    for {
        // Read
        input, err := reader.ReadLine()
        if err != nil {
            // handle error
            out(err.Error())

            // keep reading
            continue
        }

        // Exec command and get response.
        response := process_command(input, conn)

        // Process responses
        switch response {
        case "":
            continue
        case "EXIT":
            break READ
        case "ERROR":
            out("ERROR: " + input)
        default:
            // Return response
            writer.PrintfLine(response)
        }
    }
}
