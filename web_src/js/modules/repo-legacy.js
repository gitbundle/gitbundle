// import { initRepoBranchTagDropdown } from '../components/RepoBranchTagDropdown.js'
import { createCommentEasyMDE, getAttachedEasyMDE } from '../common/EasyMDE.js'
import { initEasyMDEImagePaste } from '../common/ImagePaste.js'
import { initCompLabelEdit } from '../common/LabelEdit.js'
import { initCompMarkupContentPreviewTab } from '../common/MarkupContentPreview.js'
import { createDropzone } from '../common/dropzone.js'
import $ from '../features/jquery.js'
import { initCommentContent, initMarkupContent } from '../markup/content.js'
import { initReactionSelector } from './reaction-selector.js'
import {
  initRepoCloneLink,
  initRepoCommonBranchOrTagDropdown,
  initRepoCommonFilterSearchDropdown,
  initRepoCommonLanguageStats
} from './repo-common.js'
// import { initRepoDiffConversationNav } from './repo-diff.js'
// import initRepoPullRequestMergeForm from './repo-issue-pr-form.js'
import RepoOwnerSelect from '../components/RepoOwnerSelect.svelte'
import {
  initRepoIssueBranchSelect,
  // initRepoIssueCodeCommentCancel,
  initRepoIssueCommentDelete,
  // initRepoIssueComments,
  // initRepoIssueDependencyDelete,
  initRepoIssueReferenceIssue,
  // initRepoIssueStatusButton,
  initRepoIssueTitleEdit,
  initRepoIssueWipToggle,
  initRepoPullRequestUpdate
} from './repo-issue.js'
import { initRepoSettingBranches } from './repo-settings.js'
import { initUnicodeEscapeButton } from './repo-unicode-escape.js'
import attachTribute from './tribute.js'

const { csrfToken } = window.config

