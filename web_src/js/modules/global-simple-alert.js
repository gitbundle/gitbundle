import GlobalSimpleAlert from '../components/GlobalSimpleAlert.svelte'

const selector = '[svelte-simple-alert]'
const els = document.querySelectorAll(selector)
els.forEach((el) => {
  let data = {}
  data.selector = selector
  data.title = el.getAttribute('data-title') ?? undefined
  data.desc = el.getAttribute('data-desc') ?? undefined
  data.type = el.getAttribute('data-type') ?? undefined
  data.icon = el.getAttribute('data-icon') ?? true
  data.iconName = el.getAttribute('data-icon-name') ?? undefined
  data.close = el.getAttribute('data-close') ?? undefined
  data.opacity = el.getAttribute('data-opacity') ?? true
  data.extraClass = el.getAttribute('data-extra-class') ?? undefined

  new GlobalSimpleAlert({
    target: el,
    props: { data }
  })
})
