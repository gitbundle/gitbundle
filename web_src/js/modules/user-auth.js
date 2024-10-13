import $ from 'jquery'

export function initUserAuthOauth2() {
  const $oauth2LoginNav = $('[jq-oauth2-login-navigator]')
  if ($oauth2LoginNav.length === 0) return

  $oauth2LoginNav.find('[jq-oauth-login-image]').click(() => {
    const oauthLoader = $('[jq-oauth2-login-loader]')
    const oauthNav = $('[jq-oauth2-login-navigator]')

    oauthNav.addClass('disabled')
    oauthLoader.removeClass('hidden')

    setTimeout(() => {
      // recover previous content to let user try again
      // usually redirection will be performed before this action
      oauthLoader.addClass('hidden')
      oauthNav.removeClass('disabled')
    }, 5000)
  })
}

export function initUserAuthLinkAccountView() {
  const $lnkUserPage = $('[jq-user-link-account]')
  if ($lnkUserPage.length === 0) {
    return false
  }

  const $tabs = $lnkUserPage.find('[jq-tabs]')
  const $contents = $lnkUserPage.find('[jq-contents]')

  const $signinTab = $tabs.find('[data-tab="auth-link-signin-tab"]')
  const $signUpTab = $tabs.find('[data-tab="auth-link-signup-tab"]')
  const $signInView = $contents.find('[data-tab="auth-link-signin-tab"]')
  const $signUpView = $contents.find('[data-tab="auth-link-signup-tab"]')

  $signUpTab.on('click', () => {
    $signinTab.removeClass('tab-active')
    $signInView.addClass('hidden')
    $signUpTab.addClass('tab-active')
    $signUpView.removeClass('hidden')
    return false
  })

  $signinTab.on('click', () => {
    $signUpTab.removeClass('tab-active')
    $signUpView.addClass('hidden')
    $signinTab.addClass('tab-active')
    $signInView.removeClass('hidden')
  })
}
