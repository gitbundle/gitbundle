<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Select from 'svelte-select'
  import WrappedSelectStyle from './WrappedSelectStyle.svelte'

  export let pageData

  let filterText = '',
    value,
    placeholder = ''

  if (pageData.IssueLabels !== null) {
    value.value = pageData.IssueLabels
    value.label = pageData.IssueLabels
  } else {
    placeholder = pageData.IssueLabelsHelper
  }

  const handleChange = ({ detail }) => {
    value = detail
  }

  beforeUpdate(() => {
    // console.log(value, items)
  })

  onMount(() => {
    // console.log(value, pageData)
  })
</script>

<Select items={pageData.items} class={pageData.ClassName ?? '!select-bordered !select w-full max-w-md'} clearable on:change={handleChange} bind:value bind:filterText {placeholder}>
  <div slot="selection" let:selection class="flex flex-col gap-1 whitespace-nowrap leading-none text-base-content">
    {selection.label !== null ? selection.label : ''}<i class="text-xs">{selection.desc !== null ? selection.desc : ''}</i>
  </div>
  <div slot="item" let:item class="flex flex-col gap-1 whitespace-pre-line leading-none text-base-content">
    {item.label}<i class="text-xs">{item.desc}</i>
  </div>
  <div slot="input-hidden" let:value>
    <input type="hidden" name="issue_labels" value={value?.value ?? ''} required />
  </div>
</Select>

<WrappedSelectStyle />
