cask 'if-six-versions-six-appcasts' do
  if MacOS.release == :snow_leopard
    version '0.1.0'
    sha256 '6ad9613a455798d6d92e5f5f390ab4baa70596bc869ed6b17f5cdd2b28635f06'

    url "https://example.com/snowleopard/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/snowleopard.xml",
            checkpoint: 'a93e9e53c90ab95e1ce83cbc1cbd76102da1bce5330b649872dbd95a1793a03e'
  elsif MacOS.release == :lion
    version '0.2.0'
    sha256 '911fc0c48cb0c70601db5775a9bef1b740dc4cc9f9b46389b9f0563fe7eb94d7'

    url "https://example.com/lion/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/lion.xml",
            checkpoint: '13dfb3758d65d265e4c12336815b2db327683ad38b2a1162cc88ab3579bbfaa1'
  elsif MacOS.release == :mountain_lion
    version '0.3.0'
    sha256 '550613537fc488f3b372af74a4001879f012c8465b816f1b85c6d3446b2cfb49'

    url "https://example.com/mountainlion/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/mountainlion.xml",
            checkpoint: '00af55f25d0c6e53f017a972b77fe4def95f9bb4ec4dc217c520e875fa0071a9'
  elsif MacOS.release == :mavericks
    version '0.4.0'
    sha256 'cd78534ed15ad46912b71339d1417d0d043d8309c2b94415f3ed1b9d1fdfaed0'

    url "https://example.com/mavericks/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/mavericks.xml",
            checkpoint: '9cbe5cfd22b0eb5f159ae634acf615d9c8032699b5a79d37a3046bdaf5677c84'
  elsif MacOS.release == :yosemite
    version '0.5.0'
    sha256 'd1f62539db82b51da84bda2f4885db5e847db8389183be41389efd0ae6edab94'

    url "https://example.com/yosemite/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/yosemite.xml",
            checkpoint: 'f309466aea57120e04b214292d54a9d5e32d018582344b3a62021a91ed8dd69d'
  else
    version '2.0.0'
    sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'

    url "https://example.com/elcapitan/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version.major}/appcast.xml",
            checkpoint: '57956bd3fb23a5673e30dc83ed19d51b43e5a9235756887f3ed90662e6c68fb7'
  end

  name 'Example'
  name 'Example (if-six-versions-six-appcasts)'
  homepage 'https://example.com/'
  license :commercial

  app 'Example (if-six-versions-six-appcasts).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
