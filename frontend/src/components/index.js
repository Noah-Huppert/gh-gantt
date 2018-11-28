import HelloComponent from "./hello"

export const components = {
	HelloComponent
}

export const routes = (store) => {
	return [
		{ 
			path: "/",
			component: HelloComponent,
			props: {
				store: store
			}
		}
	]
}
