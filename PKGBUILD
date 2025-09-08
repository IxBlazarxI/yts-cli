# Maintainer: Your Name blazar11111@gmail.com
pkgname=yts-cli
pkgver=1.1.0
pkgrel=1
pkgdesc="CLI YouTube Web Scraper in GO (no API needed)"
arch=('x86_64')
url="https://github.com/IxBlazarxI/yts-cli"
license=('MIT')
depends=('bash' 'curl' 'grep' 'awk')
makedepends=('go')
source=("https://github.com/IxBlazarxI/yts-cli/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('daea67461ed6487577c64d1ca36abbb987a26606ef03d3d7c61da1da1e7a5b16')

build() {
    cd "$srcdir/yts-cli-$pkgver"
    go build -o yts-cli
}

package() {
    cd "$srcdir/yts-cli-$pkgver"
    install -Dm755 yts-cli "$pkgdir/usr/bin/yts-cli"
}
