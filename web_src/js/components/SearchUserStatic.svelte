<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'

  export let pageData

  let filterText = '',
    value,
    placeholder = ''

  if (pageData.defaultValue !== null && pageData.defaultValue !== '') {
    value = pageData.multiple ? [] : {}
    const values = pageData.defaultValue.split(',').map((item) => parseInt(item))
    pageData.items.forEach((item) => {
      if (values.includes(item.value)) {
        if (pageData.multiple) {
          value.push(item)
        } else {
          value = item
        }
      }
    })
  } else {
    placeholder = pageData.placeholder
  }

  const handleChange = ({ detail }) => {
    value = detail
  }

  const handleClear = ({ detail }) => {
    // console.log(detail, value)
  }

  beforeUpdate(() => {
    // console.log(value, items)
  })

  onMount(() => {
    // console.log(value, pageData)
  })
</script>

<Select items={pageData.items} class="!input-bordered !input" on:change={handleChange} on:clear={handleClear} bind:value bind:filterText {placeholder} clearable multiple={pageData.multiple ? true : false}>
  <div slot="selection" let:selection class="inline-block align-middle leading-none text-base-content">
    <div class="flex items-center gap-2">
      {#if selection.avatar_img}
        <div class="avatar">
          <span class="h-5 w-5">
            {@html selection.avatar}
          </span>
        </div>
      {:else}
        {@html selection.avatar}
      {/if}
      {selection.label}
    </div>
  </div>
  <div slot="item" let:item class="inline-block align-middle leading-none text-base-content">
    <div class="flex items-center gap-2">
      {#if item.avatar_img}
        <div class="avatar">
          <span class="h-7 w-7">
            {@html item.avatar}
          </span>
        </div>
      {:else}
        {@html item.avatar}
      {/if}
      {item.label}
    </div>
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name={pageData.fieldName} value={pageData.multiple ? value?.map((item) => item.value).join(',') : value?.value ?? 0} required />
  </div>
</Select>

<WrappedSelectStyle />
