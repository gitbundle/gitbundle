<script>
  import { onMount, beforeUpdate, tick } from 'svelte'
  // import Typeahead from 'svelte-typeahead'
  import Typeahead from './SvelteTypeahead.svelte'

  // const dispatch = createEventDispatcher()
  // const defaultHandleClear = (e) => {
  //   // [...e, { event: 'clear' }]

  //   console.log(e)
  // }

  // const defaultHandleKeydown = (e) => {
  //   if ((e.keyCode === 75 && e.metaKey) || (e.keyCode === 75 && e.ctrlKey)) {
  //     e.preventDefault()
  //     searchbox.querySelector('input[type=search]').focus()
  //     dispatch('focus')
  //   }
  // }

  export let placeholder = 'Search...'
  export let label = 'Search'
  export let data = [] // {title: '', content: ''}
  export let handleKeydown
  export let handleClear
  export let handleInput
  export let handleType
  export let handleChange
  export let handleSelect
  export let handleBlur
  export let handleFocus
  export let value
  export let limit
  export let className = null
  export let handleResults
  export let showAllResultsOnOpen = false
  export let searchRef = null
  export let defaultSelectedValue = null
  export let extract = (item) => {
    if (typeof item === 'object') {
      return item.content
    }
    return item
  }
  export let disable = (item) => false

  const inputWrapper = async (e, f) => {
    // e.preventDefault()
    !f || f(e)
    // ref.querySelector('input[type=search]').focus()
    // dispatch('focus')
    searchRef.focus()
  }

  if (!className) {
    className = `rounded-lg placeholder:text-xs placeholder:text-base-content focus:border-none !input !input-sm !border-base-content/10 min-w-[8rem]`
  }
</script>

<!-- <svelte:window on:keydown={handleKeydown} /> -->

<!-- svelte-ignore a11y-label-has-associated-control -->
<div class="searchbox relative h-auto w-full">
  <div class="absolute left-0 top-0 z-10 h-fit">
    <slot name="typeahead-left" />
  </div>
  <!-- prettier-ignore -->
  <Typeahead
    bind:value
    bind:searchRef
    bind:defaultSelectedValue
    autoselect={false}
    {placeholder}
    {limit}
    {label}
    {data}
    {extract}
    {disable}
    {showAllResultsOnOpen}
    class={className}
    {handleResults}
    on:type={(e) => inputWrapper(e, handleType)}
    on:input={(e) => inputWrapper(e, handleInput)}
    on:keydown={(e) => inputWrapper(e, handleKeydown)}
    on:clear={(e) => inputWrapper(e, handleClear)}
    on:change={handleChange}
    on:select={handleSelect}
    on:blur={handleBlur}
    on:focus={handleFocus}
    let:result
  >
    <slot name="typeahead-result" {result}>
      {@html result.string}
    </slot>
    <svelte:fragment slot="no-results" let:value>
      <slot name="typeahead-no-results" {value} />
    </svelte:fragment>
  </Typeahead>
  <div class="absolute right-0 top-0 z-10 h-fit">
    <slot name="typeahead-right" />
  </div>
</div>

<style lang="postcss">
  :global(.searchbox [data-svelte-typeahead]) {
    @apply h-full w-full max-w-full bg-inherit;
    /* @apply !bg-transparent; */
  }
  :global(.searchbox [data-svelte-typeahead] label),
  :global(.searchbox [data-svelte-search] label) {
    @apply !hidden;
  }
  :global(.searchbox [data-svelte-search] input) {
    /* @apply !rounded-lg !border-none !pl-8; */
    /* @apply rounded-lg placeholder:text-xs placeholder:text-base-content focus:border-none; */
    /* @apply !input !input-sm !border-base-content/10; */
    @apply !input-bordered !border-base-content/10;
  }
  :global(.searchbox [data-svelte-search] input::placeholder) {
    @apply !text-gray-500;
  }
  :global(.searchbox [data-svelte-search] input:focus) {
    /* @apply !outline-1 !outline-offset-2 !outline-base-content/20; */
    @apply !border-none !outline-1 !outline-offset-1 !outline-mako-800;
  }
  :global(.searchbox [data-svelte-typeahead] .svelte-typeahead-list) {
    /* @apply rounded-btn translate-y-2 overflow-hidden !bg-base-100; */
    @apply !static max-h-52 overflow-y-auto bg-inherit !shadow-none;
  }
  :global(.searchbox [data-svelte-typeahead] .svelte-typeahead-list .selected),
  :global(.searchbox [data-svelte-typeahead] .svelte-typeahead-list .selected:hover) {
    /* @apply !bg-neutral text-neutral-content; */
    @apply rounded-lg !bg-base-content/10;
  }
  :global(.searchbox [data-svelte-typeahead] ul.svelte-typeahead-list li) {
    /* @apply text-base-content; */
    @apply p-2 leading-5;
  }
  :global(.searchbox [data-svelte-typeahead] .svelte-typeahead-list li:hover),
  :global(.searchbox [data-svelte-typeahead] .svelte-typeahead-list li.selected:hover) {
    @apply !rounded-lg !bg-base-content/10;
  }
  :global(.searchbox [data-svelte-typeahead] .svelte-typeahead-list li:not(:last-of-type)) {
    @apply !border-none;
  }
</style>
