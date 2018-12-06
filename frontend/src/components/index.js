import AppHeader from "./appHeader"

import HomePage from "./homePage"
import GHLoginPage from "./ghLoginPage"
import GHLoginCallbackPage from "./ghLoginCallbackPage"

// All components which can be used in Vue templates
export const components = {
	AppHeader,

	HomePage,
	GHLoginPage,
	GHLoginCallbackPage
}

// Page routes
export const HomePageRoute = "/"
export const GHLoginPageRoute = "/login"
export const GHLoginCallbackPageRoute = "/auth/github"

export const routes = [
	{
		path: HomePageRoute,
		component: HomePage
	},
	{ 
		path: GHLoginPageRoute,
		component: GHLoginPage,
	},
	{
		path: GHLoginCallbackPageRoute,
		component: GHLoginCallbackPage,
	}
]
