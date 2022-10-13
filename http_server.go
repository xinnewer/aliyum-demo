package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/thb-cmyk/aliyum-demo/databasic"
)

type requestresponse struct {
	resp http.ResponseWriter
	requ *http.Request
	wg   *sync.WaitGroup
}

/*
the funciton aims to initialize a http server, which receive http request and send a response to the http client.
*/
func IntrefaceInit() {
	http.HandleFunc("/voltage", voltageHandler)
	http.HandleFunc("/check_mode", checkmodeHandler)
	http.HandleFunc("/error_info", errorinfoHandler)

	http.ListenAndServe(":8080", nil)
}

/*
the function is a handler, the router route request received from client to proper handler
*/
func voltageHandler(writer http.ResponseWriter, reader *http.Request) {

	var wg sync.WaitGroup

	wg.Add(1)

	rr := requestresponse{writer, reader, &wg}

	rawnode := databasic.RawNode_create("voltage", &rr)

	databasic.Send_raw(rawnode)

	wg.Wait()

}

/*
the function is a handler, the router route request received from client to proper handler.
*/
func checkmodeHandler(writer http.ResponseWriter, reader *http.Request) {
	var wg sync.WaitGroup

	wg.Add(1)

	rr := requestresponse{writer, reader, &wg}

	rawnode := databasic.RawNode_create("check_mode", &rr)

	databasic.Send_raw(rawnode)

	wg.Wait()
}

/*
the function is a handler, the router route request received from client to proper handler
*/
func errorinfoHandler(writer http.ResponseWriter, reader *http.Request) {
	var wg sync.WaitGroup

	wg.Add(1)

	rr := requestresponse{writer, reader, &wg}

	rawnode := databasic.RawNode_create("error_info", &rr)

	databasic.Send_raw(rawnode)

	wg.Wait()
}

/*
create a processor to handle the request received from the http client.we should registry it to databasic
*/
func voltageProccesser(tasknode *databasic.TaskNode, rawnode *databasic.RawNode) bool {

	log.Print("voltageProcesser\n\r")

	rr := rawnode.Raw.(*requestresponse)

	writer := rr.resp

	reader := rr.requ

	var index int

	err := reader.ParseForm()
	if err != nil {
		index = 1
	}
	index, err = strconv.Atoi(reader.FormValue("index"))
	if err != nil {
		fmt.Print(err.Error())
	}

	result := VoltageSelect(index)

	fmt.Printf("len: %d, content: %s\n\r", len(result), result)

	writer.Write(result)

	rr.wg.Done()

	return true
}

/*
create a processor to handle the request received from the http client.we should registry it to databasic
*/
func checkmodeProccesser(tasknode *databasic.TaskNode, rawnode *databasic.RawNode) bool {
	rr := rawnode.Raw.(*requestresponse)

	writer := rr.resp

	reader := rr.requ

	var index int

	err := reader.ParseForm()
	if err != nil {
		index = 1
	}
	index, err = strconv.Atoi(reader.FormValue("index"))
	if err != nil {
		fmt.Print(err.Error())
	}

	result := CheckModeSelect(index)
	fmt.Printf("len: %d, content: %s\n\r", len(result), result)

	writer.Write(result)

	rr.wg.Done()

	return true
}

/*
create a processor to handle the request received from the http client.we should registry it to databasic
*/
func errorinfoProccesser(tasknode *databasic.TaskNode, rawnode *databasic.RawNode) bool {
	rr := rawnode.Raw.(*requestresponse)

	writer := rr.resp

	reader := rr.requ

	var index int

	err := reader.ParseForm()
	if err != nil {
		index = 1
	}
	index, err = strconv.Atoi(reader.FormValue("index"))
	if err != nil {
		fmt.Print(err.Error())
	}

	result := ErrorInfoSelect(index)

	fmt.Printf("len: %d, content: %s\n\r", len(result), result)

	writer.Write(result)

	rr.wg.Done()

	return true
}
