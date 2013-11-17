package main

import "strings"

func process_command(comm string) string {
    var response, command, value string
    var parsed []string

    parsed = strings.SplitN(comm, " ", 2)

    if len(parsed) > 0 {
        command = parsed[0]

        if len(parsed) > 1 {
            value = parsed[1]
        }

        switch command {
            case "GET":
                response = cachetag_get(value)
            case "SET":
                response = cachetag_set(value)
            case "DEL":
                response = cachetag_del(value)
        }
    }

    return response
}

func cachetag_get(value string) string {
    return value
}

func cachetag_set(value string) string {
    return value
}

func cachetag_del(value string) string {
    return value
}
