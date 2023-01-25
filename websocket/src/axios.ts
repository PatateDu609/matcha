import {Axios} from "axios";
import {Agent} from "http";

export function createRouter(sid: string): Axios {
	return new Axios({
		baseURL: "http://localhost:4001",
		withCredentials: true,
		httpAgent: new Agent({family: 4}),
		headers: {
			"Origin": "http://localhost:3000",
			"Cookie": `gosessid=${sid}`
		},
	})
}