import $ from 'jquery'
import RepoTemplateSelect from '../components/RepoTemplateSelect.svelte'

const { appSubUrl } = window.config

export function initRepoTemplateSearch() {
  const uidInputSelector = 'input[name=uid]'
  const repoTemplateInputSelector = 'input[name=repo_template]'

  const $repoTemplate = $(repoTemplateInputSelector)
  const openTemplate = function (open = false) {
    const $templateUnits = $('[jq-template_units]')
    const $nonTemplate = $('[jq-non_template]')
    if (!open) {
      $templateUnits.removeClass('hidden')
      $nonTemplate.addClass('hidden')
    } else {
      $templateUnits.addClass('hidden')
      $nonTemplate.removeClass('hidden')
    }
  }

  const getOwnerUID = () => $(uidInputSelector).val()

  return {
    getOwnerUID,
    openTemplate
  }
}

export function initRepoTemplateSelect(selector) {
  const $els = $(selector)
  if (!$els.length) return

  const pageData = window.config.pageData.repoTemplate
  const el = $els[0]
  new RepoTemplateSelect({
    target: el,
    props: { pageData }
  })
}
