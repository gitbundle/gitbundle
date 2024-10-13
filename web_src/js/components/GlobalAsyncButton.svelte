<script>
  const formData = new FormData()
  formData.append('_csrf', window.config.csrfToken)

  const els = document.querySelectorAll('[svelte-async-button]')
  els.forEach((el) => {
    el.addEventListener('click', async () => {
      Array.from(el.attributes).forEach(({ name, value }) => {
        // console.log(name, value)
        const match = name.match(/data-form-key\[([0-9-a-zA-Z]+)\]/)
        if (match && match.length == 2) {
          // console.log(match, match[1])
          const fieldValue = el.getAttribute(`data-form-val[${match[1]}]`)
          if (!fieldValue) {
            // TODO: what we should do about not matched form value @low
            console.error(`Unmatched value for '${match}'`)
            return
          }
          formData.append(value, fieldValue)
        }
      })

      const redirect = el.getAttribute('data-redirect') ?? undefined

      const href = el.getAttribute('data-href')
      const response = await fetch(href, {
        method: 'POST',
        body: formData
      })
      // console.log(response)
      if (response.status === 301) {
        const body = await response.json()
        // console.log(body)
        window.location.href = body.redirect
      } else {
        if (response.status === 200 && redirect) {
          window.location.href = redirect
        } else {
          window.location.reload()
        }
      }
    })
  })
</script>