export async function onEditContent(event) {
  event.preventDefault()

  // $(this).closest('.dropdown').find('.menu').toggle('visible')
  const $segment = $(this).closest('[jq-content]')
  const $editContentZone = $segment.find('.edit-content-zone')
  const $renderContent = $segment.find('.render-content')
  const $rawContent = $segment.find('.raw-content')
  let $textarea
  let easyMDE

  // Setup new form
  if ($editContentZone.html().length === 0) {
    $editContentZone.html($('#edit-content-form').html())
    $textarea = $editContentZone.find('textarea')
    await attachTribute($textarea.get(), { mentions: true, emoji: true })

    let dz
    const $dropzone = $editContentZone.find('.dropzone')
    if ($dropzone.length === 1) {
      $dropzone.data('saved', false)

      const fileUuidDict = {}
      dz = await createDropzone($dropzone[0], {
        url: $dropzone.data('upload-url'),
        headers: { 'X-Csrf-Token': csrfToken },
        maxFiles: $dropzone.data('max-file'),
        maxFilesize: $dropzone.data('max-size'),
        acceptedFiles: ['*/*', ''].includes($dropzone.data('accepts'))
          ? null
          : $dropzone.data('accepts'),
        addRemoveLinks: true,
        dictDefaultMessage: $dropzone.data('default-message'),
        dictInvalidFileType: $dropzone.data('invalid-input-type'),
        dictFileTooBig: $dropzone.data('file-too-big'),
        dictRemoveFile: $dropzone.data('remove-file'),
        timeout: 0,
        thumbnailMethod: 'contain',
        thumbnailWidth: 480,
        thumbnailHeight: 480,
        init() {
          this.on('success', (file, data) => {
            file.uuid = data.uuid
            fileUuidDict[file.uuid] = { submitted: false }
            const input = $(
              `<input id="${data.uuid}" name="files" type="hidden">`
            ).val(data.uuid)
            $dropzone.find('.files').append(input)
          })
          this.on('removedfile', (file) => {
            $(`#${file.uuid}`).remove()
            if (
              $dropzone.data('remove-url') &&
              !fileUuidDict[file.uuid].submitted
            ) {
              $.post($dropzone.data('remove-url'), {
                file: file.uuid,
                _csrf: csrfToken
              })
            }
          })
          this.on('submit', () => {
            $.each(fileUuidDict, (fileUuid) => {
              fileUuidDict[fileUuid].submitted = true
            })
          })
          this.on('reload', () => {
            $.getJSON($editContentZone.data('attachment-url'), (data) => {
              dz.removeAllFiles(true)
              $dropzone.find('.files').empty()
              $.each(data, function () {
                const imgSrc = `${$dropzone.data('link-url')}/${this.uuid}`
                dz.emit('addedfile', this)
                dz.emit('thumbnail', this, imgSrc)
                dz.emit('complete', this)
                dz.files.push(this)
                fileUuidDict[this.uuid] = { submitted: true }
                $dropzone.find(`img[src='${imgSrc}']`).css('max-width', '100%')
                const input = $(
                  `<input id="${this.uuid}" name="files" type="hidden">`
                ).val(this.uuid)
                $dropzone.find('.files').append(input)
              })
            })
          })
        }
      })
      dz.emit('reload')
    }
    // Give new write/preview data-tab name to distinguish from others
    const $editContentForm = $editContentZone.find('[jq-edit-comment-form]')
    // TODO: we do not have the following class or attribute @low
    // const $tabMenu = $editContentForm.find('[jq-tabs]')
    // $tabMenu.attr('data-write', $editContentZone.data('write'))
    // $tabMenu.attr('data-preview', $editContentZone.data('preview'))
    // $tabMenu
    //   .find('.write.item')
    //   .attr('data-tab', $editContentZone.data('write'))
    // $tabMenu
    //   .find('.preview.item')
    //   .attr('data-tab', $editContentZone.data('preview'))
    // $editContentForm
    //   .find('.write')
    //   .attr('data-tab', $editContentZone.data('write'))
    // $editContentForm
    //   .find('.preview')
    //   .attr('data-tab', $editContentZone.data('preview'))

    // easyMDE = getAttachedEasyMDE($textarea)//TODO: why can't get initialized easyMDE
    easyMDE = await createCommentEasyMDE($textarea)

    initCompMarkupContentPreviewTab($editContentForm)
    if ($dropzone.length === 1) {
      initEasyMDEImagePaste(easyMDE, $dropzone[0], $dropzone.find('.files'))
    }

    // .cancel.button
    const $saveButton = $editContentZone.find('[jq-save]')
    $textarea.off('ce-quick-submit').on('ce-quick-submit', () => {
      $saveButton.trigger('click')
    })

    // .cancel.button
    $editContentZone
      .find('[jq-cancel]')
      .off('click')
      .on('click', () => {
        $renderContent.show()
        $editContentZone.hide()
        if (dz) {
          dz.emit('reload')
        }
      })

    $saveButton.off('click').on('click', () => {
      $renderContent.show()
      $editContentZone.hide()
      const $attachments = $dropzone
        .find('.files')
        .find('[name=files]')
        .map(function () {
          return $(this).val()
        })
        .get()
      $.post(
        $editContentZone.data('update-url'),
        {
          _csrf: csrfToken,
          content: $textarea.val(),
          context: $editContentZone.data('context'),
          files: $attachments
        },
        (data) => {
          if (data.length === 0 || data.content.length === 0) {
            $renderContent.html($('#no-content').html())
            $rawContent.text('')
          } else {
            $renderContent.html(data.content)
            $rawContent.text($textarea.val())
          }
          const $content = $renderContent.parent()
          const $tmpAttachments = $content.find('[jq-dropzone-attachments]')
          if (!$tmpAttachments.length) {
            if (data.attachments !== '') {
              $content.append(`<div jq-dropzone-attachments></div>`)
              $content
                .find('[jq-dropzone-attachments]')
                .replaceWith(data.attachments)
            }
          } else if (data.attachments === '') {
            $content.find('[jq-dropzone-attachments]').remove()
          } else {
            $content
              .find('[jq-dropzone-attachments]')
              .replaceWith(data.attachments)
          }
          if (dz) {
            dz.emit('submit')
            dz.emit('reload')
          }
          initMarkupContent()
          initCommentContent()
        }
      )
    })
  } else {
    // use existing form
    $textarea = $segment.find('textarea')
    easyMDE = getAttachedEasyMDE($textarea)
  }

  // Show write/preview tab and copy raw content as needed
  $editContentZone.show()
  $renderContent.hide()
  if ($textarea.val().length === 0) {
    $textarea.val($rawContent.text())
    easyMDE.value($rawContent.text())
  }
  requestAnimationFrame(() => {
    $textarea.focus()
    easyMDE.codemirror.focus()
  })
}

