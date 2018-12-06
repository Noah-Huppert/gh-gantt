import AppHeader from "./appHeader"

import HomePage from "./homePage"
import LoginPage from "./loginPage"
import LoginCallbackPage from "./loginCallbackPage"

// All components which can be used in Vue templates
export const components = {
	AppHeader,

	HomePage,
	LoginPage,
	LoginCallbackPage
}

// Page routes
export const HomePageRoute = "/"
export const LoginPageRoute = "/login"
export const LoginCallbackPageRoute = "/auth/github"

export const routes = [
	{
		path: HomePageRoute,
		component: HomePage
	},
	{ 
		path: LoginPageRoute,
		component: LoginPage,
	},
	{
		path: LoginCallbackPageRoute,
		component: LoginCallbackPage,
	}
]
