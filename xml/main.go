package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `XML:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	// XMLName    xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main2() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"hogehoge_vps", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"piyopiyo_vps", "127.1.0.1"})
	v.Svs = append(v.Svs, server{"fugafuga_vps", "127.2.0.1"})
	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("error: %n\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

}
