package main

import (
	"fmt"
	"time"
	"strconv"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)


func main() {
	topic := "python@test@mqtt/publish/DeviceProfile/PythonNode1/PySens"
	broker := "ssl://mqtt.ardich.com:8883"

	password := "12345678"
	user := "pythontest"
	clientId := "python@test@mqtt"

	qos := 1

	payload:= "{data:{sensorData:[{date:"+strconv.FormatInt(time.Now().Unix(), 10)+"000,values:[555]}],formatVersion:2}}"

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(user)
	opts.SetPassword(password)


	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Sending Data...")
	token := client.Publish(topic, byte(qos), false, payload)
	token.Wait()
		
	client.Disconnect(250)
	fmt.Println("Data Sending Completed...")
	
}
