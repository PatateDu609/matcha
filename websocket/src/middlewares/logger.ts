import {ExtendedError} from "socket.io/dist/namespace";
import {Socket} from "socket.io";
import {log} from "../log";

export function logger(namespace: string): (socket: Socket, next: (err?: ExtendedError) => void) => void {
	return (socket: Socket, next: (err?: ExtendedError) => void) => {
		const subLogger = log.getSubLogger({name: namespace})
		if (socket.request.headers["sec-websocket-key"])
			subLogger.info(`received connection identified by: ${socket.request.headers["sec-websocket-key"]}`)
		else
			subLogger.info("received connection")

		subLogger.trace(`socket has ID = ${socket.id}`)

		next()
	}
}