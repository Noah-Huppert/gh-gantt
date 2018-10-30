import { AUTH_TOKEN } from "./actionTypes";

export const authToken = authToken => {
	return {
		type: AUTH_TOKEN,
		authToken: authToken
	}
}
