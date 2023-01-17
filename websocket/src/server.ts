import {Server} from "socket.io";

const io = new Server(3000, {
	path: '/'
});

io.on("connection", (socket) => {
	console.log("a user connected...")
	socket.on("disconnect", args => {
		console.log(`user has left (${args}`)
	})

	socket.on("chat message", function (message) {
		console.log(message)
		socket.emit("chat message", message)
	})
});