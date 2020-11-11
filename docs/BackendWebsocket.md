# Websocket
```
/message/{socketId}?token=<token>
```
# Backend Websocket test
```
wscat -c ws://localhost:5000/message/1{socketId}?token=<token>
```
# Event Type
## 1. Send Message 
- Client send message in group

**Request**
```json
{
    "type":"send_text",
    "data":{
             "groupId": <idgroup>,
             "body":"<message>"
           }
}
```
**Response**
```json
{
"type":"send_text",
"data":{
        "groupId":1,
        "id":63,
        "body":"hello",
        "Sender":"893a4692-63bb-4919-80d9-aece678c0422",
        "socketId":"",
        "idContinueOldMess":0,
        "createdAt":"2020-11-11T15:04:24.800875Z",
        "updatedAt":"2020-11-11T15:04:24.800875Z"
       }
}
```
## 2. Subcribe Group 
- Client join group and receive history message (20 message)

**Request**
```json
{
    "type":"subcribe_group",
    "data":{
            "groupId": <idgroup>,
            "socketId":"<socketId>"
           }
}
```
**Response**
```json
{
    "type":"subcribe_group",
    "data":{"groupId":1,
            "id":62,
            "body":"hello",
            "Sender":"ffb63922-8f99-46ba-9648-d07f3ac14757",
            "socketId":"",
            "idContinueOldMess":0,
            "createdAt":"2020-11-11T14:37:06.818823Z",
            "updatedAt":"2020-11-11T14:37:06.818823Z"}
}
```
## 3. Load Old Message
- Client receive continue history message before idContinueOldMess  (20 message)

**Request**
```json
{ 
    "type":"load_old_mess",
    "data":{
            "groupId": <idgroup>,
            "socketId":"<socketId>",
            "idContinueOldMess": <idOldMess>
            }
}
```
**Response**
```json
{
    "type":"load_old_mess",
    "data":{"groupId":1,
            "id":42,"body":"hello",
            "Sender":"893a4692-63bb-4919-80d9-aece678c0422",
            "socketId":"",
            "idContinueOldMess":0,
            "createdAt":"2020-11-04T18:12:47.907925Z",
            "updatedAt":"2020-11-04T18:12:47.907925Z"
            }
}
```