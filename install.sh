#!/bin/bash
if [[ $(/usr/bin/id -u) -ne 0 ]]; then
  echo "Must be run as root"
  exit
fi

repo="djh20/openjukebox"
package="openjukebox"
release_base_url="https://api.github.com/repos/$repo/releases"
release_type="$1"

extract_json_string() {
  local json=$1
  local key=$2

  echo $json | grep -o "\"$key\": \"[^\"]*" | grep -o "[^\"]*$"
}

if [ "$release_type" == "beta" ]; then
  release_url="$release_base_url?per_page=1"
else
  release_url="$release_base_url/latest"
fi

release=`curl -s "$release_url"`

installed_version=$(dpkg-query --showformat='${Version}' --show $package)
latest_version=$(extract_json_string "$release" tag_name)

echo "Installed Version: $installed_version"
echo "Latest Version: $latest_version"

if [ "$installed_version" != "" ]; then
  dpkg --compare-versions $latest_version gt $installed_version

  if [ $? -ne 0 ]; then
    echo "No update available"
    exit 0
  fi
fi

arch=$(dpkg --print-architecture)

echo "Arch: $arch"

download_url=$(echo $release | grep -o "https:[^\"]*_$arch\.deb")

if [ "$download_url" != "" ]; then
  echo "Downloading package..."
  wget -O /tmp/$package.deb -q "$download_url"

  echo "Installing package..."
  dpkg -i /tmp/$package.deb

else
  echo "Failed to find compatible download URL"
  exit 1
fi