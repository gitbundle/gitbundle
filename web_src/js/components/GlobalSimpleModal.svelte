<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Svg from './Svg.svelte'
  import jquery from 'jquery'

  export let data
  export let formData

  /*
  let data = {
      csrfToken: window.config.csrfToken,
      asyncForm: true,
      checked: false,
      ...window.config.pageData.modalDeletion
    },
    formData = null
  const dyFieldType = {
      _csrf: 'hidden'
    },
    dyFieldLabel = {},
    dyFieldRequired = {}

  const els = document.querySelectorAll('[svelte-simple-modal]')
  els.forEach((el) => {
    el.addEventListener('click', () => {
      formData = new FormData()
      formData.append('_csrf', window.config.csrfToken)

      Array.from(el.attributes).forEach(({ name, value }) => {
        // console.log(name, value)
        const match = name.match(/data-form-key\[([0-9-a-zA-Z]+)\]/)
        if (match && match.length == 2) {
          // console.log(match, match[1])
          const fieldValue = el.getAttribute(`data-form-val[${match[1]}]`)
          if (!fieldValue) {
            console.log(
              `Unmatched value for '${match}' '${value}' '${fieldValue}'`
            )
          }
          formData.append(value, fieldValue ?? '')

          const fieldType = el.getAttribute(`data-form-type[${match[1]}]`)
          dyFieldType[value] = fieldType ?? 'hidden'

          const fieldLabel = el.getAttribute(`data-form-label[${match[1]}]`)
          dyFieldLabel[value] = fieldLabel ?? undefined

          const fieldRequired = el.getAttribute(
            `data-form-required[${match[1]}]`
          )
          dyFieldRequired[value] = fieldRequired ?? undefined
        }
      })
      data.method = el.getAttribute('data-method') ?? 'post'
      data.href = el.getAttribute('data-href')
      data.name = el.getAttribute('data-name')
      data.desc = el.getAttribute('data-desc') ?? undefined
      data.descClass = el.getAttribute('data-desc-class') ?? undefined
      data.title = el.getAttribute('data-title')
      data.titleSvgName = el.getAttribute('data-title-svg-name') ?? undefined
      data.approveText = el.getAttribute('data-action-approve-text')
      data.approveColor = el.getAttribute('data-action-approve-color')
      data.negativeText = el.getAttribute('data-action-negative-text')
      data.confirmForm = el.getAttribute('data-confirm-form') ?? undefined

      const asyncForm = el.getAttribute('data-async')
      if (asyncForm !== undefined) {
        data.async = asyncForm
      }

      if (el.tagName !== 'label') {
        data.checked = true
      }
    })
  })
  */

  const simpleModalId = `svelte-simple-modal` + Math.random()

  const handleSubmit = async () => {
    if (data.confirmForm) {
      jquery(data.confirmForm).submit()
      return
    }
    // formData.append('_csrf', window.config.csrfToken)
    // for (const pair of formData.entries()) {
    //   console.log(pair)
    // }
    // return
    const response = await fetch(data.href, {
      method: data.method,
      body: formData
    })
    if (response.status === 301 || response.status === 200) {
      handleDismiss(null, true)
      try {
        const body = await response.json()
        if (body.redirect) {
          window.location.href = body.redirect
        } else {
          window.location.reload()
        }
      } catch (error) {
        console.log(error)
        window.location.reload()
      }
    }
  }

  const handleDismiss = (e, clearHistory = false) => {
    data.checked = false
    if (clearHistory) {
      window.history.pushState(null, null, ' ')
    }
  }

  beforeUpdate(() => {
    // console.log(data)
  })

  onMount(() => {})
</script>

<input type="checkbox" id={simpleModalId} class="modal-toggle" checked={data.checked} />
<div for={simpleModalId} class="modal">
  <div class="modal-box relative overflow-y-visible">
    <h3 class="flex flex-wrap items-center gap-2 text-lg font-bold">
      {#if data.titleSvgName === undefined}
        <Svg name="octicon-trash" size={16} />
      {:else if data.titleSvgName !== ''}
        <Svg name={data.titleSvgName} size={16} />
      {/if}
      {@html data.title ?? data.defaultTextModalTitle}
      {#if data.name}
        <span>{data.name}</span>
      {/if}
    </h3>
    {#if data.desc !== undefined}
      <div class:py-4={data.desc !== ''} class={data.descClass}>
        {@html data.desc ?? data.defaultTextModalDesc}
      </div>
    {/if}
    {#if data.async === 'false'}
      <form action={data.href} method={data.method} class="space-y-4">
        <!-- <input type="hidden" name="_csrf" value={data.csrfToken} /> -->
        {#each Array.from(formData.entries()) as [name, value]}
          {#if data.dyFieldLabel[name]}
            <div class="form-control !mt-0 w-full" class:required={data.dyFieldRequired[name]}>
              <div class="label">
                <span class="label-text">{@html data.dyFieldLabel[name]}</span>
              </div>
              <input type={data.dyFieldType[name]} {name} {value} class="input-bordered input w-full" required={data.dyFieldRequired[name] !== undefined} />
            </div>
          {:else}
            <input type={data.dyFieldType[name]} {name} {value} />
          {/if}
        {/each}

        <div class="modal-action">
          <div class="grid max-w-xs grid-cols-2 gap-2">
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <span class="btn" type="reset" on:click={handleDismiss}>
              {data.negativeText ?? data.textModalNo}
            </span>
            <button class={`btn-${data.approveColor ?? 'primary'} btn`} type="submit" on:click={(e) => handleDismiss(e, true)}>
              {data.approveText ?? data.textModalYes}
            </button>
          </div>
        </div>
      </form>
    {:else}
      <div class="modal-action">
        <div class="grid max-w-xs grid-cols-2 gap-2">
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <span class="btn" on:click={handleDismiss}>
            {data.negativeText ?? data.textModalNo}
          </span>
          <button on:click={handleSubmit} class={`btn-${data.approveColor ?? 'primary'} btn`}>
            {data.approveText ?? data.textModalYes}
          </button>
        </div>
      </div>
    {/if}
  </div>
</div>
