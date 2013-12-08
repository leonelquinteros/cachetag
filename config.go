package cachetag

type Configuration struct {
    address string
    port    string
}

func GetConfig() Configuration {
    conf := Configuration{
        "",
        "23481",
    }

    return conf
}
