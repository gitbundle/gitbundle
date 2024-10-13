<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'

  export let pageData
  export let handleChangeCallback
  export let handleClearCallback
  export let value

  let filterText = '',
    placeholder = ''

  if (pageData.defaultValue !== null && pageData.defaultValue !== '') {
    value = pageData.defaultValue
  } else {
    placeholder = pageData.placeholder
  }

  const handleChange = ({ detail }) => {
    value = detail
    // console.log(value, Array.isArray(value))
    handleChangeCallback && handleChangeCallback(detail)
  }

  const handleClear = ({ detail }) => {
    // console.log(detail, value)
    handleClearCallback && handleClearCallback(detail)
  }

  beforeUpdate(() => {
    // console.log(value, pageData.multiple)
  })

  onMount(() => {
    // console.log(value, pageData)
  })
</script>

<Select items={pageData.items} class={pageData.ClassName ?? '!select-bordered !select w-full max-w-md'} on:change={handleChange} on:clear={handleClear} bind:value bind:filterText {placeholder} clearable={pageData.clearable} multiple={pageData.multiple}>
  <div slot="selection" let:selection class="inline-block align-middle leading-none text-base-content">
    {selection.label !== null ? selection.label : ''}
  </div>
  <div slot="item" let:item class="inline-block align-middle leading-none text-base-content">
    {item.label}
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name={pageData.fieldName} value={pageData.multiple && Array.isArray(value) ? value.map((item) => item.value).join(',') : typeof value === 'object' && value?.value ? value?.value : value} />
  </div>
</Select>

<WrappedSelectStyle />
