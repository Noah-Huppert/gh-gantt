import { HELLO } from "./actionTypes";

export const hello = text => {
	return {
		type: HELLO,
		hello: text
	}
}
