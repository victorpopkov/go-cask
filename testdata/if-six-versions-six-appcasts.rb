cask 'if-six-versions-six-appcasts' do
  if MacOS.version == :snow_leopard
    version '0.1.0'
    sha256 '6ad9613a455798d6d92e5f5f390ab4baa70596bc869ed6b17f5cdd2b28635f06'

    url "https://example.com/snowleopard/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/snowleopard.xml"
  elsif MacOS.version == :lion
    version '0.2.0'
    sha256 '911fc0c48cb0c70601db5775a9bef1b740dc4cc9f9b46389b9f0563fe7eb94d7'

    url "https://example.com/lion/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/lion.xml"
  elsif MacOS.version == :mountain_lion
    version '0.3.0'
    sha256 '550613537fc488f3b372af74a4001879f012c8465b816f1b85c6d3446b2cfb49'

    url "https://example.com/mountainlion/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/mountainlion.xml"
  elsif MacOS.version == :mavericks
    version '0.4.0'
    sha256 'cd78534ed15ad46912b71339d1417d0d043d8309c2b94415f3ed1b9d1fdfaed0'

    url "https://example.com/mavericks/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/mavericks.xml"
  elsif MacOS.version == :yosemite
    version '0.5.0'
    sha256 'd1f62539db82b51da84bda2f4885db5e847db8389183be41389efd0ae6edab94'

    url "https://example.com/yosemite/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version}/yosemite.xml"
  else
    version '2.0.0'
    sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'

    url "https://example.com/elcapitan/app_#{version}.dmg"
    appcast "https://example.com/sparkle/#{version.major}/appcast.xml"
  end

  name 'Example'
  name 'Example (if-six-versions-six-appcasts)'
  homepage 'https://example.com/'

  app 'Example (if-six-versions-six-appcasts).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
