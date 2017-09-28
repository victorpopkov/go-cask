cask 'if-three-versions-one-appcast' do
  if MacOS.version <= :tiger
    version '0.9.0'
    sha256 '30c99e8b103eacbe6f6d6e1b54b06ca6d5f3164b4f50094334a517ae95ca8fba'
  elsif MacOS.version <= :leopard
    version '1.0.0'
    sha256 '92521fc3cbd964bdc9f584a991b89fddaa5754ed1cc96d6d42445338669c1305'
  else
    version '2.0.0'
    sha256 'f22abd6773ab232869321ad4b1e47ac0c908febf4f3a2bd10c8066140f741261'

    appcast "https://example.com/sparkle/#{version.major}/appcast.xml",
            checkpoint: '8dc47a4bcec6e46b79fb6fc7b84224f1461f18a2d9f2e5adc94612bb9b97072d'
  end

  url "https://example.com/app_#{version}.dmg"
  name 'Example'
  name 'Example (if-three-versions-one-appcast)'
  homepage 'https://example.com/'

  auto_updates true

  app 'Example (if-three-versions-one-appcast).app', target: 'Example.app'
  binary "#{appdir}/Example.app/Contents/MacOS/example-if", target: 'example'
end
