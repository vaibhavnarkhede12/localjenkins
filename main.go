package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// ddddddeclarationssssss
type Databin struct {
	Id        int     `json:"id"`
	Quantity  int     `json:"quantity"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Filledpct int     `json:"filledpct"`
}

type Data struct {
	Slice     []Databin `json:""`
	Maptojson string    `json:""`
	Datimemap string    `json:""`
}

var Dataarray []Databin

var addConnection = make(chan *websocket.Conn)
var newHttpData = make(chan string)

var wsConnMap = make(map[int]*websocket.Conn)
var prevGarbageData = make(map[string][]int)    ///
var prevDatimerData = make(map[string][]string) //
var wsConnSrNo = 0

////////////////////////////////
func main() {
	// router := mux.NewRouter()
	go mapNewConnections()
	go sendToAllConnections()

	router := mux.NewRouter()
	router.HandleFunc("/", rxHttpFunc)
	router.HandleFunc("/ws", wsEndpoint)
	http.ListenAndServe(":9000", router)

	// http.HandleFunc("/", rxHttpFunc)
	// http.HandleFunc("/ws", wsEndpoint)
	// http.ListenAndServe(":9000", nil)
	defaultDataBin := Databin{
		Id:        999,
		Quantity:  999,
		Lat:       999,
		Lng:       999,
		Filledpct: 99,
	}
	Dataarray = append(Dataarray, defaultDataBin)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	maptojson, _ := json.Marshal(prevGarbageData)
	datimemap, _ := json.Marshal(prevDatimerData)
	JSONDataarray := Data{
		Slice:     Dataarray,
		Maptojson: string(maptojson),
		Datimemap: string(datimemap),
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	conn.WriteJSON(JSONDataarray)

	// maptojson, _ := json.Marshal(prevGarbageData)
	// conn.WriteJSON(maptojson)

	// fmt.Println(JSONDataarray)
	// fmt.Println("Data sent to Client")
	// go mapNewConnections
	addConnection <- conn
	// reader(conn)
}

func mapNewConnections() {
	for value := range addConnection {
		fmt.Println("mapped a new websocket Connection")
		wsConnMap[wsConnSrNo] = value
		wsConnSrNo++
	}
}

func sendToAllConnections() {
	// isThereNewData := <-newHttpData
	for value := range newHttpData {
		if value != "" {
			// fmt.Println("ready to send new data for refresh")
		}

		if wsConnMap != nil {
			JSONDataarrayRefreshed := Data{
				Slice: Dataarray[len(Dataarray)-1 : len(Dataarray)],
			}

			for k := range wsConnMap {
				// log.Println(len(wsConnMap))
				err := wsConnMap[k].WriteJSON(JSONDataarrayRefreshed)
				if err != nil {
					// fmt.Println("error in reffreshing the data")
				} else {
					// fmt.Println("data refreshed")
				}
			}
		}
	}
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		// fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func rxHttpFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "POST":
		fmt.Println("recieving some http Data")
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
		var databin Databin
		json.Unmarshal([]byte(reqBody), &databin)
		flag := 0
		for i, Itr := range Dataarray {
			if Itr.Id == databin.Id {
				flag = 1
				copy(Dataarray[i:], Dataarray[i+1:]) // Shift a[i+1:] left one index.
				// Dataarray[len(Dataarray)-1] = {}   // Erase last element (write zero value).
				Dataarray = Dataarray[:len(Dataarray)-1]
				num := databin.Id
				idstr := "id" + strconv.Itoa(num)
				prevGarbageData[idstr] = append(prevGarbageData[idstr], databin.Quantity)

				dt := time.Now()
				str := dt.Format("01-02-2006 15:04:05")
				prevDatimerData[idstr] = append(prevDatimerData[idstr], str)

			}
		}
		if flag == 0 {
			// fmt.Println("updated!!!!!!")
			num := databin.Id
			idstr := "id" + strconv.Itoa(num)
			prevGarbageData[idstr] = make([]int, 1)
			prevGarbageData[idstr] = append(prevGarbageData[idstr], databin.Quantity)

			dt := time.Now()
			str := dt.Format("01-02-2006 15:04:05")
			prevDatimerData[idstr] = make([]string, 1)
			prevDatimerData[idstr] = append(prevDatimerData[idstr], str)

		} else {
			flag = 0
		}
		fmt.Println("updated map")
		fmt.Println(prevGarbageData)
		// Maptojson, _ := json.Marshal(prevGarbageData)
		// fmt.Println(string(maptojson))
		Dataarray = append(Dataarray, databin)
		// fmt.Println(Dataarray)
		newHttpData <- "newDataForRefresh"

		// for _, Itr := range Dataarray {
		// 	fmt.Println(Itr.Id, Itr.Quantity, Itr.City)
		// }

	default:
		fmt.Println("deafulted http ")
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}
