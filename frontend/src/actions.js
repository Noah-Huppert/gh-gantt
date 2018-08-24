import { OAUTH_CALLBACK } from "./actionTypes";

export const oauthCallback = tempOAuthCode => {
	return {
		type: OAUTH_CALLBACK,
		tempOAuthCode: tempOAuthCode
	}
}
