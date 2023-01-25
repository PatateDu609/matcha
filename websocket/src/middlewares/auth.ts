import {ExtendedError} from "socket.io/dist/namespace";
import {Socket} from "socket.io";
import {parse} from "cookie";
import {log} from "../log";
import {createRouter} from "../axios";
import {Logger} from "tslog";

async function getSessionUUID(subLogger: Logger<unknown>, sid: string): Promise<string> {
	return new Promise((resolve, reject) => {
		const router = createRouter(sid)

		router.get("/session/uuid").then(result => {
			const data: { uuid: string } = JSON.parse(result.data)
			subLogger.info(`got uuid: ${data.uuid}`)

			if (data && data.uuid)
				resolve(data.uuid)
			else
				reject("got invalid uuid")
		}).catch(reason => {
			subLogger.error(`couldn't get uuid for \`${sid}\`: ${reason}`)
			reject(reason)
		})
	})
}

export function auth(namespace: string): (socket: Socket, next: (err?: ExtendedError) => void) => void {
	return (socket: Socket, next: (err?: ExtendedError) => void) => {
		const val = socket.handshake.headers.cookie
		const subLogger = log.getSubLogger({name: namespace})

		const isNotAuth = new Error("user is not auth")

		subLogger.silly(`cookies: ${val}`)

		if (val == undefined) {
			subLogger.error("user is not authenticated")
			next(isNotAuth)
			return
		}
		const cookies = parse(val)
		for (let cookiesKey in cookies) {
			subLogger.debug(` - ${cookiesKey} -> ${cookies[cookiesKey]}`)
		}

		const sid = cookies['gosessid']
		subLogger.info(`sid: ${sid}`)
		if (sid == '') {
			subLogger.error("user is not authenticated")
			next(isNotAuth)
			return
		}

		getSessionUUID(subLogger, sid).then(uuid => {
			socket.handshake.auth['user'] = uuid
			next()
		}).catch(reason => {
			subLogger.error(`error while checking session: ${reason}`)
			subLogger.error("user is not authenticated")
			next(isNotAuth)
		})
	}
}