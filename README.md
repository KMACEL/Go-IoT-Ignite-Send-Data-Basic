## Go-IoT-Ignite-Send-Data-Basic
Simple program that only sends data to the "IoT-Ignite" platform with the "go" language

### First Step : 

Go Install = https://golang.org/dl/


### Second Step :

 GOBOT - Mqtt Install = Open Terminal.  

```
go get -d -u gobot.io/x/gobot/... && go install gobot.io/x/gobot/platforms/mqtt
```
https://gobot.io/documentation/platforms/mqtt/


### Third Step :
Open Terminal. 

```
go run GO_IoT_Ignite_Send_Message.go
```



### IoT - Ignite Connection Topic :

```
topic := "golang@gopher@mqtt/publish/DeviceProfile/goNode/goSensor"
```

```
topic:="<Gateway ID>/publish/DEviceProfile/<Node ID>/<Sensor ID>"
```


### IoT - Ignite Broker Host :
```
broker := "ssl://mqtt.ardich.com:8883"
```
8883 = SSL



### IoT - Ignite Client ID :
```
clientId := "golang@gopher@mqtt"
```
```
clientId:="<Device ID>"
```



"Client Id" is your Gateway Id(Device ID)

### IoT - Ignite User Name :
```
user := "gophergo"
```
```
user := "<User Name>"
```
"User Name" is your Gateway User Name 



### IoT - Ignite Password :
```
password := "12345678"
```
```
password := "<Password>"
```
"Password" is your Gateway Password
