let socket;


function init() {
    socket = new WebSocket("ws://localhost:3000/ws");
    socket.onmessage = function (event) {
        const messageData = JSON.parse(event.data);
        const messageList = document.getElementById("messages");
        const messageItem = document.createElement("li");
        const messageContent = document.createElement("p");
        const senderId = parseInt(document.getElementById("senderIdInput").value);
        // Set the text content of the message
        messageContent.textContent = messageData.content;

        // Set the common classes for the message content
        
        // Check if the message is from the sender or receiver
        if (messageData.senderId == senderId) {
            messageContent.className = "bg-background px-3 py-3 text-end max-w-max rounded-lg text-sm";
            messageItem.className = "flex items-center justify-end w-full px-4";
        } else {
            messageContent.className = "bg-accent px-3 py-3 text-end max-w-max rounded-lg text-sm";
            messageItem.className = "flex items-center justify-start w-full px-4";
        }

        // Append the message content to the list item
        messageItem.appendChild(messageContent);

        // Append the list item to the message list
        messageList.appendChild(messageItem);
    };

    const form = document.getElementById("chatForm");
    form.onsubmit = function (event) {
        event.preventDefault();
        const senderId = parseInt(document.getElementById("senderIdInput").value);
        const receiverId = parseInt(document.getElementById("receiverIdInput").value);
        const input = document.getElementById("messageInput");
        const message = {
            content: input.value,
            senderId: senderId,
            receiverId: receiverId,
            createdAt: new Date().toISOString()
        }
        socket.send(JSON.stringify(message));
        input.value = '';
    };
}

window.onload = init;