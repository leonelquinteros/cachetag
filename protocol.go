package cachetag

import (
    "net"
    "strings"
)

// Define public storage
var ct_storage map[string]string

func process_command(input string, conn net.Conn) string {
    var response, command, key, data string
    var parsed []string

    parsed = strings.SplitN(input, " ", 3)

    if len(parsed) > 0 {
        command = parsed[0]

        if len(parsed) > 1 {
            key = parsed[1]

            if len(parsed) > 2 {
                data = parsed[2]
            }
        }

        switch command {
        case "GET":
            response = cachetag_get(key)
        case "SET":
            cachetag_set(key, data)
        case "DEL":
            cachetag_del(key)
        case "EXIT":
            response = "EXIT"
        default:
            response = "ERROR"
        }
    }

    return response
}

func cachetag_get(key string) string {
    value := ct_storage[key]

    return value
}

func cachetag_set(key, data string) {
    ct_storage[key] = data
}

func cachetag_del(key string) {
    delete(ct_storage, key)
}
