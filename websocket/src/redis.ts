import {createClient, RedisClientType} from "redis";

export function NewClient(): RedisClientType {
	return createClient({
		url: "redis://localhost:6379/4",
	})
}