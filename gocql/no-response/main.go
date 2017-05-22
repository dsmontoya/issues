package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	"github.com/entropyx/tools/intutils"
	"github.com/gocql/gocql"
)

func main() {
	var session *gocql.Session
	var err error
	for {
		ips, err := lookupHost(os.Getenv("CASSANDRA_HOST"))
		if err != nil {
			fmt.Println("err", err)
			continue
		}
		cluster := gocql.NewCluster(ips)
		cluster.Consistency = gocql.Any
		session, err = cluster.CreateSession()
		if err != nil {
			fmt.Println("err", err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	b, err := ioutil.ReadFile("schema.cql")
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt := string(b)
	for _, query := range strings.Split(stmt, ";") {
		fmt.Println("=====================")
		fmt.Println(query)
		err = session.Query(query).Exec()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func lookupHost(domain string) (string, error) {
	ips, err := net.LookupHost(domain)
	fmt.Println("ips", ips)
	if err != nil {
		return "", err
	}
	r := intutils.RandInt(0, len(ips))
	return ips[r], nil
}
