<script>
  import { beforeUpdate, onMount } from 'svelte'

  import { createCodeEditor } from '../modules/codeeditor.js'
  import LoadingDots from './LoadingDots.svelte'

  export let filename
  export let className
  export let content
  export let automaticLayout
  export let handleCallback

  let editorContainer,
    filenameInput,
    loading = true

  onMount(async () => {
    await createCodeEditor(editorContainer, filenameInput, undefined, className, automaticLayout)
    loading = false
  })

  const handleChange = (e) => {
    // console.log(e.target.value)
    handleCallback && handleCallback(e.target.value)
  }

  beforeUpdate(async () => {
    // console.log(content)
  })
</script>

<input bind:this={filenameInput} type="hidden" value={filename} />
<textarea bind:this={editorContainer} class="hidden" value={content} on:change={handleChange} />
{#if loading}
  <LoadingDots />
{/if}
