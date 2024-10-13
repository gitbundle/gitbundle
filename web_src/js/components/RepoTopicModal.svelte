<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import Svg from './Svg.svelte'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'
  import { initRepoTopicBar } from '../modules/repo-home.js'

  export let pageData

  const { appSubUrl, csrfToken } = window.config
  const { handleToggler, handleFail, handleOk } = initRepoTopicBar()

  const maxCount = 25
  const regExp = /^[a-z0-9][a-z0-9-]{0,35}$/

  let filterText = '',
    value = pageData.items,
    items = [],
    isCountInValid = false,
    isFormatInValid = false,
    hasError = false,
    topics = ''

  // $: isCountInValid = value.length > maxCount
  // $: isFormatInValid = filterText.length > 0 && !filterText.match(regExp)
  $: hasError = isCountInValid || isFormatInValid
  $: value = value.map((item, index) => {
    isCountInValid = index + 1 > maxCount
    isFormatInValid = item.value.length > 0 && !item.value.match(regExp)
    item.error = isCountInValid || isFormatInValid
    return item
  })
  $: topics = value.map((item) => item.value).join(',')

  const handleFilter = ({ detail }) => {
    if (value?.find((i) => i.label === filterText)) return
    if (detail.length === 0 && filterText.length > 0) {
      const prev = items.filter((i) => !i.created)
      items = [...prev, { value: filterText, label: filterText, created: true }]
    }
  }

  const handleFocus = () => getItems()

  const handleChange = ({ detail }) => {
    console.log(detail)
    items = items.map((i) => {
      delete i.created
      return i
    })

    // detail.forEach((item, index) => {
    //   isCountInValid = index + 1 > maxCount
    //   isFormatInValid = item.value.length > 0 && !item.value.match(regExp)
    //   item.error = isCountInValid || isFormatInValid
    // })
  }

  const getItems = async () => {
    const response = await fetch(`${appSubUrl}/explore/topics/search?q=${filterText}`)
    if (!response.ok) {
      console.log(response)
      console.error('query topics error')
      return
    }
    const res = await response.json()
    items = []
    res.topics.forEach((item, index) => {
      if (value.includes(item.topic_name)) {
        return
      }
      items.push({ label: item.topic_name, value: item.topic_name })
    })
  }

  const handleSave = async () => {
    if (hasError) return

    const formData = new FormData()
    formData.append('_csrf', csrfToken)
    formData.append('topics', topics)

    const response = await fetch(`${pageData.RepoLink}/topics`, {
      method: 'POST',
      body: formData
    })
    if (!response.ok) {
      const res = await response.json()
      res.invalidTopics.forEach((item, index) => {
        for (let i = 0; i < value.length; i++) {
          if (item === value[i].value) {
            value[i].error = true
            break
          }
        }
      })
      handleFail(res.message)
    } else {
      handleOk(topics)
    }
  }

  beforeUpdate(() => {
    // console.log(isCountInValid, isFormatInValid, filterText, value)
  })

  onMount(() => {
    handleToggler()
  })
</script>

<div class="form-control w-full">
  <Select {items} class={`!gap-1 !rounded-lg !border !border-base-content/10 ${hasError ? '!border-red-600 !bg-base-100' : '!bg-base-100'}`} multiple bind:filterText on:filter={handleFilter} on:change={handleChange} on:focus={handleFocus} bind:hasError bind:value>
    <div slot="selection" let:selection let:index>
      <div class="badge badge-ghost" class:!border-red-600={selection.error} data-value={selection.value}>
        {selection.value}
      </div>
    </div>
    <div slot="item" let:item title={item.label}>
      {item.value}
    </div>
    <div slot="input-hidden" let:value>
      <input type="hidden" name="topics" value={topics} required />
    </div>
  </Select>
  <div class="label pb-0 pt-1">
    <span class="label-text-alt" class:text-red-600={isCountInValid || isFormatInValid}>
      {#if isCountInValid}
        {pageData.RepoTopicCountPrompt}
      {/if}
      {#if isFormatInValid}
        {pageData.RepoTopicFormatPrompt}
      {/if}
    </span>
  </div>
</div>

<button type="button" class="btn-primary btn-sm btn w-fit" disabled={hasError} on:click={handleSave}>{pageData.RepoTopicDone} </button>

<WrappedSelectStyle />

<style lang="postcss">
  :global([svelte-repo-topics-manager] div.svelte-select) {
    @apply !px-2;
  }
</style>
