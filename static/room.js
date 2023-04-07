
const roomWebsocketUrl = `ws://localhost:1323/ws/rooms/${roomId}`;

const chatLog = document.getElementById("chatLog");
const chatText = document.querySelector("input[name='chatText']");
const sendButton = document.getElementById("sendButton")

const roomWebsocket = new WebSocket(roomWebsocketUrl);

roomWebsocket.addEventListener("open", (event) => {
    roomWebsocket.send(`${currentUser} just joined!`);
})

roomWebsocket.addEventListener("message", (event) => {
    chatLog.value += event.data + '\n'
})

roomWebsocket.addEventListener("error", (event) => {
    console.log(event)
})


function sendMessage(){
    roomWebsocket.send(`${currentUser}: ${chatText.value}`);
    chatText.value = "";
}

chatText.addEventListener("keyup", (event) => {
    if(event.key == "Enter"){
        sendMessage();
    }
})
sendButton.addEventListener("click", sendMessage)
chatText.focus()