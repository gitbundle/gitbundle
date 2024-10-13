<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'

  export let pageData

  const { appSubUrl, csrfToken } = window.config

  let filterText = '',
    value,
    items = []

  const handleChange = ({ detail }) => {
    value = detail
  }

  const handleFocus = async () => {
    const response = await fetch(`${appSubUrl}/repo/search?q=${filterText}&uid=${pageData.uid}`)
    if (!response.ok) return

    const res = await response.json()
    // console.log(res)
    items = res.data.map((item) => ({
      value: item.id,
      label: item.full_name.split('/')[1],
      full_name: item.full_name
    }))
  }

  beforeUpdate(() => {
    // console.log(value)
  })

  onMount(() => {
    // console.log(pageData)
  })
</script>

<Select {items} class="!input-bordered !input" on:change={handleChange} on:focus={handleFocus} bind:value bind:filterText clearable placeholder={pageData.placeholder}>
  <div slot="selection" let:selection class="inline-block align-middle leading-none text-base-content">
    <div class="flex flex-col">
      <span>{selection.label}</span>
      <span class="text-xs text-gray-500">{selection.full_name}</span>
    </div>
  </div>
  <div slot="item" let:item class="inline-block align-middle leading-none text-base-content">
    <div class="flex flex-col">
      <span>{item.label}</span>
      <span class="text-xs text-gray-500">{item.full_name}</span>
    </div>
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name={pageData.fieldName} value={value?.label ?? ''} autocomplete={pageData.autocomplete} required={pageData.required} />
  </div>
</Select>

<WrappedSelectStyle />
