package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
)

type myRowsEventHandler struct {
}

func (h *myRowsEventHandler) Do(e *canal.RowsEvent) error {
	for i := 0; i < len(e.Table.Columns); i++ {
    	fmt.Printf(e.Table.Columns[i].Name)
	}

	fmt.Println("%s %s %s %v %v\n", e.Action, e.Table.Schema,e.Table.Name, e.Table.Columns[0].Name, e.Rows)
	return nil
}

func (h *myRowsEventHandler) String() string {
	return "myRowsEventHandler"
}

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "172.17.10.38:3306"
	cfg.User = "canal"
	cfg.Password = "canal"
	cfg.ServerID = 2345
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"t_user"}

	c, _ := canal.NewCanal(cfg)

	// Register a handler to handle RowsEvent
	rowsEventHandler := &myRowsEventHandler{}
	c.RegRowsEventHandler(rowsEventHandler)

	// Start canal
	c.Start()
	for {
		select {}
	}
}
