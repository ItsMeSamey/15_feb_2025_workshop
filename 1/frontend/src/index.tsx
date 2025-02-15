/* @refresh reload */
import { render } from 'solid-js/web'
import './index.css'

render(() => {
  document.body.setAttribute('data-kb-theme', 'dark')
  return (
    <h1 class='text-5xl w-full text-center bg-red-500'> Hi </h1>
  )
}, document.body)

