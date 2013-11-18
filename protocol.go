package main

import (
    "strings"
)

func process_command(input string) string {
    var response, command, data string
    var parsed []string

    parsed = strings.SplitN(input, " ", 2)

    if len(parsed) > 0 {
        command = parsed[0]

        if len(parsed) > 1 {
            data = parsed[1]
        }

        switch command {
            case "GET":
                response = cachetag_get(data)
            case "SET":
                response = cachetag_set(data)
            case "DEL":
                response = cachetag_del(data)
            default:
                response = "ERROR"
        }
    }

    return response
}

func cachetag_get(data string) string {
    return data
}

func cachetag_set(data string) string {
    return data
}

func cachetag_del(data string) string {
    return data
}
