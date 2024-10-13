<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'

  export let pageData

  const { appSubUrl } = window.config

  let filterText = '',
    value,
    items = []

  const handleChange = ({ detail }) => {
    value = detail
  }

  const handleFocus = async () => {
    const response = await fetch(`${appSubUrl}/org/${pageData.OrgName}/teams/-/search?q=${filterText}`)
    if (!response.ok) return

    const res = await response.json()
    items = res.data.map((item) => ({
      value: item.id,
      label: item.name,
      description: item.description
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
    {selection.label !== null ? selection.label : ''}
  </div>
  <div slot="item" let:item class="inline-block align-middle leading-none text-base-content">
    {item.label}
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name={pageData.fieldName} value={value?.label ?? ''} autocomplete={pageData.autocomplete} required={pageData.required} />
  </div>
</Select>

<WrappedSelectStyle />
