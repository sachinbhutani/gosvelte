import App from './App.svelte'
const app = new App({
  target: document.getElementById('app'),
  props:  document.getElementById('app').getAttribute('props')
})

export default app
