<script>
  import { beforeUpdate, onMount, tick } from 'svelte'
  import Svg from './Svg.svelte'
  import Typeahead from './WrappedTypeahead.svelte'
  import { htmlEscape } from 'escape-goat'

  export let pageData

  let comboboxRef,
    searchRef,
    data = {
      appSubUrl: window.config.appSubUrl,
      active: false,
      ...pageData
    },
    items = [],
    query = '',
    selected,
    repoSearchUrl

  $: repoSearchUrl = `${data.appSubUrl}/repo/search?q=${query}&limit=20`

  const getItems = async () => {
    const response = await fetch(repoSearchUrl)
    // console.log(response)
    if (response.ok) {
      const { data } = await response.json()
      items = data
    }
  }

  const handleSelect = ({ detail }) => {
    // console.log(detail)
    selected = detail.original
    comboboxRef.closest('form').setAttribute('action', `${data.appSubUrl}/${detail.original.full_name}/issues/new`)
    data.active = false
  }

  const handleClear = (e) => {
    // console.log(e)
    selected = null
    query = ''
  }

  //TODO: do we need this for realtime search?
  const handleType = async (e) => {
    // console.log(e)
    query = e.detail
    // await getItems()
  }

  const extract = (item) => {
    // #${issue.number} ${htmlEscape(issue.title)}<div class="text small dont-break-out">${htmlEscape(issue.repository.full_name)}</div>
    return htmlEscape(item.full_name)
  }

  beforeUpdate(() => {
    // console.log(data.active)
  })

  onMount(async () => {
    await getItems()
  })
</script>

<svelte:window
  on:click={({ target }) => {
    if (!comboboxRef?.contains(target)) {
      data.active = false
    }
  }}
/>

<div bind:this={comboboxRef} class="dropdown w-full" class:dropdown-open={data.active}>
  <!-- svelte-ignore a11y-label-has-associated-control -->
  <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <label
    tabindex="0"
    class="btn-outline btn flex items-center"
    on:click={() => {
      searchRef?.focus()
      data.active = true
    }}
  >
    <span class="line-clamp-1 flex flex-1 p-0 text-left text-base-content">
      {#if selected}
        {selected.full_name}
      {:else}
        {data.selected}
      {/if}
    </span>
    <Svg name="octicon-triangle-down" size="16" className="flex w-4 h-4 text-base-content" />
  </label>
  <div class="card dropdown-content card-compact w-full p-2" class:hidden={!data.active}>
    <div class="card-body !p-0">
      <Typeahead bind:searchRef {handleSelect} {handleClear} {handleType} showAllResultsOnOpen={true} data={items} {extract} className="!pr-8 !rounded-lg !input !input-sm !bg-base-200">
        <div slot="typeahead-result" let:result class="flex flex-col">
          <div class="line-clamp-1 flex items-center">
            {result.original.full_name}
          </div>
        </div>
        <span slot="typeahead-right">
          <Svg name="octicon-search" size="16" className="m-2 h-4 w-4 text-base-content/60" />
        </span>
      </Typeahead>
    </div>
  </div>
</div>