export function initRepository() {
  if ($('.repository').length === 0) {
    return
  }

  // File list and commits
  // TODO: we use svelte component for this, should double check @high
  // if (
  //   $('.repository.file.list').length > 0 ||
  //   $('.branch-dropdown').length > 0 ||
  //   $('.repository.commits').length > 0 ||
  //   $('.repository.release').length > 0
  // ) {
  //   initRepoBranchTagDropdown('.choose.reference .dropdown')
  // }

  // Wiki
  // if ($('.repository.wiki.view').length > 0) {
  //   initRepoCommonFilterSearchDropdown('.choose.page .dropdown')
  // }

  // Options

  // Labels
  initCompLabelEdit('.repository.labels')

  // Repo Creation
  if ($('.repository.new.repo').length > 0) {
    $('input[name="gitignores"], input[name="license"]').on('change', () => {
      const gitignores = $('input[name="gitignores"]').val()
      const license = $('input[name="license"]').val()
      if (gitignores || license) {
        $('input[name="auto_init"]').prop('checked', true)
      }
    })
  }

  // Compare or pull request
  const $repoDiff = $('.repository.diff')
  if ($repoDiff.length) {
    initRepoCommonBranchOrTagDropdown('.choose.branch .dropdown')
    initRepoCommonFilterSearchDropdown('.choose.branch .dropdown')
  }

  initRepoCloneLink()
  initRepoCommonLanguageStats()
  initRepoSettingBranches()

  // Issues
  if ($('.repository.view.issue').length > 0) {
    initRepoIssueCommentEdit()

    initRepoIssueBranchSelect()
    initRepoIssueTitleEdit()
    initRepoIssueWipToggle()
    initRepoIssueComments()

    initRepoDiffConversationNav()
    initRepoIssueReferenceIssue()

    initRepoIssueCommentDelete()
    initRepoIssueDependencyDelete()
    initRepoIssueCodeCommentCancel()
    initRepoIssueStatusButton()
    initRepoPullRequestUpdate()
    initReactionSelector()

    initRepoPullRequestMergeForm()
  }

  // Pull request
  const $repoComparePull = $('.repository.compare.pull')
  if ($repoComparePull.length > 0) {
    // show pull request form
    $repoComparePull.find('button.show-form').on('click', function (e) {
      e.preventDefault()
      $(this).parent().hide()

      const $form = $repoComparePull.find('.pullrequest-form')
      const easyMDE = getAttachedEasyMDE($form.find('textarea.edit_area'))
      $form.show()
      easyMDE.codemirror.refresh()
    })
  }

  initUnicodeEscapeButton()
}

export function initRepoIssueCommentEdit(selector) {
  // Issue/PR Context Menus
  // TODO: we do not have dropdown plugin, need double check @low
  // $('.comment-header-right .context-dropdown').dropdown({ action: 'hide' })

  // Edit issue or comment content
  // $(document).on('click', '.edit-content', onEditContent)
  $(document)
    .off('click', `${selector} [jq-edit-content]`)
    .on('click', `${selector} [jq-edit-content]`, onEditContent)

  // Quote reply
  $(document)
    .off('click', `${selector} [jq-quote-reply]`)
    .on('click', `${selector} [jq-quote-reply]`, function (event) {
      // $(this).closest('.dropdown').find('.menu').toggle('visible')
      const target = $(this).data('target')
      const quote = $(`#comment-${target}`).text().replace(/\n/g, '\n> ')
      const content = `> ${quote}\n\n`
      let easyMDE
      if ($(this).is('[jq-quote-reply-diff]')) {
        const $parent = $(this).closest('[jq-comment-code-cloud]')
        $parent.find('button[jq-comment-form-reply]').trigger('click')
        easyMDE = getAttachedEasyMDE($parent.find('[name="content"]'))
      } else {
        // for normal issue/comment page
        easyMDE = getAttachedEasyMDE($('[jq-issue-comment-form] .edit_area'))
      }
      if (easyMDE) {
        if (easyMDE.value() !== '') {
          easyMDE.value(`${easyMDE.value()}\n\n${content}`)
        } else {
          easyMDE.value(`${content}`)
        }
        requestAnimationFrame(() => {
          easyMDE.codemirror.focus()
          easyMDE.codemirror.setCursor(easyMDE.codemirror.lineCount(), 0)
        })
      }
      event.preventDefault()
    })
}

export function initNewRepo(selector) {
  if (!$(selector).length) return

  const $els = $(`${selector} [svelte-repo-owner-select]`)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData.repoOwnerSelect
  new RepoOwnerSelect({
    target: el,
    props: { pageData }
  })
}
