<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'
  import { htmlEscape } from 'escape-goat'
  import { initRepoTemplateSearch } from '../modules/repo-template.js'

  export let pageData

  const { appSubUrl } = window.config
  const { openTemplate, getOwnerUID } = initRepoTemplateSearch()

  let filterText = '',
    value,
    placeholder = '',
    items

  if (pageData.RepoTemplate !== null) {
    value = {}
    value.value = pageData.RepoTemplate
    value.label = pageData.RepoTemplateName
  } else {
    placeholder = pageData.RepoTemplateName
  }

  const handleChange = ({ detail }) => {
    value = detail
    openTemplate(false)
  }

  const handleClear = ({ detail }) => {
    console.log(detail)
    openTemplate(true)
  }

  const getItems = async () => {
    const uid = getOwnerUID()
    if (!uid) {
      console.error('uid not initialized')
      return
    }
    const response = await fetch(`${appSubUrl}/repo/search?q=${filterText}&template=true&priority_owner_id=${uid}`)
    if (!response.ok) return
    const res = await response.json()
    // console.log(res)
    items = res.data.map((item) => ({
      value: item.id,
      label: htmlEscape(item.full_name)
    }))
  }

  beforeUpdate(() => {
    // console.log(value, items)
  })

  onMount(() => {
    // console.log(value, items)
    openTemplate(value === null || value.value === '' || value.value === '0')
  })
</script>

<Select {items} class={pageData.ClassName ?? '!select-bordered !select w-full max-w-md'} clearable on:change={handleChange} on:focus={getItems} on:clear={handleClear} bind:value bind:filterText {placeholder}>
  <div slot="selection" let:selection class="flex items-center gap-2 text-base-content">
    <span class="overflow-hidden text-ellipsis whitespace-nowrap">{selection.label}</span>
  </div>
  <div slot="item" let:item class="flex items-center gap-2 text-base-content">
    <span class="overflow-hidden text-ellipsis whitespace-nowrap">{item.label}</span>
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name="repo_template" value={value?.value ?? 0} required />
  </div>
</Select>

<WrappedSelectStyle />
