let socket;

function init() {
    socket = new WebSocket("ws://localhost:3000/ws");
    console.log("Socket connection", socket);
    socket.onmessage = function (event) {
        console.log("Inside event", event);
        const messageList = document.getElementById("messages");
        const messageItem = document.createElement("li");
        messageItem.textContent = event.data;
        messageList.appendChild(messageItem);
    };

    const form = document.getElementById("chatForm");
    form.onsubmit = function (event) {
        event.preventDefault();
        const input = document.getElementById("messageInput");
        socket.send(input.value);
        console.log("Message sent", input.value);
        input.value = '';
    };
}

window.onload = init;