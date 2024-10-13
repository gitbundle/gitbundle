import { htmlEscape } from 'escape-goat'
import $ from '../features/jquery.js'
import { initMarkupContent } from '../markup/content.js'
import { createCodeEditor } from './codeeditor.js'

const { csrfToken } = window.config
let previewFileModes

function initEditPreviewTab($form) {
  const $tabMenu = $form.find('[jq-tabs]')
  const $contents = $form.find('[jq-contents]')
  const $previewTab = $tabMenu.find(`[data-tab="${$tabMenu.data('preview')}"]`)
  if ($previewTab.length) {
    previewFileModes = $previewTab.data('preview-file-modes').split(',')
    $previewTab.off('click').on('click', function () {
      const $this = $(this)
      $this.addClass('tab-active').siblings().removeClass('tab-active')

      let context = `${$this.data('context')}/`
      const mode = $this.data('markdown-mode') || 'comment'
      const treePathEl = $form.find('input#tree_path')
      if (treePathEl.length > 0) {
        context += treePathEl.val()
      }
      context = context.substring(0, context.lastIndexOf('/'))
      $.post(
        $this.data('url'),
        {
          _csrf: csrfToken,
          mode,
          context,
          text: $contents
            .find(`[data-tab="${$tabMenu.data('write')}"] textarea`)
            .val()
        },
        (data) => {
          const $previewPanel = $contents.find(
            `[data-tab="${$tabMenu.data('preview')}"]`
          )
          $previewPanel
            .removeClass('hidden')
            .html(data)
            .siblings()
            .addClass('hidden')
          initMarkupContent()
        }
      )
    })
  }
}

function initEditDiffTab($form) {
  const $tabMenu = $form.find('[jq-tabs]')
  const $contents = $form.find('[jq-contents]')
  $tabMenu
    .find(`[data-tab="${$tabMenu.data('diff')}"]`)
    .off('click')
    .on('click', function () {
      const $this = $(this)
      $this.addClass('tab-active').siblings().removeClass('tab-active')

      $.post(
        $this.data('url'),
        {
          _csrf: csrfToken,
          context: $this.data('context'),
          content: $contents
            .find(`[data-tab="${$tabMenu.data('write')}"] textarea`)
            .val()
        },
        (data) => {
          const $diffPreviewPanel = $contents.find(
            `[data-tab="${$tabMenu.data('diff')}"]`
          )
          $diffPreviewPanel
            .removeClass('hidden')
            .html(data)
            .siblings()
            .addClass('hidden')
        }
      )
    })
}

function initEditWriteTab($form) {
  const $tabMenu = $form.find('[jq-tabs]')
  const $contents = $form.find('[jq-contents]')
  $tabMenu
    .find(`[data-tab="${$tabMenu.data('write')}"]`)
    .off('click')
    .on('click', function () {
      $(this).addClass('tab-active').siblings().removeClass('tab-active')
      $contents
        .find(`[data-tab="${$tabMenu.data('write')}"]`)
        .removeClass('hidden')
        .siblings()
        .addClass('hidden')
    })
}

function initEditorForm() {
  const $form = $('[jq-repository-edit-form]')
  if ($form.length === 0) {
    return
  }

  initEditWriteTab($form)
  initEditPreviewTab($form)
  initEditDiffTab($form)
}

function getCursorPosition($e) {
  const el = $e.get(0)
  let pos = 0
  if ('selectionStart' in el) {
    pos = el.selectionStart
  } else if ('selection' in document) {
    el.focus()
    const Sel = document.selection.createRange()
    const SelLength = document.selection.createRange().text.length
    Sel.moveStart('character', -el.value.length)
    pos = Sel.text.length - SelLength
  }
  return pos
}

export function initRepoEditor(selector) {
  if (!$(selector).length) return

  initEditorForm()

  const $editFilename = $('#file-name')
  $editFilename
    .on('keyup', function (e) {
      const $section = $('[jq-breadcrumb] span[jq-section]')
      const $divider = $('[jq-breadcrumb] span[jq-divider]')
      let value
      let parts

      if (
        e.keyCode === 8 &&
        getCursorPosition($(this)) === 0 &&
        $section.length > 0
      ) {
        value = $section.last().find('a').text()
        $(this).val(value + $(this).val())
        $(this)[0].setSelectionRange(value.length, value.length)
        $section.last().remove()
        $divider.last().remove()
      }
      if (e.keyCode === 191) {
        parts = $(this).val().split('/')
        for (let i = 0; i < parts.length; ++i) {
          value = parts[i]
          if (i < parts.length - 1) {
            if (value.length) {
              $(
                `<span jq-section><a class="link-hover link-primary link" href="#">${htmlEscape(
                  value
                )}</a></span>`
              ).insertBefore($(this))
              $('<span jq-divider>/</span>').insertBefore($(this))
            }
          } else {
            $(this).val(value)
          }
          $(this)[0].setSelectionRange(0, 0)
        }
      }
      parts = []
      $('[jq-breadcrumb] span[jq-section]').each(function () {
        const element = $(this)
        if (element.find('a').length) {
          parts.push(element.find('a').text())
        } else {
          parts.push(element.text())
        }
      })
      if ($(this).val()) parts.push($(this).val())
      $('#tree_path').val(parts.join('/'))
    })
    .trigger('keyup')

  const $editArea = $('textarea[jq-edit_area]')
  if (!$editArea.length) return
  ;(async () => {
    const editor = await createCodeEditor(
      $editArea[0],
      $editFilename[0],
      previewFileModes
    )

    // Using events from https://github.com/codedance/jquery.AreYouSure#advanced-usage
    // to enable or disable the commit button
    const $commitButton = $('[jq-commit-button]')
    const $editForm = $('[jq-repository-edit-form]')
    const dirtyFileClass = 'dirty-file'

    // Disabling the button at the start
    if ($('input[name="page_has_posted"]').val() !== 'true') {
      $commitButton.prop('disabled', true)
    }

    // Registering a custom listener for the file path and the file content
    $editForm.areYouSure({
      silent: true,
      dirtyClass: dirtyFileClass,
      fieldSelector: ':input:not([jq-commit-form-wrapper] :input)',
      change() {
        const dirty = $(this).hasClass(dirtyFileClass)
        $commitButton.prop('disabled', !dirty)
      }
    })

    // Update the editor from query params, if available,
    // only after the dirtyFileClass initialization
    const params = new URLSearchParams(window.location.search)
    const value = params.get('value')
    if (value) {
      editor.setValue(value)
    }

    $commitButton.off('click').on('click', (event) => {
      // A modal which asks if an empty file should be committed
      if ($editArea.val().length === 0) {
        $('[jq-edit-empty-content-modal] :checkbox').prop('checked', true)
        $('[jq-edit-empty-content-modal] button[type="button"]').prop(
          'disabled',
          false
        )
        $('[jq-edit-empty-content-modal] [jq-save]')
          .off('click')
          .on('click', function () {
            $(this)
              .prop('disabled', true)
              .closest('[jq-edit-empty-content-modal]')
              .find(':checkbox')
              .prop('checked', false)
            $editForm.trigger('submit')
          })
        $('[jq-edit-empty-content-modal] [jq-cancel]')
          .off('click')
          .on('click', function () {
            $(this)
              .prop('disabled', true)
              .closest('[jq-edit-empty-content-modal]')
              .find(':checkbox')
              .prop('checked', false)
          })
        event.preventDefault()
      }
    })
  })()
}
