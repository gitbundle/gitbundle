import TypeaheadDropdown from '../components/TypeaheadRepoSearchDropdown.svelte'

export function initTypeaheadRepoSearchDropdown(selector) {
  const els = document.querySelectorAll(selector)
  if (!els.length) return
  if (els.length > 1) {
    console.error(`found more than one container "${selector}"`)
    return
  }

  const container = els[0]
  const pageData = {
    selected: container.getAttribute('data-selected')
  }
  new TypeaheadDropdown({
    target: container,
    props: { pageData }
  })
}

// export function initRepoIssueReferenceRepositorySearch() {
//   $('.issue_reference_repository_search').dropdown({
//     apiSettings: {
//       url: `${appSubUrl}/repo/search?q={query}&limit=20`,
//       onResponse(response) {
//         const filteredResponse = { success: true, results: [] }
//         $.each(response.data, (_r, repo) => {
//           filteredResponse.results.push({
//             name: htmlEscape(repo.full_name),
//             value: repo.full_name
//           })
//         })
//         return filteredResponse
//       },
//       cache: false
//     },
//     onChange(_value, _text, $choice) {
//       const $form = $choice.closest('form')
//       $form.attr('action', `${appSubUrl}/${_text}/issues/new`)
//     },
//     fullTextSearch: true
//   })
// }
