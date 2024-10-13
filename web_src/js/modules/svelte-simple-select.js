import $ from 'jquery'
import SimpleSvelteSelect from '../components/SimpleSvelteSelect.svelte'

export function initSimpleSvelteSelect(
  selector,
  dataKey,
  handleChangeCallback,
  handleClearCallback
) {
  const $els = $(selector)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData[dataKey]
  new SimpleSvelteSelect({
    target: el,
    props: { pageData, handleChangeCallback, handleClearCallback }
  })
}
