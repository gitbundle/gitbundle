#!/usr/bin/env bash
# This is an update script for gitbundle installed via the binary distribution
# from dl.gitbundle.com on linux as systemd service. It performs a backup and updates
# GitBundle in place.
# NOTE: This adds the GPG Signing Key of the GitBundle maintainers to the keyring.
# Depends on: bash, curl, xz, sha256sum. optionally jq, gpg
#   See section below for available environment vars.
#   When no version is specified, updates to the latest release.
# Examples:
#   upgrade.sh 1.15.10
#   gitbundlehome=/opt/gitbundle gitbundleconf=$gitbundlehome/app.ini upgrade.sh

# apply variables from environment
: "${gitbundlebin:="/usr/local/bin/gitbundle"}"
: "${gitbundlehome:="/var/lib/gitbundle"}"
: "${gitbundleconf:="/etc/gitbundle/app.ini"}"
: "${gitbundleuser:="git"}"
: "${sudocmd:="sudo"}"
: "${arch:="linux-amd64"}"
: "${service_start:="$sudocmd systemctl start gitbundle"}"
: "${service_stop:="$sudocmd systemctl stop gitbundle"}"
: "${service_status:="$sudocmd systemctl status gitbundle"}"
: "${backupopts:=""}" # see `gitbundle dump --help` for available options

function gitbundlecmd {
  if [[ $sudocmd = "su" ]]; then
    # `-c` only accept one string as argument.
    "$sudocmd" - "$gitbundleuser" -c "$(printf "%q " "$gitbundlebin" "--config" "$gitbundleconf" "--work-path" "$gitbundlehome" "$@")"
  else
    "$sudocmd" --user "$gitbundleuser" "$gitbundlebin" --config "$gitbundleconf" --work-path "$gitbundlehome" "$@"
  fi
}

function require {
  for exe in "$@"; do
    command -v "$exe" &>/dev/null || (echo "missing dependency '$exe'"; exit 1)
  done
}

# parse command line arguments
while true; do
  case "$1" in
    -v | --version ) gitbundleversion="$2"; shift 2 ;;
    -y | --yes ) no_confirm="yes"; shift ;;
    --ignore-gpg) ignore_gpg="yes"; shift ;;
    "" | -- ) shift; break ;;
    * ) echo "Usage:  [<environment vars>] upgrade.sh [-v <version>] [-y] [--ignore-gpg]"; exit 1;; 
  esac
done

# exit once any command fails. this means that each step should be idempotent!
set -euo pipefail

if [[ -f /etc/os-release ]]; then
  os_release=$(cat /etc/os-release)

  if [[ "$os_release" =~ "OpenWrt" ]]; then
    sudocmd="su"
    service_start="/etc/init.d/gitbundle start"
    service_stop="/etc/init.d/gitbundle stop"
    service_status="/etc/init.d/gitbundle status"
  else
    require systemctl
  fi
fi

require curl xz sha256sum "$sudocmd"

# select version to install
if [[ -z "${gitbundleversion:-}" ]]; then
  require jq
  gitbundleversion=$(curl --connect-timeout 10 -sL https://dl.gitbundle.com/gitbundle/version.json | jq -r .latest.version)
  echo "Latest available version is $gitbundleversion"
fi

# confirm update
echo "Checking currently installed version..."
current=$(gitbundlecmd --version | cut -d ' ' -f 3)
[[ "$current" == "$gitbundleversion" ]] && echo "$current is already installed, stopping." && exit 1
if [[ -z "${no_confirm:-}"  ]]; then
  echo "Make sure to read the changelog first: https://github.com/gitbundle/gitbundle/blob/main/CHANGELOG.md"
  echo "Are you ready to update GitBundle from ${current} to ${gitbundleversion}? (y/N)"
  read -r confirm
  [[ "$confirm" == "y" ]] || [[ "$confirm" == "Y" ]] || exit 1
fi

echo "Upgrading gitbundle from $current to $gitbundleversion ..."

pushd "$(pwd)" &>/dev/null
cd "$gitbundlehome" # needed for gitbundle dump later

# download new binary
binname="gitbundle-${gitbundleversion}-${arch}"
binurl="https://dl.gitbundle.com/gitbundle/${gitbundleversion}/${binname}.xz"
echo "Downloading $binurl..."
curl --connect-timeout 10 --silent --show-error --fail --location -O "$binurl{,.sha256,.asc}"

# validate checksum & gpg signature
sha256sum -c "${binname}.xz.sha256"
if [[ -z "${ignore_gpg:-}" ]]; then
  require gpg
  gpg --keyserver keys.openpgp.org --recv 7C9E68152594688862D62AF62D9AE806EC1592E2
  gpg --verify "${binname}.xz.asc" "${binname}.xz" || { echo 'Signature does not match'; exit 1; }
fi
rm "${binname}".xz.{sha256,asc}

# unpack binary + make executable
xz --decompress --force "${binname}.xz"
chown "$gitbundleuser" "$binname"
chmod +x "$binname"

# stop gitbundle, create backup, replace binary, restart gitbundle
echo "Flushing gitbundle queues at $(date)"
gitbundlecmd manager flush-queues
echo "Stopping gitbundle at $(date)"
$service_stop
echo "Creating backup in $gitbundlehome"
gitbundlecmd dump $backupopts
echo "Updating binary at $gitbundlebin"
cp -f "$gitbundlebin" "$gitbundlebin.bak" && mv -f "$binname" "$gitbundlebin"
$service_start
$service_status

echo "Upgrade to $gitbundleversion successful!"

popd
