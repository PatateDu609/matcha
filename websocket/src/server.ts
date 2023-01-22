import {Server} from "socket.io";
import {logger} from "./middlewares/logger";
import {auth} from "./middlewares/auth";
import {log} from "./log";

const io = new Server(3000, {
	path: "/",
});

io.on("connection", (socket) => {
	const subLogger = log.getSubLogger({name: "main"}).getSubLogger({ name: "connection" })

	subLogger.info(`user ${socket.handshake.auth['user']} connected`)

	socket.on("disconnect", args => {
		console.log(`user has left (${args})`)
	})

	socket.on("chat message", function (message) {
		console.log(message)
		socket.emit("chat message", message)
	})
});

io.use(logger("main"))

io.use(auth("main"))