# Svelte served on Go! 

Proof of Concept to serve svelte generated pages on GoLang HTTP server with server data being sent as props to svelte components. 

## Build With 
 
 [Svelte](https://svelte.dev) - Cybernatically enhanced webapps

 [Vite](https://vitejs.dev) -  Next Generation Frontend Tooling

 [Fiber](https://gofiber.io) - An Express-inspired web framework written in Go.

### clone the repo
```
git clone https://github.com/sachinbhutani/gosvelte gosvelte
```

###  compile front end 
``` 
cd gosvelte/svelte
npm install 
npm build
```

### start fiber server
```
cd gosvelte/go 
go run go/server.go
```

# Nuts and Bolts 
GoFiber can be used to render static html pages with javascript code 
As svelte is combined to pure js, JS build with svelte can be served directly on a static go server 

### Passing Props

With static served svelte pages, any data from backend needs to be requested using 'fetch' which needs an additional round-trip from client to server. 

Fiber allows to render HTML templates with data from the backend, this data can be passed on to the svelte components and hydrated on the client side. passing props during html page rendering, saves the additional round-trip to the server for fetching props. 

# Need an official Svelte framework?

Check out [SvelteKit](https://github.com/sveltejs/kit#readme), which is also powered by Vite. Deploy anywhere with its serverless-first approach and adapt to various platforms, with out of the box support for TypeScript, SCSS, and Less, and easily-added support for mdsvex, GraphQL, PostCSS, Tailwind CSS, and more.


