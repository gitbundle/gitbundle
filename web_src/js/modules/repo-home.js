import $ from '../features/jquery.js'

const { appSubUrl, csrfToken } = window.config

export function initRepoTopicBar() {
  const viewDiv = $('[jq-repo-topics]')

  const handleToggler = () => {
    $('[jq-repo-topics] [jq-toggle]')
      .off('click')
      .on('click', function () {
        $(this)
          .parent()
          .siblings('[svelte-repo-topics-manager]')
          .toggleClass('hidden')
        $(this).parent().toggleClass('hidden')
      })
  }

  const handleFail = (message) => {
    $('div[jq-validate_prompt]').text(message).removeClass('hidden')
  }

  const handleOk = (topics) => {
    if (topics.length) {
      viewDiv.children('a').remove()
      const topicArray = topics.split(',')
      for (let i = 0; i < topicArray.length; i++) {
        const link = $('<a class="badge-ghost badge badge-md"></a>')
        link.attr(
          'href',
          `${appSubUrl}/explore/repos?q=${encodeURIComponent(
            topicArray[i]
          )}&topic=1`
        )
        link.text(topicArray[i])
        viewDiv.prepend(link)
      }
    }
    viewDiv
      .toggleClass('hidden')
      .siblings('[svelte-repo-topics-manager]')
      .toggleClass('hidden')
  }

  return {
    handleFail,
    handleOk,
    handleToggler
  }
}
