import $ from 'jquery'

import { svg } from './common/svg.js'
import { strSubMatch } from './utils/index.js'
const { csrf } = window.config

const threshold = 50
let files = []
let $repoFindFileInput, $repoFindFileTableBody, $repoFindFileNoResult

function filterRepoFiles(filter) {
  const treeLink = $repoFindFileInput.attr('data-url-tree-link')
  $repoFindFileTableBody.empty()

  const fileRes = []
  if (filter) {
    for (let i = 0; i < files.length && fileRes.length < threshold; i++) {
      const subMatch = strSubMatch(files[i], filter)
      if (subMatch.length > 1) {
        fileRes.push(subMatch)
      }
    }
  } else {
    for (let i = 0; i < files.length && i < threshold; i++) {
      fileRes.push([files[i]])
    }
  }

  const tmplRow = `<div class="px-2 py-3 hover:bg-base-content/10 hover:first:rounded-t-lg hover:last:rounded-b-lg line-clamp-1"><a class="link link-hover text-base-content"></a></div>`

  $repoFindFileNoResult.toggle(fileRes.length === 0)
  for (const matchRes of fileRes) {
    const $row = $(tmplRow)
    const $a = $row.find('a')
    $a.attr('href', `${treeLink}/${matchRes.join('')}`)
    const $octiconFile = svg('octicon-file')
    $a.append($octiconFile)
    // if the target file path is "abc/xyz", to search "bx", then the matchRes is ['a', 'b', 'c/', 'x', 'yz']
    // the matchRes[odd] is matched and highlighted to red.
    for (let j = 0; j < matchRes.length; j++) {
      if (!matchRes[j]) continue
      const $span = $('<span>').text(matchRes[j])
      if (j % 2 === 1) $span.addClass('text-warning')
      $a.append($span)
    }
    $repoFindFileTableBody.append($row)
  }
}

async function loadRepoFiles() {
  files = await $.ajax({
    url: $repoFindFileInput.attr('data-url-data-link'),
    headers: { 'X-Csrf-Token': csrf }
  })
  filterRepoFiles($repoFindFileInput.val())
}

function _init() {
  $repoFindFileInput = $('#repo-file-find-input')
  if (!$repoFindFileInput.length) return

  $repoFindFileTableBody = $('#repo-find-file-lists')
  $repoFindFileNoResult = $('#repo-find-file-no-result')
  $repoFindFileInput.on('input', () => filterRepoFiles($repoFindFileInput.val()))

  loadRepoFiles()
}

_init()
