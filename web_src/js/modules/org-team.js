import $ from 'jquery'
import SearchRepo from '../components/SearchRepo.svelte'

const { appSubUrl } = window.config

export function initOrgTeamSettings(selector) {
  // Change team access mode
  $(`${selector} input[name=permission]`).on('change', () => {
    const val = $('input[name=permission]:checked', selector).val()
    const $teamUnits = $(`${selector} [jq-team-units]`)
    if (val === 'admin') {
      $teamUnits.addClass('hidden')
    } else {
      $teamUnits.removeClass('hidden')
    }
  })
}

export function initSearchRepoBox(selector, pageDataKey) {
  const $els = $(selector)
  if (!$els.length) return
  const el = $els[0]

  const pageData = window.config.pageData[pageDataKey]
  new SearchRepo({
    target: el,
    props: { pageData }
  })
}

// export function initOrgTeamSearchRepoBox() {
//   const $searchRepoBox = $('#search-repo-box')
//   $searchRepoBox.search({
//     minCharacters: 2,
//     apiSettings: {
//       url: `${appSubUrl}/repo/search?q={query}&uid=${$searchRepoBox.data(
//         'uid'
//       )}`,
//       onResponse(response) {
//         const items = []
//         $.each(response.data, (_i, item) => {
//           items.push({
//             title: item.full_name.split('/')[1],
//             description: item.full_name
//           })
//         })

//         return { results: items }
//       }
//     },
//     searchFields: ['full_name'],
//     showNoResults: false
//   })
// }
