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
    const response = await fetch(`${appSubUrl}/user/search?q=${filterText}`)
    if (!response.ok) return

    const res = await response.json()
    // console.log(res)
    items = res.data.map((item) => ({
      value: item.id,
      label: item.login,
      avatar_url: item.avatar_url
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
    <div class="flex items-center gap-2">
      <img class="h-5 w-5 rounded-full" src={selection.avatar_url} title={selection.label} alt={selection.label} />
      {selection.label !== null ? selection.label : ''}
    </div>
  </div>
  <div slot="item" let:item class="inline-block align-middle leading-none text-base-content">
    <div class="flex items-center gap-2">
      <img class="h-5 w-5 rounded-full" src={item.avatar_url} title={item.label} alt={item.label} />
      {item.label}
    </div>
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name={pageData.fieldName} value={value?.label ?? ''} autocomplete={pageData.autocomplete} required={pageData.required} />
  </div>
</Select>

<WrappedSelectStyle />
