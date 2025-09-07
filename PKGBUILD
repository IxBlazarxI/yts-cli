---

### PKGBUILD for `yts-cli`**

```bash
pkgname=yts-cli
pkgver=1.0.0
pkgrel=1
pkgdesc="CLI YouTube scraper in Go (no API needed)"
arch=('x86_64')
url="https://github.com/yourusername/yts-cli"
license=('MIT')
depends=('bash' 'curl' 'grep' 'awk')
source=("https://github.com/yourusername/yts-cli/archive/refs/tags/v$pkgver.tar.gz")
sha256sums=('SKIP')  # Replace with actual sum

build() {
    cd "$srcdir/yts-cli-$pkgver"
    go build -o yts-cli
}

package() {
    install -Dm755 yts-cli "$pkgdir/usr/bin/yts-cli"
}
