package main

type Configuration struct {
    address string
    port    string
}

func get_config() Configuration {
    conf := Configuration{
                "",
                "23481",
    }

    return conf
}
