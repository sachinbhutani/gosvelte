<script lang="ts">
import { compute_rest_props } from 'svelte/internal';

  import logo from './assets/svelte.png'
  import Counter from './lib/Counter.svelte'
  export let props = JSON.parse($$props)
  console.log(props)
</script>
<main>
  <img src={logo} alt="Svelte Logo" />
  <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" alt="Go Logo">
  <h1>Svelte served on GO!</h1>

  <Counter />
  <p>SSR Text: {props.Title}</p>
  <p>Server: {props.Server}</p>
</main>

<style>
  :root {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  }

  main {
    text-align: center;
    padding: 1em;
    margin: 0 auto;
  }

  img {
    height: 16rem;
    width: 16rem;
  }

  h1 {
    color: #0077ff;
    text-transform: uppercase;
    font-size: 4rem;
    font-weight: 100;
    line-height: 1.1;
    margin: 2rem auto;
    max-width: 14rem;
  }

  p {
    max-width: 14rem;
    margin: 1rem auto;
    line-height: 1.35;
  }

  @media (min-width: 480px) {
    h1 {
      max-width: none;
    }

    p {
      max-width: none;
    }
  }
</style>
