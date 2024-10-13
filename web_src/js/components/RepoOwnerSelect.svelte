<script>
  import { onMount, beforeUpdate, createEventDispatcher } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'

  export let pageData

  let value = {
    value: pageData.ContextUserValue,
    label: pageData.ContextUserLabel,
    Name: pageData.ContextUserName,
    Avatar: pageData.ContextUserAvatar
  }

  const handleChange = ({ detail }) => {
    value = detail
  }

  beforeUpdate(() => {
    // console.log(value)
  })

  onMount(() => {
    // console.log(pageData)
  })
</script>

<Select items={pageData.items} class={pageData.ClassName ?? '!select-bordered !select w-full max-w-md'} required={pageData.Required} clearable={false} on:change={handleChange} {value}>
  <div slot="selection" let:selection class="flex items-center gap-2 text-base-content">
    {@html selection.Avatar}
    <span class="overflow-hidden text-ellipsis whitespace-nowrap">{selection.label}</span>
  </div>
  <div slot="item" let:item title={item.Name} class="flex items-center gap-2 text-base-content">
    {@html item.Avatar}
    <span class="overflow-hidden text-ellipsis whitespace-nowrap">{item.label}</span>
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name="uid" value={value.value} required />
  </div>
</Select>

<WrappedSelectStyle />
