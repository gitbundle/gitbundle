<script>
  import { beforeUpdate, onMount, tick } from 'svelte'
  import Svg from './Svg.svelte'
  import Typeahead from './WrappedTypeahead.svelte'
  import { htmlEscape } from 'escape-goat'

  export let pageData

  let comboboxRef,
    searchRef,
    data = {
      csrfToken: window.config.csrfToken,
      appSubUrl: window.config.appSubUrl,
      active: false,
      ...pageData
    },
    items = [],
    query = '',
    newDependency = 0,
    selected,
    issueSearchUrl

  let exists = JSON.parse(data.exists)
  exists.push(Number(data.issueID))

  //issues/search?q=test&priority_repo_id=17&type=all
  $: issueSearchUrl = `${data.appSubUrl}/${data.repolink}/issues/search?q=${query}&type=${data.issueType}`
  $: if (data.crossRepoSearch === 'true') {
    issueSearchUrl = `${data.appSubUrl}/issues/search?q=${query}&priority_repo_id=${data.repoid}&type=${data.issueType}`
  }

  const getItems = async () => {
    const response = await fetch(issueSearchUrl)
    // console.log(response)
    if (response.ok) {
      items = await response.json()
      // console.log(items)
    }
  }

  const handleSelect = ({ detail }) => {
    newDependency = detail.original.id
    selected = detail.original
  }

  const handleClear = (e) => {
    // console.log(e)
    newDependency = 0
    selected = null
    query = ''
  }

  //TODO: do we need this for realtime search?
  const handleType = async (e) => {
    // console.log(e)
    // query = e.detail
    // await getItems()
  }

  const extract = (item) => {
    // #${issue.number} ${htmlEscape(issue.title)}<div class="text small dont-break-out">${htmlEscape(issue.repository.full_name)}</div>
    return htmlEscape(item.title)
  }

  beforeUpdate(() => {
    // console.log(showCreateNewBranch, data.searchTerm, data.active)
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

<form
  on:submit={(e) => {
    if (!newDependency || newDependency === data.issueID) {
      console.log('Can not add self or empty dependency')
      e.preventDefault()
    }
  }}
  method="POST"
  action={data.action}
  class="mt-2 w-full"
>
  <input type="hidden" name="_csrf" value={data.csrfToken} />
  <input name="newDependency" type="hidden" bind:value={newDependency} />
  <div bind:this={comboboxRef} class="dropdown w-full" class:dropdown-open={data.active}>
    <div class="btn-group w-full">
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <div
        tabindex="0"
        class="btn-outline btn-sm btn flex w-[calc(100%-2rem)] items-center"
        on:click={() => {
          searchRef?.focus()
          data.active = true
        }}
      >
        <span class="line-clamp-1 flex flex-1 p-0 text-left text-base-content">
          {#if selected}
            <div class="flex flex-col">
              <div class="line-clamp-1 flex items-center">
                #{selected.number}
                {selected.title}
              </div>
              <div class="line-clamp-1 flex items-center text-xs">
                {selected.repository.full_name}
              </div>
            </div>
          {:else}
            {data.placeholder}
          {/if}
        </span>
        <Svg name="octicon-triangle-down" size="16" className="flex w-4 h-4 text-base-content" />
      </div>
      <button
        class="btn-primary btn-sm btn w-[2rem] p-0"
        on:click={() => {
          searchRef?.blur()
          data.active = false
        }}
      >
        <Svg name="octicon-plus" size="16" className="w-4 h-4" />
      </button>
    </div>
    <div class="card dropdown-content card-compact w-full p-2" class:hidden={!data.active}>
      <div class="card-body !p-0">
        <Typeahead bind:searchRef {handleSelect} {handleClear} {handleType} showAllResultsOnOpen={true} data={items} {extract} disable={(item) => exists.includes(item.id)} className="!pr-8 !rounded-lg !input !input-sm !bg-base-200 placeholder:text-xs">
          <div slot="typeahead-result" let:result class="flex flex-col">
            <div class="line-clamp-1">
              #{result.original.number}
              {result.original.title}
            </div>
            <div class="line-clamp-1 flex items-center text-xs">
              {result.original.repository.full_name}
            </div>
          </div>
          <span slot="typeahead-right">
            <Svg name="octicon-search" size="16" className="m-2 h-4 w-4 text-base-content/60" />
          </span>
        </Typeahead>
      </div>
    </div>
  </div>
</form>
