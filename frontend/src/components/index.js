import AppHeader from "./appHeader"

import HomePage from "./homePage"
import GHLoginPage from "./ghLoginPage"
import GHLoginCallbackPage from "./ghLoginCallbackPage"
import ZenHubLoginPage from "./zenhubLoginPage"

// All components which can be used in Vue templates
export const components = {
	AppHeader,

	HomePage,
	GHLoginPage,
	GHLoginCallbackPage
}

// Page routes
export const HomePageRoute = "/"
export const GHLoginPageRoute = "/auth/github"
export const GHLoginCallbackPageRoute = "/auth/github_callback"
export const ZenHubLoginPageRoute = "/auth/zenhub"

export const routes = [
	{
		path: HomePageRoute,
		component: HomePage
	},
	{ 
		path: GHLoginPageRoute,
		component: GHLoginPage
	},
	{
		path: GHLoginCallbackPageRoute,
		component: GHLoginCallbackPage
	},
	{
		path: ZenHubLoginPageRoute,
		component: ZenHubLoginPage
	}
]
